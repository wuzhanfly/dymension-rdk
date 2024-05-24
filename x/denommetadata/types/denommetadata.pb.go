// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: denommetadata/denommetadata.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
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

// DenomMetdata defines the metadata and the denom trace for the denom.
type DenomMetadata struct {
	TokenMetadata types.Metadata `protobuf:"bytes,1,opt,name=token_metadata,json=tokenMetadata,proto3" json:"token_metadata"`
	DenomTrace    string         `protobuf:"bytes,2,opt,name=denom_trace,json=denomTrace,proto3" json:"denom_trace,omitempty"`
}

func (m *DenomMetadata) Reset()         { *m = DenomMetadata{} }
func (m *DenomMetadata) String() string { return proto.CompactTextString(m) }
func (*DenomMetadata) ProtoMessage()    {}
func (*DenomMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a11037efbe4195, []int{0}
}
func (m *DenomMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomMetadata.Merge(m, src)
}
func (m *DenomMetadata) XXX_Size() int {
	return m.Size()
}
func (m *DenomMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DenomMetadata proto.InternalMessageInfo

func (m *DenomMetadata) GetTokenMetadata() types.Metadata {
	if m != nil {
		return m.TokenMetadata
	}
	return types.Metadata{}
}

func (m *DenomMetadata) GetDenomTrace() string {
	if m != nil {
		return m.DenomTrace
	}
	return ""
}

func init() {
	proto.RegisterType((*DenomMetadata)(nil), "rollapp.denommetadata.types.DenomMetadata")
}

func init() { proto.RegisterFile("denommetadata/denommetadata.proto", fileDescriptor_74a11037efbe4195) }

var fileDescriptor_74a11037efbe4195 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x49, 0xcd, 0xcb,
	0xcf, 0xcd, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x47, 0xe1, 0xe9, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x49, 0x17, 0xe5, 0xe7, 0xe4, 0x24, 0x16, 0x14, 0xe8, 0xa1, 0x4a, 0x96, 0x54,
	0x16, 0xa4, 0x16, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xd5, 0xe9, 0x83, 0x58, 0x10, 0x2d,
	0x52, 0x72, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xfa, 0x49, 0x89, 0x79, 0xd9, 0xfa, 0x65, 0x86,
	0x49, 0xa9, 0x25, 0x89, 0x86, 0x60, 0x0e, 0x44, 0x5e, 0xa9, 0x86, 0x8b, 0xd7, 0x05, 0x64, 0x98,
	0x2f, 0xd4, 0x30, 0x21, 0x2f, 0x2e, 0xbe, 0x92, 0xfc, 0xec, 0xd4, 0xbc, 0x78, 0x98, 0xf1, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xb2, 0x7a, 0x10, 0x93, 0xf4, 0xc0, 0x9a, 0xa1, 0x26, 0xe9,
	0xc1, 0xb4, 0x39, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0xc4, 0x0b, 0xd6, 0x0a, 0x37, 0x4b, 0x9e,
	0x8b, 0x1b, 0xec, 0xd2, 0xf8, 0x92, 0xa2, 0xc4, 0xe4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x2e, 0xb0, 0x50, 0x08, 0x48, 0xc4, 0x29, 0xf4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0xac, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x53,
	0x2a, 0x73, 0x53, 0xf3, 0x8a, 0x33, 0xf3, 0xf3, 0x2a, 0x2a, 0xab, 0x10, 0x1c, 0xdd, 0xa2, 0x94,
	0x6c, 0xfd, 0x0a, 0xd4, 0x70, 0xd2, 0x07, 0x07, 0x45, 0x12, 0x1b, 0xd8, 0x6f, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x87, 0x70, 0x47, 0x14, 0x53, 0x01, 0x00, 0x00,
}

func (m *DenomMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomTrace) > 0 {
		i -= len(m.DenomTrace)
		copy(dAtA[i:], m.DenomTrace)
		i = encodeVarintDenommetadata(dAtA, i, uint64(len(m.DenomTrace)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.TokenMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDenommetadata(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDenommetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenommetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenMetadata.Size()
	n += 1 + l + sovDenommetadata(uint64(l))
	l = len(m.DenomTrace)
	if l > 0 {
		n += 1 + l + sovDenommetadata(uint64(l))
	}
	return n
}

func sovDenommetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenommetadata(x uint64) (n int) {
	return sovDenommetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenommetadata
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
			return fmt.Errorf("proto: DenomMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenommetadata
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
				return ErrInvalidLengthDenommetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDenommetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTrace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenommetadata
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
				return ErrInvalidLengthDenommetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenommetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomTrace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenommetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenommetadata
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
func skipDenommetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenommetadata
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
					return 0, ErrIntOverflowDenommetadata
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
					return 0, ErrIntOverflowDenommetadata
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
				return 0, ErrInvalidLengthDenommetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenommetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenommetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenommetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenommetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenommetadata = fmt.Errorf("proto: unexpected end of group")
)
