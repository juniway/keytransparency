// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keytransparency_v1_types.proto

/*
Package keytransparency_v1_types is a generated protocol buffer package.

Key Transparency Service

The Key Transparency Service API consists of a map of user names to public
keys. Each user name also has a history of public keys that have been
associated with it.

It is generated from these files:
	keytransparency_v1_types.proto

It has these top-level messages:
	Committed
	EntryUpdate
	Entry
	PublicKey
	KeyValue
	SignedKV
	Mutation
	GetEntryRequest
	GetEntryResponse
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	UpdateEntryRequest
	UpdateEntryResponse
	GetMutationsRequest
	GetMutationsResponse
	GetDomainInfoRequest
	GetDomainInfoResponse
*/
package keytransparency_v1_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import sigpb "github.com/google/trillian/crypto/sigpb"
import trillian "github.com/google/trillian"
import trillian1 "github.com/google/trillian"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Committed represents the data committed to in a cryptographic commitment.
// commitment = HMAC_SHA512_256(key, data)
type Committed struct {
	// key is the 16 byte random commitment key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data being committed to.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Committed) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Committed) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// EntryUpdate contains the user entry update(s).
type EntryUpdate struct {
	// update authorizes the change to entry.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// commitment contains the data committed to in update.commitment.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EntryUpdate) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *EntryUpdate) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*PublicKey `protobuf:"bytes,2,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *Entry) GetAuthorizedKeys() []*PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign MapHeads with.
type PublicKey struct {
	// Key formats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_3072
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_3072 struct {
	RsaVerifyingSha256_3072 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_3072,json=rsaVerifyingSha2563072,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_3072) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_3072() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_3072); ok {
		return x.RsaVerifyingSha256_3072
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_3072)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_3072:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_3072)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_3072
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_3072{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_3072:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_3072)))
		n += len(x.RsaVerifyingSha256_3072)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KeyValue is a map entry.
type KeyValue struct {
	// key contains the map entry key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value contains the map entry value.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// SignedKV is a signed change to a map entry.
type SignedKV struct {
	// key_value is a serialized KeyValue.
	KeyValue *KeyValue `protobuf:"bytes,1,opt,name=key_value,json=keyValue" json:"key_value,omitempty"`
	// signatures on key_value. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves that the correct owner is making this change.
	Signatures map[string]*sigpb.DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// previous contains the hash of the previous entry that this mutation is
	// modifying creating a hash chain of all mutations. The hash used is
	// CommonJSON in "github.com/benlaurie/objecthash/go/objecthash".
	Previous []byte `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SignedKV) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *SignedKV) GetSignatures() map[string]*sigpb.DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *SignedKV) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

// Mutation contains the actual mutation and the inclusion proof of the
// corresponding leaf.
type Mutation struct {
	// update contains the actual mutation information.
	Update *SignedKV `protobuf:"bytes,1,opt,name=update" json:"update,omitempty"`
	// proof contains a leaf and an inclusion proof in the map of the previous
	// epoch. This is used by Storage-less monitors.
	Proof *trillian1.MapLeafInclusion `protobuf:"bytes,2,opt,name=proof" json:"proof,omitempty"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Mutation) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *Mutation) GetProof() *trillian1.MapLeafInclusion {
	if m != nil {
		return m.Proof
	}
	return nil
}

// GetEntryRequest for a user object.
type GetEntryRequest struct {
	// user_id is the user identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,3,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetEntryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *GetEntryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

// GetEntryResponse returns a requested user entry.
type GetEntryResponse struct {
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,1,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,2,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse Merkle
	// Tree.
	LeafProof *trillian1.MapLeafInclusion `protobuf:"bytes,3,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// smr contains the signed map head for the sparse Merkle Tree.
	// smr is also stored in the append only log.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,4,opt,name=smr" json:"smr,omitempty"`
	// log_root is the latest globally consistent log root.
	// TODO: gossip the log root to verify global consistency.
	LogRoot *trillian.SignedLogRoot `protobuf:"bytes,5,opt,name=log_root,json=logRoot" json:"log_root,omitempty"`
	// log_consistency proves that log_root is consistent with previously seen roots.
	LogConsistency [][]byte `protobuf:"bytes,6,rep,name=log_consistency,json=logConsistency,proto3" json:"log_consistency,omitempty"`
	// log_inclusion proves that smr is part of log_root at index=srm.MapRevision.
	LogInclusion [][]byte `protobuf:"bytes,7,rep,name=log_inclusion,json=logInclusion,proto3" json:"log_inclusion,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetEntryResponse) GetVrfProof() []byte {
	if m != nil {
		return m.VrfProof
	}
	return nil
}

func (m *GetEntryResponse) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

func (m *GetEntryResponse) GetLeafProof() *trillian1.MapLeafInclusion {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetEntryResponse) GetLogRoot() *trillian.SignedLogRoot {
	if m != nil {
		return m.LogRoot
	}
	return nil
}

func (m *GetEntryResponse) GetLogConsistency() [][]byte {
	if m != nil {
		return m.LogConsistency
	}
	return nil
}

func (m *GetEntryResponse) GetLogInclusion() [][]byte {
	if m != nil {
		return m.LogInclusion
	}
	return nil
}

// ListEntryHistoryRequest gets a list of historical keys for a user.
type ListEntryHistoryRequest struct {
	// user_id is the user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// start is the starting epoch.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// page_size is the maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,4,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will ommit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,5,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListEntryHistoryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListEntryHistoryRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListEntryHistoryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

// ListEntryHistoryResponse requests a paginated history of keys for a user.
type ListEntryHistoryResponse struct {
	// values represents the list of keys this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// next_start is the next page token to query for pagination.
	// next_start is 0 when there are no more results to fetch.
	NextStart int64 `protobuf:"varint,2,opt,name=next_start,json=nextStart" json:"next_start,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ListEntryHistoryResponse) GetNextStart() int64 {
	if m != nil {
		return m.NextStart
	}
	return 0
}

// UpdateEntryRequest updates a user's profile.
type UpdateEntryRequest struct {
	// user_id specifies the id for the user who's profile is being updated.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,3,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
	// entry_update contains the user submitted update.
	EntryUpdate *EntryUpdate `protobuf:"bytes,4,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateEntryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *UpdateEntryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
type UpdateEntryResponse struct {
	// proof contains a proof that the update has been included in the tree.
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

// GetMutationsRequest contains the input parameters of the GetMutation APIs.
type GetMutationsRequest struct {
	// epoch specifies the epoch number in which mutations will be returned.
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,2,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
	// page_token defines the starting point for pagination. An empty
	// value means start from the beginning. A non-empty value requests the next
	// page of values.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// page_size is the maximum number of epochs to return.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *GetMutationsRequest) Reset()                    { *m = GetMutationsRequest{} }
func (m *GetMutationsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMutationsRequest) ProtoMessage()               {}
func (*GetMutationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetMutationsRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMutationsRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

func (m *GetMutationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *GetMutationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// GetMutationsResponse contains the results of GetMutation APIs.
type GetMutationsResponse struct {
	// epoch specifies the epoch number of the returned mutations.
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// smr contains the signed map root for the sparse Merkle Tree.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,2,opt,name=smr" json:"smr,omitempty"`
	// log_root is the latest globally consistent log root.
	LogRoot *trillian.SignedLogRoot `protobuf:"bytes,3,opt,name=log_root,json=logRoot" json:"log_root,omitempty"`
	// log_consistency proves that log_root is consistent with previously seen roots.
	LogConsistency [][]byte `protobuf:"bytes,4,rep,name=log_consistency,json=logConsistency,proto3" json:"log_consistency,omitempty"`
	// log_inclusion proves that smr is part of log_root at index=srm.MapRevision.
	LogInclusion [][]byte `protobuf:"bytes,5,rep,name=log_inclusion,json=logInclusion,proto3" json:"log_inclusion,omitempty"`
	// mutation contains mutation information.
	Mutations []*Mutation `protobuf:"bytes,6,rep,name=mutations" json:"mutations,omitempty"`
	// next_page_token is the next page token to query for pagination.
	// An empty value means there are no more results to fetch.
	// A non-zero value may be used by the client to fetch the next page of
	// results.
	NextPageToken string `protobuf:"bytes,7,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *GetMutationsResponse) Reset()                    { *m = GetMutationsResponse{} }
func (m *GetMutationsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMutationsResponse) ProtoMessage()               {}
func (*GetMutationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetMutationsResponse) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMutationsResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetMutationsResponse) GetLogRoot() *trillian.SignedLogRoot {
	if m != nil {
		return m.LogRoot
	}
	return nil
}

func (m *GetMutationsResponse) GetLogConsistency() [][]byte {
	if m != nil {
		return m.LogConsistency
	}
	return nil
}

func (m *GetMutationsResponse) GetLogInclusion() [][]byte {
	if m != nil {
		return m.LogInclusion
	}
	return nil
}

func (m *GetMutationsResponse) GetMutations() []*Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

func (m *GetMutationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// GetDomainInfoRequest contains an empty request to query the GetDomainInfo
// APIs.
type GetDomainInfoRequest struct {
}

func (m *GetDomainInfoRequest) Reset()                    { *m = GetDomainInfoRequest{} }
func (m *GetDomainInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDomainInfoRequest) ProtoMessage()               {}
func (*GetDomainInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

// GetDomainInfoResponse contains the results of GetDomainInfo APIs.
type GetDomainInfoResponse struct {
	// Log contains the Log-Tree's info.
	Log *trillian.Tree `protobuf:"bytes,1,opt,name=log" json:"log,omitempty"`
	// Map contains the Map-Tree's info.
	Map *trillian.Tree `protobuf:"bytes,2,opt,name=map" json:"map,omitempty"`
	// Vrf contains the VRF public key.
	Vrf *keyspb.PublicKey `protobuf:"bytes,3,opt,name=vrf" json:"vrf,omitempty"`
}

func (m *GetDomainInfoResponse) Reset()                    { *m = GetDomainInfoResponse{} }
func (m *GetDomainInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDomainInfoResponse) ProtoMessage()               {}
func (*GetDomainInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetDomainInfoResponse) GetLog() *trillian.Tree {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *GetDomainInfoResponse) GetMap() *trillian.Tree {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *GetDomainInfoResponse) GetVrf() *keyspb.PublicKey {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func init() {
	proto.RegisterType((*Committed)(nil), "keytransparency.v1.types.Committed")
	proto.RegisterType((*EntryUpdate)(nil), "keytransparency.v1.types.EntryUpdate")
	proto.RegisterType((*Entry)(nil), "keytransparency.v1.types.Entry")
	proto.RegisterType((*PublicKey)(nil), "keytransparency.v1.types.PublicKey")
	proto.RegisterType((*KeyValue)(nil), "keytransparency.v1.types.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "keytransparency.v1.types.SignedKV")
	proto.RegisterType((*Mutation)(nil), "keytransparency.v1.types.Mutation")
	proto.RegisterType((*GetEntryRequest)(nil), "keytransparency.v1.types.GetEntryRequest")
	proto.RegisterType((*GetEntryResponse)(nil), "keytransparency.v1.types.GetEntryResponse")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "keytransparency.v1.types.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "keytransparency.v1.types.ListEntryHistoryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "keytransparency.v1.types.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "keytransparency.v1.types.UpdateEntryResponse")
	proto.RegisterType((*GetMutationsRequest)(nil), "keytransparency.v1.types.GetMutationsRequest")
	proto.RegisterType((*GetMutationsResponse)(nil), "keytransparency.v1.types.GetMutationsResponse")
	proto.RegisterType((*GetDomainInfoRequest)(nil), "keytransparency.v1.types.GetDomainInfoRequest")
	proto.RegisterType((*GetDomainInfoResponse)(nil), "keytransparency.v1.types.GetDomainInfoResponse")
}

func init() { proto.RegisterFile("keytransparency_v1_types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1057 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x4e, 0x1b, 0xc7,
	0x17, 0x8f, 0xbd, 0xd8, 0xd8, 0x87, 0xaf, 0xfc, 0x27, 0x04, 0xf6, 0xef, 0x2a, 0x11, 0x5a, 0xd4,
	0x36, 0xad, 0x2a, 0x37, 0x38, 0x82, 0x36, 0xe9, 0x45, 0xd3, 0x24, 0x55, 0x40, 0x80, 0x84, 0x86,
	0x84, 0x5e, 0xae, 0x06, 0x7b, 0xbc, 0x8c, 0x58, 0xef, 0x4c, 0x67, 0x66, 0xad, 0x2c, 0x52, 0xa5,
	0x3e, 0x40, 0xa5, 0x4a, 0x7d, 0x86, 0x3e, 0x43, 0x5f, 0xa0, 0x6f, 0xd1, 0xa7, 0xa9, 0xe6, 0x63,
	0xfd, 0x01, 0x26, 0x84, 0x5c, 0xf4, 0xc6, 0x9e, 0xf3, 0xb5, 0xe7, 0x9c, 0xdf, 0xf9, 0x9d, 0xd9,
	0x85, 0x87, 0xe7, 0xb4, 0xd0, 0x92, 0x64, 0x4a, 0x10, 0x49, 0xb3, 0x6e, 0x11, 0x0f, 0xb7, 0x62,
	0x5d, 0x08, 0xaa, 0xda, 0x42, 0x72, 0xcd, 0x51, 0x78, 0xc9, 0xde, 0x1e, 0x6e, 0xb5, 0xad, 0xbd,
	0xd5, 0xea, 0xca, 0x42, 0x68, 0xfe, 0xf5, 0x39, 0x2d, 0x94, 0x38, 0xf5, 0x7f, 0x2e, 0xaa, 0x15,
	0x7a, 0x9b, 0x62, 0x89, 0x38, 0x75, 0xbf, 0xde, 0xb2, 0xac, 0x25, 0x4b, 0x53, 0x46, 0x32, 0x2f,
	0xaf, 0x95, 0x72, 0x3c, 0x20, 0x22, 0x26, 0x82, 0x39, 0x7d, 0xb4, 0x05, 0xcd, 0x97, 0x7c, 0x30,
	0x60, 0x5a, 0xd3, 0x1e, 0xba, 0x0b, 0xc1, 0x39, 0x2d, 0xc2, 0xca, 0x46, 0xe5, 0xd1, 0x22, 0x36,
	0x47, 0x84, 0x60, 0xae, 0x47, 0x34, 0x09, 0xab, 0x56, 0x65, 0xcf, 0xd1, 0x6f, 0x15, 0x58, 0xf8,
	0x31, 0xd3, 0xb2, 0x78, 0x2b, 0x7a, 0x44, 0x53, 0xf4, 0x0c, 0xea, 0xb9, 0x3d, 0x59, 0xaf, 0x85,
	0x4e, 0xd4, 0xbe, 0xae, 0x97, 0xf6, 0x31, 0x4b, 0x32, 0xda, 0xdb, 0x3f, 0xc1, 0x3e, 0x02, 0xfd,
	0x00, 0xcd, 0x6e, 0x99, 0x3e, 0x0c, 0x6c, 0xf8, 0xe6, 0xf5, 0xe1, 0xa3, 0x4a, 0xf1, 0x38, 0x2a,
	0xca, 0xa1, 0x66, 0xab, 0x41, 0x0f, 0x01, 0x9c, 0x76, 0x40, 0x33, 0xed, 0x9b, 0x98, 0xd0, 0xa0,
	0x03, 0x58, 0x21, 0xb9, 0x3e, 0xe3, 0x92, 0x5d, 0xd0, 0x5e, 0x6c, 0x70, 0x0c, 0xab, 0x1b, 0xc1,
	0xfb, 0x33, 0x1e, 0xe5, 0xa7, 0x29, 0xeb, 0xee, 0xd3, 0x02, 0x2f, 0x8f, 0x63, 0xf7, 0x69, 0xa1,
	0xa2, 0x3f, 0x2b, 0xd0, 0x1c, 0x59, 0x51, 0x0b, 0xe6, 0x69, 0xaf, 0xb3, 0xbd, 0xbd, 0xf5, 0xd4,
	0x25, 0xde, 0xbd, 0x83, 0x4b, 0x05, 0xfa, 0x0e, 0xfe, 0x2f, 0x15, 0x89, 0x87, 0x54, 0xb2, 0x7e,
	0xc1, 0xb2, 0x24, 0x56, 0x67, 0xa4, 0xb3, 0xbd, 0x13, 0x3f, 0x79, 0xfc, 0x4d, 0xc7, 0x01, 0xbb,
	0x7b, 0x07, 0xaf, 0x49, 0x45, 0x4e, 0x4a, 0x8f, 0x63, 0xeb, 0x60, 0xec, 0xa8, 0x03, 0xab, 0xb4,
	0xdb, 0x9b, 0x0a, 0x17, 0x9d, 0xed, 0x1d, 0x8b, 0x95, 0x89, 0x43, 0xd6, 0x3a, 0x8a, 0x3c, 0xea,
	0x6c, 0xef, 0xbc, 0x00, 0x68, 0x9c, 0xd3, 0xc2, 0xd2, 0x2b, 0xea, 0x40, 0x63, 0x9f, 0x16, 0x27,
	0x24, 0xcd, 0xe9, 0x8c, 0xf1, 0xae, 0x42, 0x6d, 0x68, 0x4c, 0x7e, 0xbe, 0x4e, 0x88, 0xfe, 0xa8,
	0x42, 0xa3, 0x9c, 0x14, 0xfa, 0x1e, 0x9a, 0xe6, 0x61, 0xce, 0xad, 0x72, 0xd3, 0x80, 0xcb, 0x5c,
	0xd8, 0x54, 0xe0, 0xb2, 0x62, 0x00, 0xc5, 0x92, 0x8c, 0xe8, 0x5c, 0xd2, 0x12, 0xf1, 0xce, 0xcd,
	0x14, 0xb1, 0x07, 0x17, 0x64, 0xc7, 0x8b, 0x27, 0x9e, 0x82, 0x5a, 0xd0, 0x10, 0x92, 0x0e, 0x19,
	0xcf, 0x95, 0x43, 0x02, 0x8f, 0xe4, 0xd6, 0x5b, 0x58, 0xb9, 0x14, 0x3a, 0xd9, 0x78, 0xd3, 0x35,
	0xfe, 0xd5, 0x64, 0xe3, 0x0b, 0x9d, 0xb5, 0xb6, 0xdb, 0x9d, 0x57, 0x2c, 0x61, 0x9a, 0xa4, 0x69,
	0xe1, 0xaa, 0xf0, 0x80, 0x3c, 0xab, 0x7e, 0x5b, 0x89, 0xde, 0x41, 0xe3, 0x30, 0xd7, 0x44, 0x33,
	0x9e, 0x4d, 0x30, 0xbe, 0x72, 0x6b, 0xc6, 0x3f, 0x86, 0x9a, 0x90, 0x9c, 0xf7, 0x7d, 0xe6, 0x56,
	0x7b, 0xb4, 0xa8, 0x87, 0x44, 0x1c, 0x50, 0xd2, 0xdf, 0xcb, 0xba, 0x69, 0xae, 0x18, 0xcf, 0xb0,
	0x73, 0x8c, 0x18, 0xac, 0xbc, 0xa6, 0xda, 0x81, 0x40, 0x7f, 0xce, 0xa9, 0xd2, 0x68, 0x1d, 0xe6,
	0x73, 0x45, 0x65, 0xcc, 0x7a, 0xbe, 0xa9, 0xba, 0x11, 0xf7, 0x7a, 0xe8, 0x3e, 0xd4, 0x89, 0x10,
	0x46, 0x5f, 0xb5, 0xfa, 0x1a, 0x11, 0x62, 0xaf, 0x87, 0x3e, 0x83, 0x95, 0x3e, 0x93, 0x4a, 0xc7,
	0x5a, 0x52, 0x1a, 0x2b, 0x76, 0x41, 0x2d, 0x6c, 0x01, 0x5e, 0xb2, 0xea, 0x37, 0x92, 0xd2, 0x63,
	0x76, 0x41, 0xa3, 0x7f, 0xaa, 0x70, 0x77, 0x9c, 0x4b, 0x09, 0x9e, 0x29, 0x8a, 0x3e, 0x81, 0xe6,
	0x50, 0xf6, 0x63, 0x57, 0xb5, 0x23, 0x4f, 0x63, 0x28, 0xfb, 0x47, 0x46, 0x9e, 0x5e, 0xe0, 0xea,
	0xc7, 0x2c, 0x30, 0x7a, 0x0a, 0x90, 0x52, 0x52, 0x26, 0x08, 0x6e, 0x84, 0xa5, 0x69, 0xbc, 0x5d,
	0xf6, 0x2f, 0x20, 0x50, 0x03, 0x19, 0xce, 0xd9, 0x98, 0xf5, 0x71, 0x8c, 0x43, 0xfd, 0x90, 0x08,
	0xcc, 0xb9, 0xc6, 0xc6, 0x07, 0x75, 0xa0, 0x91, 0xf2, 0x24, 0x96, 0x9c, 0xeb, 0xb0, 0x36, 0xdb,
	0xff, 0x80, 0x27, 0xd6, 0x7f, 0x3e, 0x75, 0x07, 0xf4, 0x39, 0xac, 0x98, 0x98, 0x2e, 0xcf, 0x14,
	0x53, 0xda, 0xb4, 0x12, 0xd6, 0x37, 0x82, 0x47, 0x8b, 0x78, 0x39, 0xe5, 0xc9, 0xcb, 0xb1, 0x16,
	0x6d, 0xc2, 0x92, 0x71, 0x64, 0x65, 0x8d, 0xe1, 0xbc, 0x75, 0x5b, 0x4c, 0x79, 0x32, 0xaa, 0xdb,
	0xdc, 0x18, 0xeb, 0x07, 0x4c, 0x39, 0x74, 0x77, 0x99, 0xd2, 0xfc, 0x03, 0x06, 0xba, 0x0a, 0x35,
	0xa5, 0x89, 0xd4, 0x16, 0xdb, 0x00, 0x3b, 0xc1, 0x8c, 0x44, 0x90, 0x64, 0x62, 0x92, 0x35, 0xdc,
	0x30, 0x0a, 0x33, 0xc4, 0x09, 0x0e, 0xcc, 0xdd, 0xc0, 0x81, 0xda, 0x2c, 0x0e, 0xfc, 0x02, 0xe1,
	0xd5, 0x2a, 0x3d, 0x15, 0x5e, 0x40, 0xdd, 0x6e, 0x84, 0x0a, 0x2b, 0x76, 0x8f, 0xbf, 0xbc, 0x7e,
	0xd4, 0x97, 0x69, 0x84, 0x7d, 0x24, 0x7a, 0x00, 0x90, 0xd1, 0x77, 0x3a, 0x9e, 0x6c, 0xab, 0x69,
	0x34, 0xc7, 0x46, 0x11, 0xfd, 0x55, 0x01, 0xe4, 0x5e, 0x2c, 0xff, 0x05, 0xe3, 0xd1, 0x2e, 0x2c,
	0x52, 0x93, 0x27, 0xf6, 0x0b, 0xed, 0xa8, 0xf4, 0xe9, 0xf5, 0x7d, 0x4d, 0xbc, 0xf9, 0xf0, 0x02,
	0x1d, 0x0b, 0xd1, 0x4f, 0x70, 0x6f, 0xaa, 0x6e, 0x0f, 0xd9, 0xf3, 0x72, 0xdf, 0xdd, 0x55, 0x71,
	0x1b, 0xc4, 0xfc, 0xfe, 0xff, 0x5e, 0x81, 0x7b, 0xaf, 0xa9, 0x2e, 0x6f, 0x1f, 0x55, 0x42, 0xb2,
	0x0a, 0x35, 0x2a, 0x78, 0xf7, 0xcc, 0x3e, 0x39, 0xc0, 0x4e, 0x98, 0xd5, 0x78, 0x75, 0x56, 0xe3,
	0x0f, 0x00, 0x2c, 0x85, 0x34, 0x3f, 0xa7, 0x99, 0xc5, 0xa6, 0x89, 0x2d, 0xa9, 0xde, 0x18, 0xc5,
	0x34, 0xc3, 0xe6, 0xa6, 0x19, 0x16, 0xfd, 0x5d, 0x85, 0xd5, 0xe9, 0x8a, 0x7c, 0xb3, 0xb3, 0x4b,
	0xf2, 0x5b, 0x5a, 0xbd, 0xe5, 0x96, 0x06, 0x1f, 0xbf, 0xa5, 0x73, 0x1f, 0xb6, 0xa5, 0xb5, 0xab,
	0x5b, 0x8a, 0x9e, 0x43, 0x73, 0x50, 0xf6, 0x65, 0xb7, 0xfd, 0xbd, 0xd7, 0x7b, 0x09, 0x01, 0x1e,
	0x07, 0x99, 0x09, 0x58, 0x82, 0x4f, 0xc0, 0x3b, 0x6f, 0xe1, 0x5d, 0x32, 0xea, 0xa3, 0x12, 0xe2,
	0x68, 0xcd, 0x82, 0xf8, 0x8a, 0x0f, 0x08, 0xcb, 0xf6, 0xb2, 0x3e, 0xf7, 0x73, 0x8d, 0x7e, 0xad,
	0xc0, 0xfd, 0x4b, 0x06, 0x0f, 0xef, 0x06, 0x04, 0x29, 0x4f, 0x3c, 0x93, 0x96, 0xc7, 0xc0, 0x98,
	0xa1, 0x62, 0x63, 0x32, 0x1e, 0x03, 0x22, 0x3c, 0xd4, 0x57, 0x3c, 0x06, 0x44, 0xa0, 0x4d, 0x08,
	0x86, 0xb2, 0xbc, 0x66, 0xff, 0xd7, 0xf6, 0x9f, 0x93, 0xe3, 0xef, 0x1c, 0x63, 0x3d, 0xad, 0xdb,
	0x8f, 0xc3, 0x27, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x87, 0x52, 0x53, 0x62, 0xb6, 0x0a, 0x00,
	0x00,
}
