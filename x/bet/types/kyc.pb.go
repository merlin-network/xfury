// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/bet/kyc.proto

package types

import (
	fmt "fmt"
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

// KycDataPayload is the KYC info.
type KycDataPayload struct {
	// ignore is true if kyc validation is not required.
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// approved represent if kyc validation is approved or not.
	Approved bool `protobuf:"varint,2,opt,name=approved,proto3" json:"approved,omitempty"`
	// id is the id of the kyc user.
	ID string `protobuf:"bytes,3,opt,name=id,proto3" json:"id"`
}

func (m *KycDataPayload) Reset()         { *m = KycDataPayload{} }
func (m *KycDataPayload) String() string { return proto.CompactTextString(m) }
func (*KycDataPayload) ProtoMessage()    {}
func (*KycDataPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_13abc80c3002dbc4, []int{0}
}
func (m *KycDataPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KycDataPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KycDataPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KycDataPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KycDataPayload.Merge(m, src)
}
func (m *KycDataPayload) XXX_Size() int {
	return m.Size()
}
func (m *KycDataPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_KycDataPayload.DiscardUnknown(m)
}

var xxx_messageInfo_KycDataPayload proto.InternalMessageInfo

func (m *KycDataPayload) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *KycDataPayload) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

func (m *KycDataPayload) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*KycDataPayload)(nil), "furynetwork.fury.bet.KycDataPayload")
}

func init() { proto.RegisterFile("fury/bet/kyc.proto", fileDescriptor_13abc80c3002dbc4) }

var fileDescriptor_13abc80c3002dbc4 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4e, 0x4f, 0xd5,
	0x4f, 0x4a, 0x2d, 0xd1, 0xcf, 0xae, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a,
	0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0x4f, 0xd5, 0x4b, 0x4a,
	0x2d, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x4a, 0x49,
	0x5c, 0x7c, 0xde, 0x95, 0xc9, 0x2e, 0x89, 0x25, 0x89, 0x01, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29,
	0x42, 0x62, 0x5c, 0x6c, 0x99, 0xe9, 0x79, 0xf9, 0x45, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c,
	0x41, 0x50, 0x9e, 0x90, 0x14, 0x17, 0x47, 0x62, 0x41, 0x41, 0x51, 0x7e, 0x59, 0x6a, 0x8a, 0x04,
	0x13, 0x58, 0x06, 0xce, 0x17, 0x92, 0xe1, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x56, 0x60, 0xd4, 0xe0,
	0x74, 0xe2, 0x79, 0x74, 0x4f, 0x9e, 0xc9, 0xd3, 0xe5, 0xd5, 0x3d, 0x79, 0xa6, 0xcc, 0x94, 0x20,
	0xa6, 0xcc, 0x14, 0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x52,
	0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x4e, 0x4f, 0xd5, 0x85,
	0xba, 0x19, 0xc4, 0xd6, 0xaf, 0x00, 0xfb, 0xa9, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec,
	0x58, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x34, 0x34, 0x32, 0xeb, 0x00, 0x00, 0x00,
}

func (m *KycDataPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KycDataPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KycDataPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintKyc(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Approved {
		i--
		if m.Approved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Ignore {
		i--
		if m.Ignore {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKyc(dAtA []byte, offset int, v uint64) int {
	offset -= sovKyc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KycDataPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ignore {
		n += 2
	}
	if m.Approved {
		n += 2
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovKyc(uint64(l))
	}
	return n
}

func sovKyc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKyc(x uint64) (n int) {
	return sovKyc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KycDataPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKyc
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
			return fmt.Errorf("proto: KycDataPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KycDataPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ignore", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKyc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ignore = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKyc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Approved = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKyc
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
				return ErrInvalidLengthKyc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKyc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKyc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKyc
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
func skipKyc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKyc
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
					return 0, ErrIntOverflowKyc
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
					return 0, ErrIntOverflowKyc
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
				return 0, ErrInvalidLengthKyc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKyc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKyc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKyc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKyc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKyc = fmt.Errorf("proto: unexpected end of group")
)
