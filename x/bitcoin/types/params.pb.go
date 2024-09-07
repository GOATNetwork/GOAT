// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: goat/bitcoin/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ChainConfig
type ChainConfig struct {
	NetworkName          string `protobuf:"bytes,1,opt,name=network_name,json=networkName,proto3" json:"network_name,omitempty"`
	PubkeyHashAddrPrefix uint32 `protobuf:"varint,2,opt,name=pubkey_hash_addr_prefix,json=pubkeyHashAddrPrefix,proto3" json:"pubkey_hash_addr_prefix,omitempty"`
	ScriptHashAddrPrefix uint32 `protobuf:"varint,3,opt,name=script_hash_addr_prefix,json=scriptHashAddrPrefix,proto3" json:"script_hash_addr_prefix,omitempty"`
	Bech32Hrp            string `protobuf:"bytes,4,opt,name=bech32_hrp,json=bech32Hrp,proto3" json:"bech32_hrp,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_96fcb82db75b9cd1, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func (m *ChainConfig) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

func (m *ChainConfig) GetPubkeyHashAddrPrefix() uint32 {
	if m != nil {
		return m.PubkeyHashAddrPrefix
	}
	return 0
}

func (m *ChainConfig) GetScriptHashAddrPrefix() uint32 {
	if m != nil {
		return m.ScriptHashAddrPrefix
	}
	return 0
}

func (m *ChainConfig) GetBech32Hrp() string {
	if m != nil {
		return m.Bech32Hrp
	}
	return ""
}

// Params defines the parameters for the module.
type Params struct {
	ChainConfig           *ChainConfig `protobuf:"bytes,1,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config,omitempty"`
	SafeConfirmationBlock uint32       `protobuf:"varint,2,opt,name=safe_confirmation_block,json=safeConfirmationBlock,proto3" json:"safe_confirmation_block,omitempty"`
	HardConfirmationBlock uint32       `protobuf:"varint,3,opt,name=hard_confirmation_block,json=hardConfirmationBlock,proto3" json:"hard_confirmation_block,omitempty"`
	MinDepositAmount      uint64       `protobuf:"varint,4,opt,name=min_deposit_amount,json=minDepositAmount,proto3" json:"min_deposit_amount,omitempty"`
	DepositMagicPrefix    []byte       `protobuf:"bytes,5,opt,name=deposit_magic_prefix,json=depositMagicPrefix,proto3" json:"deposit_magic_prefix,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_96fcb82db75b9cd1, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetChainConfig() *ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return nil
}

func (m *Params) GetSafeConfirmationBlock() uint32 {
	if m != nil {
		return m.SafeConfirmationBlock
	}
	return 0
}

func (m *Params) GetHardConfirmationBlock() uint32 {
	if m != nil {
		return m.HardConfirmationBlock
	}
	return 0
}

func (m *Params) GetMinDepositAmount() uint64 {
	if m != nil {
		return m.MinDepositAmount
	}
	return 0
}

func (m *Params) GetDepositMagicPrefix() []byte {
	if m != nil {
		return m.DepositMagicPrefix
	}
	return nil
}

func init() {
	proto.RegisterType((*ChainConfig)(nil), "goat.bitcoin.v1.ChainConfig")
	proto.RegisterType((*Params)(nil), "goat.bitcoin.v1.Params")
}

func init() { proto.RegisterFile("goat/bitcoin/v1/params.proto", fileDescriptor_96fcb82db75b9cd1) }

var fileDescriptor_96fcb82db75b9cd1 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x6f, 0x13, 0x31,
	0x18, 0xc6, 0xe3, 0x52, 0x2a, 0xd5, 0x17, 0x04, 0x9c, 0x1a, 0x35, 0xaa, 0xca, 0x29, 0x74, 0x8a,
	0x00, 0xdd, 0xd1, 0x56, 0x30, 0xb0, 0xa0, 0x36, 0x20, 0x65, 0x01, 0x45, 0x37, 0xb2, 0x58, 0x3e,
	0x9f, 0x73, 0xb6, 0x82, 0xff, 0xc8, 0xe7, 0x84, 0xe4, 0x2b, 0x30, 0xf1, 0x51, 0x58, 0x18, 0xd9,
	0x19, 0x33, 0x32, 0xa2, 0x64, 0xe0, 0x6b, 0x20, 0xdb, 0x17, 0x14, 0x25, 0x5d, 0x6e, 0x78, 0x7f,
	0xfe, 0x9d, 0xfc, 0x3c, 0x7e, 0xe1, 0x79, 0xa5, 0xb0, 0xcd, 0x0a, 0x6e, 0x89, 0xe2, 0x32, 0x9b,
	0x5d, 0x66, 0x1a, 0x1b, 0x2c, 0xea, 0x54, 0x1b, 0x65, 0x55, 0xfc, 0xd0, 0xd1, 0xb4, 0xa1, 0xe9,
	0xec, 0xf2, 0xec, 0x31, 0x16, 0x5c, 0xaa, 0xcc, 0x7f, 0xc3, 0x99, 0x8b, 0x9f, 0x00, 0x46, 0x03,
	0x86, 0xb9, 0x1c, 0x28, 0x39, 0xe6, 0x55, 0xfc, 0x14, 0xb6, 0x25, 0xb5, 0x5f, 0x94, 0x99, 0x20,
	0x89, 0x05, 0xed, 0x82, 0x1e, 0xe8, 0x1f, 0xe7, 0x51, 0x33, 0xfb, 0x88, 0x05, 0x8d, 0x5f, 0xc1,
	0x53, 0x3d, 0x2d, 0x26, 0x74, 0x81, 0x18, 0xae, 0x19, 0xc2, 0x65, 0x69, 0x90, 0x36, 0x74, 0xcc,
	0xe7, 0xdd, 0x83, 0x1e, 0xe8, 0x3f, 0xc8, 0x4f, 0x02, 0x1e, 0xe2, 0x9a, 0xdd, 0x94, 0xa5, 0x19,
	0x79, 0xe6, 0xb4, 0x9a, 0x18, 0xae, 0xed, 0xbe, 0x76, 0x2f, 0x68, 0x01, 0xef, 0x68, 0x4f, 0x20,
	0x2c, 0x28, 0x61, 0xd7, 0x57, 0x88, 0x19, 0xdd, 0x3d, 0xf4, 0xd7, 0x39, 0x0e, 0x93, 0xa1, 0xd1,
	0x17, 0x3f, 0x0e, 0xe0, 0xd1, 0xc8, 0x87, 0x8e, 0xdf, 0xc2, 0x36, 0x71, 0x49, 0x10, 0xf1, 0x51,
	0xfc, 0xd5, 0xa3, 0xab, 0xf3, 0x74, 0xa7, 0x85, 0x74, 0x2b, 0x6e, 0x1e, 0x91, 0xad, 0xec, 0xaf,
	0xe1, 0x69, 0x8d, 0xc7, 0x34, 0xf8, 0x46, 0x60, 0xcb, 0x95, 0x44, 0xc5, 0x67, 0x45, 0x26, 0x4d,
	0xb0, 0x8e, 0xc3, 0x83, 0x2d, 0x7a, 0xeb, 0xa0, 0xf3, 0x18, 0x36, 0xe5, 0x5d, 0x5e, 0x48, 0xd6,
	0x71, 0x78, 0xdf, 0x7b, 0x01, 0x63, 0xc1, 0x25, 0x2a, 0xa9, 0x56, 0x35, 0xb7, 0x08, 0x0b, 0x35,
	0x95, 0xd6, 0x47, 0x3c, 0xcc, 0x1f, 0x09, 0x2e, 0xdf, 0x05, 0x70, 0xe3, 0xe7, 0xf1, 0x4b, 0x78,
	0xb2, 0x39, 0x29, 0x70, 0xc5, 0xc9, 0xa6, 0xbc, 0xfb, 0x3d, 0xd0, 0x6f, 0xe7, 0x71, 0xc3, 0x3e,
	0x38, 0x14, 0xaa, 0x7b, 0x73, 0xf6, 0xf5, 0xef, 0xf7, 0x67, 0x1d, 0xbf, 0x22, 0xf3, 0xff, 0x4b,
	0x12, 0xca, 0xba, 0x7d, 0xff, 0x6b, 0x95, 0x80, 0xe5, 0x2a, 0x01, 0x7f, 0x56, 0x09, 0xf8, 0xb6,
	0x4e, 0x5a, 0xcb, 0x75, 0xd2, 0xfa, 0xbd, 0x4e, 0x5a, 0x9f, 0x9e, 0x57, 0xdc, 0xb2, 0x69, 0x91,
	0x12, 0x25, 0x32, 0xe7, 0x36, 0x4f, 0x9f, 0xed, 0xfc, 0xc7, 0x2e, 0x34, 0xad, 0x8b, 0x23, 0xbf,
	0x45, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xa8, 0x25, 0xe1, 0x89, 0x02, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bech32Hrp) > 0 {
		i -= len(m.Bech32Hrp)
		copy(dAtA[i:], m.Bech32Hrp)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Bech32Hrp)))
		i--
		dAtA[i] = 0x22
	}
	if m.ScriptHashAddrPrefix != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ScriptHashAddrPrefix))
		i--
		dAtA[i] = 0x18
	}
	if m.PubkeyHashAddrPrefix != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PubkeyHashAddrPrefix))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NetworkName) > 0 {
		i -= len(m.NetworkName)
		copy(dAtA[i:], m.NetworkName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.NetworkName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DepositMagicPrefix) > 0 {
		i -= len(m.DepositMagicPrefix)
		copy(dAtA[i:], m.DepositMagicPrefix)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DepositMagicPrefix)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MinDepositAmount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinDepositAmount))
		i--
		dAtA[i] = 0x20
	}
	if m.HardConfirmationBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HardConfirmationBlock))
		i--
		dAtA[i] = 0x18
	}
	if m.SafeConfirmationBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafeConfirmationBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainConfig != nil {
		{
			size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NetworkName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.PubkeyHashAddrPrefix != 0 {
		n += 1 + sovParams(uint64(m.PubkeyHashAddrPrefix))
	}
	if m.ScriptHashAddrPrefix != 0 {
		n += 1 + sovParams(uint64(m.ScriptHashAddrPrefix))
	}
	l = len(m.Bech32Hrp)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainConfig != nil {
		l = m.ChainConfig.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.SafeConfirmationBlock != 0 {
		n += 1 + sovParams(uint64(m.SafeConfirmationBlock))
	}
	if m.HardConfirmationBlock != 0 {
		n += 1 + sovParams(uint64(m.HardConfirmationBlock))
	}
	if m.MinDepositAmount != 0 {
		n += 1 + sovParams(uint64(m.MinDepositAmount))
	}
	l = len(m.DepositMagicPrefix)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubkeyHashAddrPrefix", wireType)
			}
			m.PubkeyHashAddrPrefix = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PubkeyHashAddrPrefix |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptHashAddrPrefix", wireType)
			}
			m.ScriptHashAddrPrefix = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScriptHashAddrPrefix |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32Hrp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32Hrp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainConfig == nil {
				m.ChainConfig = &ChainConfig{}
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafeConfirmationBlock", wireType)
			}
			m.SafeConfirmationBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafeConfirmationBlock |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HardConfirmationBlock", wireType)
			}
			m.HardConfirmationBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HardConfirmationBlock |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDepositAmount", wireType)
			}
			m.MinDepositAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDepositAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositMagicPrefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositMagicPrefix = append(m.DepositMagicPrefix[:0], dAtA[iNdEx:postIndex]...)
			if m.DepositMagicPrefix == nil {
				m.DepositMagicPrefix = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
