// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sequencers/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Sequencer struct {
	// Validator is a convenient storage for e.g operator address and consensus pub key
	Validator *types.Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	// RewardAddr is the sdk acc address where the sequencer has opted to receive rewards. Empty if not set.
	RewardAddr string `protobuf:"bytes,2,opt,name=reward_addr,json=rewardAddr,proto3" json:"reward_addr,omitempty"`
}

func (m *Sequencer) Reset()         { *m = Sequencer{} }
func (m *Sequencer) String() string { return proto.CompactTextString(m) }
func (*Sequencer) ProtoMessage()    {}
func (*Sequencer) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ef324f83e482d4, []int{0}
}
func (m *Sequencer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sequencer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequencer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sequencer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequencer.Merge(m, src)
}
func (m *Sequencer) XXX_Size() int {
	return m.Size()
}
func (m *Sequencer) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequencer.DiscardUnknown(m)
}

var xxx_messageInfo_Sequencer proto.InternalMessageInfo

func (m *Sequencer) GetValidator() *types.Validator {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *Sequencer) GetRewardAddr() string {
	if m != nil {
		return m.RewardAddr
	}
	return ""
}

// GenesisState defines the module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// Sequencers all stored sequencers
	Sequencers []Sequencer `protobuf:"bytes,3,rep,name=sequencers,proto3" json:"sequencers"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ef324f83e482d4, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSequencers() []Sequencer {
	if m != nil {
		return m.Sequencers
	}
	return nil
}

func init() {
	proto.RegisterType((*Sequencer)(nil), "rollapp.sequencers.types.Sequencer")
	proto.RegisterType((*GenesisState)(nil), "rollapp.sequencers.types.GenesisState")
}

func init() { proto.RegisterFile("sequencers/genesis.proto", fileDescriptor_56ef324f83e482d4) }

var fileDescriptor_56ef324f83e482d4 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0xe3, 0xb6, 0xaa, 0xfe, 0xba, 0xff, 0x80, 0x22, 0x24, 0xa2, 0x0e, 0x6e, 0x28, 0x0c,
	0x5d, 0xb0, 0xd5, 0xb2, 0x30, 0x81, 0xe8, 0x82, 0x60, 0x42, 0xa9, 0xc4, 0xc0, 0x82, 0xdc, 0xda,
	0x0a, 0x51, 0x9b, 0x38, 0xd8, 0x6e, 0x69, 0x38, 0x05, 0x37, 0xe0, 0x3a, 0x1d, 0x3b, 0x32, 0x21,
	0x94, 0x5c, 0x04, 0x11, 0x27, 0x24, 0x4b, 0xb7, 0x7c, 0xd1, 0xf3, 0x3e, 0x9f, 0x5f, 0x1b, 0x3a,
	0x8a, 0xbf, 0xac, 0x78, 0x34, 0xe7, 0x52, 0x11, 0x9f, 0x47, 0x5c, 0x05, 0x0a, 0xc7, 0x52, 0x68,
	0x61, 0x3b, 0x52, 0x2c, 0x97, 0x34, 0x8e, 0x71, 0x45, 0x60, 0x9d, 0xc4, 0x5c, 0xf5, 0x0e, 0x7d,
	0xe1, 0x8b, 0x1c, 0x22, 0xbf, 0x5f, 0x86, 0xef, 0x1d, 0xd5, 0x4c, 0x31, 0x95, 0x34, 0x2c, 0x44,
	0xbd, 0xd3, 0xb9, 0x50, 0xa1, 0x50, 0x44, 0x69, 0xba, 0x08, 0x22, 0x9f, 0xac, 0x47, 0x33, 0xae,
	0xe9, 0xa8, 0x9c, 0x0d, 0x35, 0x08, 0x61, 0x67, 0x5a, 0x0a, 0xec, 0x2b, 0xd8, 0x59, 0xd3, 0x65,
	0xc0, 0xa8, 0x16, 0xd2, 0x01, 0x2e, 0x18, 0x76, 0xc7, 0xc7, 0xd8, 0x68, 0x70, 0x19, 0x2b, 0x34,
	0xf8, 0xa1, 0x04, 0xbd, 0x2a, 0x63, 0xf7, 0x61, 0x57, 0xf2, 0x57, 0x2a, 0xd9, 0x13, 0x65, 0x4c,
	0x3a, 0x0d, 0x17, 0x0c, 0x3b, 0x1e, 0x34, 0xbf, 0xae, 0x19, 0x93, 0x83, 0x0f, 0x00, 0xff, 0xdf,
	0x98, 0xbe, 0x53, 0x4d, 0x35, 0xb7, 0x2f, 0x61, 0xdb, 0x9c, 0xba, 0xd8, 0xe7, 0xe2, 0x7d, 0xfd,
	0xf1, 0x7d, 0xce, 0x4d, 0x5a, 0xdb, 0xaf, 0xbe, 0xe5, 0x15, 0x29, 0xfb, 0x16, 0xc2, 0x0a, 0x74,
	0x9a, 0x6e, 0x73, 0xd8, 0x1d, 0x9f, 0xec, 0x77, 0xfc, 0x75, 0x2d, 0x34, 0xb5, 0xf0, 0x5d, 0xeb,
	0x5f, 0xe3, 0xa0, 0x39, 0xf1, 0xb6, 0x29, 0x02, 0xbb, 0x14, 0x81, 0xef, 0x14, 0x81, 0xf7, 0x0c,
	0x59, 0xbb, 0x0c, 0x59, 0x9f, 0x19, 0xb2, 0x1e, 0x2f, 0xfc, 0x40, 0x3f, 0xaf, 0x66, 0x78, 0x2e,
	0x42, 0xc2, 0x92, 0x90, 0x47, 0x2a, 0x10, 0xd1, 0x26, 0x79, 0xab, 0x86, 0x33, 0xc9, 0x16, 0x64,
	0x43, 0x6a, 0x2f, 0x92, 0x6f, 0x9d, 0xb5, 0xf3, 0xbb, 0x3e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x89, 0x24, 0x98, 0xf6, 0x01, 0x00, 0x00,
}

func (m *Sequencer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequencer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sequencer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardAddr) > 0 {
		i -= len(m.RewardAddr)
		copy(dAtA[i:], m.RewardAddr)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.RewardAddr)))
		i--
		dAtA[i] = 0x12
	}
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sequencers) > 0 {
		for iNdEx := len(m.Sequencers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sequencers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sequencer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.RewardAddr)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Sequencers) > 0 {
		for _, e := range m.Sequencers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sequencer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Sequencer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequencer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &types.Validator{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequencers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequencers = append(m.Sequencers, Sequencer{})
			if err := m.Sequencers[len(m.Sequencers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
