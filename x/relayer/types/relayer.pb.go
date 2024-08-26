// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: goat/relayer/v1/relayer.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// VoterStatus is the status of a voter.
type VoterStatus int32

const (
	// UNSPECIFIED defines an invalid voter status.
	Unspecified VoterStatus = 0
	// UNACTIVATED defines a voter that is not able to vote and sign a new relayer tx
	Unactivated VoterStatus = 1
	// PENDING defines a voter that is pending to Activated status
	Pending VoterStatus = 2
	// ACTIVATED defines a voter that is actived.
	Activated VoterStatus = 3
)

var VoterStatus_name = map[int32]string{
	0: "VOTER_STATUS_UNSPECIFIED",
	1: "VOTER_STATUS_UNACTIVATED",
	2: "VOTER_STATUS_PENDING",
	3: "VOTER_STATUS_ACTIVATED",
}

var VoterStatus_value = map[string]int32{
	"VOTER_STATUS_UNSPECIFIED": 0,
	"VOTER_STATUS_UNACTIVATED": 1,
	"VOTER_STATUS_PENDING":     2,
	"VOTER_STATUS_ACTIVATED":   3,
}

func (x VoterStatus) String() string {
	return proto.EnumName(VoterStatus_name, int32(x))
}

func (VoterStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0a671d8480f2e392, []int{0}
}

// Relayer represents the current relayer group state
type Relayer struct {
	Version     uint64    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Threshold   uint64    `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Proposer    string    `protobuf:"bytes,3,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Voters      []string  `protobuf:"bytes,4,rep,name=voters,proto3" json:"voters,omitempty"`
	LastElected time.Time `protobuf:"bytes,5,opt,name=last_elected,json=lastElected,proto3,stdtime" json:"last_elected"`
}

func (m *Relayer) Reset()         { *m = Relayer{} }
func (m *Relayer) String() string { return proto.CompactTextString(m) }
func (*Relayer) ProtoMessage()    {}
func (*Relayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a671d8480f2e392, []int{0}
}
func (m *Relayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relayer.Merge(m, src)
}
func (m *Relayer) XXX_Size() int {
	return m.Size()
}
func (m *Relayer) XXX_DiscardUnknown() {
	xxx_messageInfo_Relayer.DiscardUnknown(m)
}

var xxx_messageInfo_Relayer proto.InternalMessageInfo

func (m *Relayer) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Relayer) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *Relayer) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *Relayer) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *Relayer) GetLastElected() time.Time {
	if m != nil {
		return m.LastElected
	}
	return time.Time{}
}

// Voter the relayer voter
type Voter struct {
	// tx_key is the uncompressed secp256k1 pubkey, it's for submitting new relayer proposals
	TxKey []byte `protobuf:"bytes,1,opt,name=tx_key,json=txKey,proto3" json:"tx_key,omitempty"`
	// vote key is the bls12-381 G2 compressed pubkey(96 bytes)
	VoteKey []byte `protobuf:"bytes,2,opt,name=vote_key,json=voteKey,proto3" json:"vote_key,omitempty"`
	// status represents the current voter status
	Status VoterStatus `protobuf:"varint,3,opt,name=status,proto3,enum=goat.relayer.v1.VoterStatus" json:"status,omitempty"`
}

func (m *Voter) Reset()         { *m = Voter{} }
func (m *Voter) String() string { return proto.CompactTextString(m) }
func (*Voter) ProtoMessage()    {}
func (*Voter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a671d8480f2e392, []int{1}
}
func (m *Voter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Voter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Voter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Voter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voter.Merge(m, src)
}
func (m *Voter) XXX_Size() int {
	return m.Size()
}
func (m *Voter) XXX_DiscardUnknown() {
	xxx_messageInfo_Voter.DiscardUnknown(m)
}

var xxx_messageInfo_Voter proto.InternalMessageInfo

func (m *Voter) GetTxKey() []byte {
	if m != nil {
		return m.TxKey
	}
	return nil
}

func (m *Voter) GetVoteKey() []byte {
	if m != nil {
		return m.VoteKey
	}
	return nil
}

func (m *Voter) GetStatus() VoterStatus {
	if m != nil {
		return m.Status
	}
	return Unspecified
}

// PublicKey defines the keys available for use with relayer
type PublicKey struct {
	// the key
	//
	// Types that are valid to be assigned to Key:
	//
	//	*PublicKey_Secp256K1
	//	*PublicKey_Schnorr
	Key isPublicKey_Key `protobuf_oneof:"key"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a671d8480f2e392, []int{2}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(m, src)
}
func (m *PublicKey) XXX_Size() int {
	return m.Size()
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

type isPublicKey_Key interface {
	isPublicKey_Key()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
	Compare(interface{}) int
}

type PublicKey_Secp256K1 struct {
	Secp256K1 []byte `protobuf:"bytes,1,opt,name=secp256k1,proto3,oneof" json:"secp256k1,omitempty"`
}
type PublicKey_Schnorr struct {
	Schnorr []byte `protobuf:"bytes,2,opt,name=schnorr,proto3,oneof" json:"schnorr,omitempty"`
}

func (*PublicKey_Secp256K1) isPublicKey_Key() {}
func (*PublicKey_Schnorr) isPublicKey_Key()   {}

func (m *PublicKey) GetKey() isPublicKey_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PublicKey) GetSecp256K1() []byte {
	if x, ok := m.GetKey().(*PublicKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

func (m *PublicKey) GetSchnorr() []byte {
	if x, ok := m.GetKey().(*PublicKey_Schnorr); ok {
		return x.Schnorr
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PublicKey) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PublicKey_Secp256K1)(nil),
		(*PublicKey_Schnorr)(nil),
	}
}

// message Votes the proposal vote result
type Votes struct {
	// the proposal sequence
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// voters represents the voter bitmap
	Voters []byte `protobuf:"bytes,2,opt,name=voters,proto3" json:"voters,omitempty"`
	// signature is the aggregate signature by voters
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Votes) Reset()         { *m = Votes{} }
func (m *Votes) String() string { return proto.CompactTextString(m) }
func (*Votes) ProtoMessage()    {}
func (*Votes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a671d8480f2e392, []int{3}
}
func (m *Votes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Votes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Votes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Votes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Votes.Merge(m, src)
}
func (m *Votes) XXX_Size() int {
	return m.Size()
}
func (m *Votes) XXX_DiscardUnknown() {
	xxx_messageInfo_Votes.DiscardUnknown(m)
}

var xxx_messageInfo_Votes proto.InternalMessageInfo

func (m *Votes) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Votes) GetVoters() []byte {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *Votes) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("goat.relayer.v1.VoterStatus", VoterStatus_name, VoterStatus_value)
	proto.RegisterType((*Relayer)(nil), "goat.relayer.v1.Relayer")
	proto.RegisterType((*Voter)(nil), "goat.relayer.v1.Voter")
	proto.RegisterType((*PublicKey)(nil), "goat.relayer.v1.PublicKey")
	proto.RegisterType((*Votes)(nil), "goat.relayer.v1.Votes")
}

func init() { proto.RegisterFile("goat/relayer/v1/relayer.proto", fileDescriptor_0a671d8480f2e392) }

var fileDescriptor_0a671d8480f2e392 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xf6, 0x36, 0xcd, 0x8f, 0x37, 0x85, 0x56, 0x56, 0xa9, 0x8c, 0x55, 0x1c, 0x2b, 0x12, 0x52,
	0x00, 0x61, 0x2b, 0xe5, 0xe7, 0xc0, 0x2d, 0x6d, 0x4d, 0x1b, 0x21, 0x95, 0xc8, 0x49, 0x2b, 0xe0,
	0x12, 0x39, 0xce, 0xd4, 0xb1, 0x92, 0x78, 0xcd, 0xee, 0xc6, 0x34, 0x6f, 0x80, 0x7a, 0xea, 0x0b,
	0x54, 0x02, 0xf1, 0x1e, 0x5c, 0xb8, 0xf4, 0xd8, 0x23, 0x27, 0x40, 0xed, 0xa5, 0x8f, 0x81, 0xbc,
	0x8e, 0x9b, 0x16, 0x71, 0x9b, 0xef, 0x9b, 0x6f, 0x66, 0xbe, 0x99, 0xb5, 0xf1, 0x03, 0x9f, 0xb8,
	0xdc, 0xa2, 0x30, 0x72, 0xa7, 0x40, 0xad, 0xb8, 0x9e, 0x85, 0x66, 0x44, 0x09, 0x27, 0xca, 0x72,
	0x92, 0x36, 0x33, 0x2e, 0xae, 0x6b, 0xba, 0x47, 0xd8, 0x98, 0x30, 0xab, 0xe7, 0x32, 0xb0, 0xe2,
	0x7a, 0x0f, 0xb8, 0x5b, 0xb7, 0x3c, 0x12, 0x84, 0x69, 0x81, 0xb6, 0xea, 0x13, 0x9f, 0x88, 0xd0,
	0x4a, 0xa2, 0x19, 0x5b, 0xf1, 0x09, 0xf1, 0x47, 0x60, 0x09, 0xd4, 0x9b, 0x1c, 0x5a, 0x3c, 0x18,
	0x03, 0xe3, 0xee, 0x38, 0x4a, 0x05, 0xd5, 0xef, 0x08, 0x17, 0x9d, 0x74, 0x8a, 0xa2, 0xe2, 0x62,
	0x0c, 0x94, 0x05, 0x24, 0x54, 0x91, 0x81, 0x6a, 0x8b, 0x4e, 0x06, 0x95, 0x75, 0x2c, 0xf3, 0x01,
	0x05, 0x36, 0x20, 0xa3, 0xbe, 0xba, 0x20, 0x72, 0x73, 0x42, 0xd1, 0x70, 0x29, 0xa2, 0x24, 0x22,
	0x0c, 0xa8, 0x9a, 0x33, 0x50, 0x4d, 0x76, 0xae, 0xb1, 0xb2, 0x86, 0x0b, 0x31, 0xe1, 0x40, 0x99,
	0xba, 0x68, 0xe4, 0x6a, 0xb2, 0x33, 0x43, 0xca, 0x0e, 0x5e, 0x1a, 0xb9, 0x8c, 0x77, 0x61, 0x04,
	0x1e, 0x87, 0xbe, 0x9a, 0x37, 0x50, 0xad, 0xbc, 0xa1, 0x99, 0xa9, 0x5f, 0x33, 0xf3, 0x6b, 0x76,
	0x32, 0xbf, 0x9b, 0xa5, 0xb3, 0x5f, 0x15, 0xe9, 0xe4, 0x77, 0x05, 0x39, 0xe5, 0xa4, 0xd2, 0x4e,
	0x0b, 0xab, 0x04, 0xe7, 0x0f, 0x92, 0x96, 0xca, 0x3d, 0x5c, 0xe0, 0x47, 0xdd, 0x21, 0x4c, 0x85,
	0xf9, 0x25, 0x27, 0xcf, 0x8f, 0xde, 0xc0, 0x54, 0xb9, 0x8f, 0x4b, 0xc9, 0x48, 0x91, 0x58, 0x10,
	0x89, 0x62, 0x82, 0x93, 0xd4, 0x73, 0x5c, 0x60, 0xdc, 0xe5, 0x13, 0x26, 0x5c, 0xdf, 0xdd, 0x58,
	0x37, 0xff, 0x39, 0xba, 0x29, 0x3a, 0xb7, 0x85, 0xc6, 0x99, 0x69, 0xab, 0xef, 0xb0, 0xdc, 0x9a,
	0xf4, 0x46, 0x81, 0x97, 0xb4, 0xd0, 0xb1, 0xcc, 0xc0, 0x8b, 0x36, 0x5e, 0xbc, 0x1c, 0xd6, 0xd3,
	0xb9, 0xbb, 0x92, 0x33, 0xa7, 0x14, 0x0d, 0x17, 0x99, 0x37, 0x08, 0x09, 0xa5, 0xe9, 0xf0, 0x5d,
	0xc9, 0xc9, 0x88, 0x57, 0xa5, 0xab, 0x2f, 0x15, 0x74, 0xf5, 0xb5, 0x82, 0x36, 0xf3, 0x38, 0x37,
	0x84, 0x69, 0xf5, 0x7d, 0xba, 0x0a, 0x4b, 0x0e, 0xca, 0xe0, 0xe3, 0x04, 0x42, 0x0f, 0x66, 0x2f,
	0x71, 0x8d, 0x6f, 0x1c, 0x34, 0xdd, 0x26, 0x3b, 0xe8, 0x3a, 0x96, 0x59, 0xe0, 0x87, 0x2e, 0x9f,
	0x50, 0x10, 0xfb, 0x2c, 0x39, 0x73, 0xe2, 0xf1, 0x0f, 0x84, 0xcb, 0x37, 0x96, 0x51, 0x9e, 0x62,
	0xf5, 0xe0, 0x6d, 0xc7, 0x76, 0xba, 0xed, 0x4e, 0xa3, 0xb3, 0xdf, 0xee, 0xee, 0xef, 0xb5, 0x5b,
	0xf6, 0x56, 0xf3, 0x75, 0xd3, 0xde, 0x5e, 0x91, 0xb4, 0xe5, 0xe3, 0x53, 0xa3, 0xbc, 0x1f, 0xb2,
	0x08, 0xbc, 0xe0, 0x30, 0x80, 0xfe, 0x7f, 0xe4, 0x8d, 0xad, 0x4e, 0xf3, 0xa0, 0xd1, 0xb1, 0xb7,
	0x57, 0x50, 0x26, 0x77, 0x3d, 0x1e, 0xc4, 0x2e, 0x87, 0xbe, 0xf2, 0x10, 0xaf, 0xde, 0x92, 0xb7,
	0xec, 0xbd, 0xed, 0xe6, 0xde, 0xce, 0xca, 0x82, 0x56, 0x3e, 0x3e, 0x35, 0x8a, 0x2d, 0x08, 0xfb,
	0x41, 0xe8, 0x2b, 0x8f, 0xf0, 0xda, 0x2d, 0xd9, 0xbc, 0x67, 0x4e, 0xbb, 0x73, 0x7c, 0x6a, 0xc8,
	0x8d, 0xac, 0xa3, 0xb6, 0xf8, 0xf9, 0x9b, 0x2e, 0x6d, 0xda, 0x67, 0x17, 0x3a, 0x3a, 0xbf, 0xd0,
	0xd1, 0x9f, 0x0b, 0x1d, 0x9d, 0x5c, 0xea, 0xd2, 0xf9, 0xa5, 0x2e, 0xfd, 0xbc, 0xd4, 0xa5, 0x0f,
	0x4f, 0xfc, 0x80, 0x0f, 0x26, 0x3d, 0xd3, 0x23, 0x63, 0x2b, 0x79, 0xc4, 0x10, 0xf8, 0x27, 0x42,
	0x87, 0x22, 0xb6, 0x8e, 0xae, 0x7f, 0x33, 0x3e, 0x8d, 0x80, 0xf5, 0x0a, 0xe2, 0xeb, 0x7a, 0xf6,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x41, 0x4b, 0x5f, 0x83, 0x03, 0x00, 0x00,
}

func (this *PublicKey) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*PublicKey)
	if !ok {
		that2, ok := that.(PublicKey)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if that1.Key == nil {
		if this.Key != nil {
			return 1
		}
	} else if this.Key == nil {
		return -1
	} else {
		thisType := -1
		switch this.Key.(type) {
		case *PublicKey_Secp256K1:
			thisType = 0
		case *PublicKey_Schnorr:
			thisType = 1
		default:
			panic(fmt.Sprintf("compare: unexpected type %T in oneof", this.Key))
		}
		that1Type := -1
		switch that1.Key.(type) {
		case *PublicKey_Secp256K1:
			that1Type = 0
		case *PublicKey_Schnorr:
			that1Type = 1
		default:
			panic(fmt.Sprintf("compare: unexpected type %T in oneof", that1.Key))
		}
		if thisType == that1Type {
			if c := this.Key.Compare(that1.Key); c != 0 {
				return c
			}
		} else if thisType < that1Type {
			return -1
		} else if thisType > that1Type {
			return 1
		}
	}
	return 0
}
func (this *PublicKey_Secp256K1) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*PublicKey_Secp256K1)
	if !ok {
		that2, ok := that.(PublicKey_Secp256K1)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := bytes.Compare(this.Secp256K1, that1.Secp256K1); c != 0 {
		return c
	}
	return 0
}
func (this *PublicKey_Schnorr) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*PublicKey_Schnorr)
	if !ok {
		that2, ok := that.(PublicKey_Schnorr)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := bytes.Compare(this.Schnorr, that1.Schnorr); c != 0 {
		return c
	}
	return 0
}
func (this *PublicKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PublicKey)
	if !ok {
		that2, ok := that.(PublicKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Key == nil {
		if this.Key != nil {
			return false
		}
	} else if this.Key == nil {
		return false
	} else if !this.Key.Equal(that1.Key) {
		return false
	}
	return true
}
func (this *PublicKey_Secp256K1) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PublicKey_Secp256K1)
	if !ok {
		that2, ok := that.(PublicKey_Secp256K1)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Secp256K1, that1.Secp256K1) {
		return false
	}
	return true
}
func (this *PublicKey_Schnorr) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PublicKey_Schnorr)
	if !ok {
		that2, ok := that.(PublicKey_Schnorr)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Schnorr, that1.Schnorr) {
		return false
	}
	return true
}
func (m *Relayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastElected, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastElected):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintRelayer(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintRelayer(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Threshold != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x10
	}
	if m.Version != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Voter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Voter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Voter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.VoteKey) > 0 {
		i -= len(m.VoteKey)
		copy(dAtA[i:], m.VoteKey)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.VoteKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxKey) > 0 {
		i -= len(m.TxKey)
		copy(dAtA[i:], m.TxKey)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.TxKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PublicKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		{
			size := m.Key.Size()
			i -= size
			if _, err := m.Key.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PublicKey_Secp256K1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Secp256K1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Secp256K1 != nil {
		i -= len(m.Secp256K1)
		copy(dAtA[i:], m.Secp256K1)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Secp256K1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Schnorr) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Schnorr) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Schnorr != nil {
		i -= len(m.Schnorr)
		copy(dAtA[i:], m.Schnorr)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Schnorr)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Votes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Votes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Votes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Voters) > 0 {
		i -= len(m.Voters)
		copy(dAtA[i:], m.Voters)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Voters)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRelayer(dAtA []byte, offset int, v uint64) int {
	offset -= sovRelayer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Relayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovRelayer(uint64(m.Version))
	}
	if m.Threshold != 0 {
		n += 1 + sovRelayer(uint64(m.Threshold))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	if len(m.Voters) > 0 {
		for _, s := range m.Voters {
			l = len(s)
			n += 1 + l + sovRelayer(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastElected)
	n += 1 + l + sovRelayer(uint64(l))
	return n
}

func (m *Voter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxKey)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.VoteKey)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovRelayer(uint64(m.Status))
	}
	return n
}

func (m *PublicKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		n += m.Key.Size()
	}
	return n
}

func (m *PublicKey_Secp256K1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Secp256K1 != nil {
		l = len(m.Secp256K1)
		n += 1 + l + sovRelayer(uint64(l))
	}
	return n
}
func (m *PublicKey_Schnorr) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Schnorr != nil {
		l = len(m.Schnorr)
		n += 1 + l + sovRelayer(uint64(l))
	}
	return n
}
func (m *Votes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovRelayer(uint64(m.Sequence))
	}
	l = len(m.Voters)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	return n
}

func sovRelayer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRelayer(x uint64) (n int) {
	return sovRelayer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Relayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Relayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastElected", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastElected, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Voter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Voter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Voter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxKey = append(m.TxKey[:0], dAtA[iNdEx:postIndex]...)
			if m.TxKey == nil {
				m.TxKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteKey = append(m.VoteKey[:0], dAtA[iNdEx:postIndex]...)
			if m.VoteKey == nil {
				m.VoteKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= VoterStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PublicKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PublicKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secp256K1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Key = &PublicKey_Secp256K1{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schnorr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Key = &PublicKey_Schnorr{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Votes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Votes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Votes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters[:0], dAtA[iNdEx:postIndex]...)
			if m.Voters == nil {
				m.Voters = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRelayer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRelayer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRelayer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRelayer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRelayer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRelayer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRelayer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRelayer = fmt.Errorf("proto: unexpected end of group")
)
