// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signer

import (
	"fmt"
	"math"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/keytransparency/core/appender"
	"github.com/google/keytransparency/core/mutator"
	"github.com/google/keytransparency/core/transaction"

	"github.com/golang/glog"
	"github.com/google/trillian"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"

	tpb "github.com/google/keytransparency/core/proto/keytransparency_v1_types"
	"github.com/google/trillian/util"
)

var (
	mutationsCtr = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "kt_signer_mutations",
		Help: "Number of mutations the signer has processed.",
	})
	indexCtr = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "kt_signer_mutations_unique",
		Help: "Number of mutations the signer has processed post per epoch dedupe.",
	})
	mapUpdateHist = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "kt_signer_map_update_seconds",
		Help:    "Seconds waiting for map update",
		Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, math.Inf(1)},
	})
	createEpochHist = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "kt_signer_create_epoch_seconds",
		Help:    "Seconds spent generating epoch",
		Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, math.Inf(1)},
	})
)

func init() {
	prometheus.MustRegister(mutationsCtr)
	prometheus.MustRegister(indexCtr)
	prometheus.MustRegister(mapUpdateHist)
	prometheus.MustRegister(createEpochHist)
}

// Signer processes mutations and sends them to the trillian map.
type Signer struct {
	realm     string
	mapID     int64
	tmap      trillian.TrillianMapClient
	logID     int64
	sths      appender.Remote
	mutator   mutator.Mutator
	mutations mutator.Mutation
	factory   transaction.Factory
}

// New creates a new instance of the signer.
func New(realm string,
	mapID int64,
	tmap trillian.TrillianMapClient,
	logID int64,
	sths appender.Remote,
	mutator mutator.Mutator,
	mutations mutator.Mutation,
	factory transaction.Factory) *Signer {
	return &Signer{
		realm:     realm,
		mapID:     mapID,
		tmap:      tmap,
		logID:     logID,
		sths:      sths,
		mutator:   mutator,
		mutations: mutations,
		factory:   factory,
	}
}

// StartSigning advance epochs once per minInterval, if there were mutations,
// and at least once per maxElapsed minIntervals.
func (s *Signer) StartSigning(ctx context.Context, minInterval, maxInterval time.Duration) {
	var rootResp *trillian.GetSignedMapRootResponse
	ctxTime, cancel := context.WithTimeout(ctx, minInterval)
	rootResp, err := s.tmap.GetSignedMapRoot(ctxTime, &trillian.GetSignedMapRootRequest{
		MapId: s.mapID,
	})
	if err != nil {
		glog.Infof("GetSignedMapRoot failed: %v", err)
		// Immediately create new epoch and write new sth:
		if err := s.CreateEpoch(ctxTime, true); err != nil {
			glog.Errorf("CreateEpoch failed: %v", err)
		}
		// Request map head again to get the exact time it was created:
		rootResp, err = s.tmap.GetSignedMapRoot(ctxTime, &trillian.GetSignedMapRootRequest{
			MapId: s.mapID,
		})
		if err != nil {
			glog.Errorf("GetSignedMapRoot failed after CreateEpoch: %v", err)
		}
	}
	cancel()
	// Fetch last time from previous map head (as stored in the map server)
	mapRoot := rootResp.GetMapRoot()
	last := time.Unix(0, mapRoot.GetTimestampNanos())
	// Start issuing epochs:
	clock := util.SystemTimeSource{}
	tc := time.Tick(minInterval)
	for f := range genEpochTicks(clock, last, tc, minInterval, maxInterval) {
		ctxTime, cancel := context.WithTimeout(ctx, minInterval)
		if err := s.CreateEpoch(ctxTime, f); err != nil {
			glog.Errorf("CreateEpoch failed: %v", err)
		}
		cancel()
	}
}

// genEpochTicks returns and sends to a bool channel every time an epoch should
// be created. If the boolean value is true this indicates that the epoch should
// be created regardless of whether mutations exist.
func genEpochTicks(t util.TimeSource, last time.Time, minTick <-chan time.Time, minElapsed, maxElapsed time.Duration) <-chan bool {
	enforce := make(chan bool)
	go func() {
		// Do not wait for the first minDuration to pass but directly resume from
		// last
		if (t.Now().Sub(last) + minElapsed) >= maxElapsed {
			enforce <- true
			last = t.Now()
		}

		for now := range minTick {
			if (now.Sub(last) + minElapsed) >= maxElapsed {
				enforce <- true
				last = now
			} else {
				enforce <- false
			}
		}
	}()

	return enforce
}

// newMutations returns a list of mutations to process and highest sequence
// number returned.
func (s *Signer) newMutations(ctx context.Context, startSequence int64) ([]*tpb.SignedKV, int64, error) {
	txn, err := s.factory.NewTxn(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("NewDBTxn(): %v", err)
	}

	maxSequence, mutations, err := s.mutations.ReadAll(txn, uint64(startSequence))
	if err != nil {
		if err := txn.Rollback(); err != nil {
			glog.Errorf("Cannot rollback the transaction: %v", err)
		}
		return nil, 0, fmt.Errorf("ReadAll(%v): %v", startSequence, err)
	}

	if err := txn.Commit(); err != nil {
		return nil, 0, fmt.Errorf("txn.Commit(): %v", err)
	}
	return mutations, int64(maxSequence), nil
}

// toArray returns the first 32 bytes from b.
// If b is less than 32 bytes long, the output is zero padded.
func toArray(b []byte) [32]byte {
	var i [32]byte
	copy(i[:], b)
	return i
}

// applyMutations takes the set of mutations and applies them to given leafs.
// Multiple mutations for the same leaf will be applied to provided leaf.
// The last valid mutation for each leaf is included in the output.
// Returns a list of map leaves that should be updated.
func (s *Signer) applyMutations(mutations []*tpb.SignedKV, leaves []*trillian.MapLeaf) ([]*trillian.MapLeaf, error) {
	// Put leaves in a map from index to leaf value.
	leafMap := make(map[[32]byte]*trillian.MapLeaf)
	for _, l := range leaves {
		leafMap[toArray(l.Index)] = l
	}

	retMap := make(map[[32]byte]*trillian.MapLeaf)
	for _, m := range mutations {
		index := m.GetKeyValue().Key
		var oldValue []byte // If no map leaf was found, oldValue will be nil.
		if leaf, ok := leafMap[toArray(index)]; ok {
			oldValue = leaf.LeafValue
		}

		// TODO: change mutator interface to accept objects directly.
		mData, err := proto.Marshal(m)
		if err != nil {
			return nil, err
		}
		newValue, err := s.mutator.Mutate(oldValue, mData)
		if err != nil {
			glog.Warningf("Mutate(): %v", err)
			continue // A bad mutation should not make the whole batch fail.
		}

		retMap[toArray(index)] = &trillian.MapLeaf{
			Index:     index,
			LeafValue: newValue,
		}
	}
	// Convert return map back into a list.
	ret := make([]*trillian.MapLeaf, 0, len(retMap))
	for _, v := range retMap {
		ret = append(ret, v)
	}
	return ret, nil
}

// CreateEpoch signs the current map head.
func (s *Signer) CreateEpoch(ctx context.Context, forceNewEpoch bool) error {
	glog.V(2).Infof("CreateEpoch: starting sequencing run")
	start := time.Now()
	// Get the current root.
	rootResp, err := s.tmap.GetSignedMapRoot(ctx, &trillian.GetSignedMapRootRequest{
		MapId: s.mapID,
	})
	if err != nil {
		return fmt.Errorf("GetSignedMapRoot(%v): %v", s.mapID, err)
	}
	startSequence := rootResp.GetMapRoot().GetMetadata().GetHighestFullyCompletedSeq()
	revision := rootResp.GetMapRoot().GetMapRevision()
	glog.V(3).Infof("CreateEpoch: Previous SignedMapRoot: {Revision: %v, HighestFullyCompletedSeq: %v}", revision, startSequence)

	// Get the list of new mutations to process.
	mutations, seq, err := s.newMutations(ctx, startSequence)
	if err != nil {
		return fmt.Errorf("newMutations(%v): %v", startSequence, err)
	}

	// Don't create epoch if there is nothing to process unless explicitly
	// specified by caller
	if len(mutations) == 0 && !forceNewEpoch {
		glog.Infof("CreateEpoch: No mutations found. Exiting.")
		return nil
	}

	// Get current leaf values.
	indexes := make([][]byte, 0, len(mutations))
	for _, m := range mutations {
		indexes = append(indexes, m.KeyValue.Key)
	}
	glog.V(2).Infof("CreateEpoch: len(mutations): %v, len(indexes): %v",
		len(mutations), len(indexes))
	getResp, err := s.tmap.GetLeaves(ctx, &trillian.GetMapLeavesRequest{
		MapId:    s.mapID,
		Index:    indexes,
		Revision: -1, // Get the latest version.
	})
	if err != nil {
		return err
	}
	glog.V(3).Infof("CreateEpoch: len(GetLeaves.MapLeafInclusions): %v",
		len(getResp.MapLeafInclusion))

	// Trust the leaf values provided by the map server.
	// If the map server is run by an untrusted entity, perform inclusion
	// and signature verification here.
	leaves := make([]*trillian.MapLeaf, 0, len(getResp.MapLeafInclusion))
	for _, m := range getResp.MapLeafInclusion {
		leaves = append(leaves, m.Leaf)
	}

	// Apply mutations to values.
	newLeaves, err := s.applyMutations(mutations, leaves)
	if err != nil {
		return err
	}
	glog.V(2).Infof("CreateEpoch: applied %v mutations to %v leaves",
		len(mutations), len(leaves))

	// Set new leaf values.
	mapSetStart := time.Now()
	setResp, err := s.tmap.SetLeaves(ctx, &trillian.SetMapLeavesRequest{
		MapId:  s.mapID,
		Leaves: newLeaves,
		MapperData: &trillian.MapperMetadata{
			HighestFullyCompletedSeq: seq,
		},
	})
	mapSetEnd := time.Now()
	if err != nil {
		return err
	}
	revision = setResp.GetMapRoot().GetMapRevision()
	glog.V(2).Infof("CreateEpoch: SetLeaves:{Revision: %v, HighestFullyCompletedSeq: %v}", revision, seq)

	// Put SignedMapHead in an append only log.
	if err := s.sths.Write(ctx, s.logID, revision, setResp.GetMapRoot()); err != nil {
		// TODO(gdbelvin): If the log doesn't do this, we need to generate an emergency alert.
		return fmt.Errorf("sths.Write(%v, %v): %v", s.logID, revision, err)
	}

	mutationsCtr.Add(float64(len(mutations)))
	indexCtr.Add(float64(len(indexes)))
	mapUpdateHist.Observe(mapSetEnd.Sub(mapSetStart).Seconds())
	createEpochHist.Observe(time.Now().Sub(start).Seconds())
	glog.Infof("CreatedEpoch: rev: %v, root: %x", revision, setResp.GetMapRoot().GetRootHash())
	return nil
}
