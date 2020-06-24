// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consul.proto

package consul

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ServiceConsul struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty" form:"ip" validate:"required"`
	Port                 int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty" form:"port" validate:"required"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	CheckName            string   `protobuf:"bytes,4,opt,name=check_name,json=checkName,proto3" json:"check_name,omitempty" form:"check_name" validate:"required"`
	ServerName           string   `protobuf:"bytes,5,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty" form:"server_name" validate:"required"`
	ServerType           string   `protobuf:"bytes,6,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty" form:"server_type" validate:"required"`
	Interval             int64    `protobuf:"varint,7,opt,name=interval,proto3" json:"interval,omitempty"`
	Deregister           int64    `protobuf:"varint,8,opt,name=deregister,proto3" json:"deregister,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceConsul) Reset()         { *m = ServiceConsul{} }
func (m *ServiceConsul) String() string { return proto.CompactTextString(m) }
func (*ServiceConsul) ProtoMessage()    {}
func (*ServiceConsul) Descriptor() ([]byte, []int) {
	return fileDescriptor_665b1cce082a77c9, []int{0}
}
func (m *ServiceConsul) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceConsul) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceConsul.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceConsul) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceConsul.Merge(m, src)
}
func (m *ServiceConsul) XXX_Size() int {
	return m.Size()
}
func (m *ServiceConsul) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceConsul.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceConsul proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ServiceConsul)(nil), "consul.ServiceConsul")
}

func init() { proto.RegisterFile("consul.proto", fileDescriptor_665b1cce082a77c9) }

var fileDescriptor_665b1cce082a77c9 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x9b, 0xa6, 0xd6, 0x76, 0xd4, 0xcd, 0xac, 0x86, 0xa2, 0x49, 0x8c, 0x28, 0x55, 0xb0,
	0x45, 0x5c, 0x08, 0x5d, 0xd6, 0x8d, 0x20, 0xb8, 0xa8, 0xae, 0xdc, 0x48, 0x9a, 0xbe, 0xa6, 0x83,
	0x4d, 0x67, 0x9c, 0x4c, 0x02, 0xbd, 0x89, 0xb7, 0xf0, 0x1a, 0x5d, 0x7a, 0x82, 0xa0, 0xf5, 0x06,
	0x39, 0x81, 0xf4, 0x8d, 0x68, 0x91, 0xe8, 0xee, 0xfd, 0xef, 0x7f, 0xff, 0xc7, 0x30, 0x3f, 0xd9,
	0x0e, 0xc5, 0x2c, 0x49, 0xa7, 0x1d, 0xa9, 0x84, 0x16, 0xb4, 0x6e, 0x54, 0xeb, 0x34, 0xe2, 0x7a,
	0x92, 0x0e, 0x3b, 0xa1, 0x88, 0xbb, 0x91, 0x88, 0x44, 0x17, 0xed, 0x61, 0x3a, 0x46, 0x85, 0x02,
	0x27, 0x13, 0xf3, 0x5f, 0x6c, 0xb2, 0x73, 0x0b, 0x2a, 0xe3, 0x21, 0x5c, 0x22, 0x80, 0x9e, 0x91,
	0x2a, 0x97, 0xcc, 0xf2, 0xac, 0x76, 0xb3, 0xbf, 0x5f, 0xe4, 0xee, 0xde, 0x58, 0xa8, 0xb8, 0xe7,
	0x73, 0xe9, 0x7b, 0x59, 0x30, 0xe5, 0xa3, 0x40, 0x43, 0xcf, 0x57, 0xf0, 0x94, 0x72, 0x05, 0x23,
	0x7f, 0x50, 0xe5, 0x92, 0x5e, 0x90, 0x9a, 0x14, 0x4a, 0xb3, 0xaa, 0x67, 0xb5, 0xed, 0xfe, 0x41,
	0x91, 0xbb, 0xae, 0x09, 0xad, 0xb6, 0xe5, 0x31, 0x0c, 0x50, 0x4a, 0x6a, 0x3a, 0x88, 0x12, 0x66,
	0x7b, 0x76, 0xbb, 0x39, 0xc0, 0x99, 0x5e, 0x11, 0x12, 0x4e, 0x20, 0x7c, 0x7c, 0x98, 0x05, 0x31,
	0xb0, 0x1a, 0xbe, 0xe3, 0xb8, 0xc8, 0xdd, 0x43, 0x83, 0xfc, 0xf1, 0xca, 0xc1, 0x4d, 0x3c, 0xb8,
	0x09, 0x62, 0xa0, 0xd7, 0x64, 0x2b, 0x01, 0x95, 0x81, 0x32, 0xa8, 0x0d, 0x44, 0x9d, 0x14, 0xb9,
	0x7b, 0x64, 0x50, 0x6b, 0x66, 0x39, 0x8b, 0x98, 0x8b, 0x5f, 0x30, 0x3d, 0x97, 0xc0, 0xea, 0x7f,
	0xc0, 0x56, 0xe6, 0xbf, 0xb0, 0xbb, 0xb9, 0x04, 0xda, 0x22, 0x0d, 0x3e, 0xd3, 0xa0, 0xb2, 0x60,
	0xca, 0x36, 0x57, 0x9f, 0x36, 0xf8, 0xd6, 0xd4, 0x21, 0x64, 0x04, 0x0a, 0x22, 0x9e, 0x68, 0x50,
	0xac, 0x81, 0xee, 0xda, 0xa6, 0xbf, 0xbb, 0x78, 0x77, 0x2a, 0x8b, 0xa5, 0x63, 0xbd, 0x2e, 0x1d,
	0xeb, 0x6d, 0xe9, 0x58, 0xcf, 0x1f, 0x4e, 0xe5, 0xfe, 0xab, 0xfe, 0x61, 0x1d, 0x6b, 0x3d, 0xff,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xaa, 0x65, 0xf0, 0x1d, 0x02, 0x00, 0x00,
}

func (m *ServiceConsul) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceConsul) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceConsul) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Deregister != 0 {
		i = encodeVarintConsul(dAtA, i, uint64(m.Deregister))
		i--
		dAtA[i] = 0x40
	}
	if m.Interval != 0 {
		i = encodeVarintConsul(dAtA, i, uint64(m.Interval))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ServerType) > 0 {
		i -= len(m.ServerType)
		copy(dAtA[i:], m.ServerType)
		i = encodeVarintConsul(dAtA, i, uint64(len(m.ServerType)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ServerName) > 0 {
		i -= len(m.ServerName)
		copy(dAtA[i:], m.ServerName)
		i = encodeVarintConsul(dAtA, i, uint64(len(m.ServerName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CheckName) > 0 {
		i -= len(m.CheckName)
		copy(dAtA[i:], m.CheckName)
		i = encodeVarintConsul(dAtA, i, uint64(len(m.CheckName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintConsul(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Port != 0 {
		i = encodeVarintConsul(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Ip) > 0 {
		i -= len(m.Ip)
		copy(dAtA[i:], m.Ip)
		i = encodeVarintConsul(dAtA, i, uint64(len(m.Ip)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConsul(dAtA []byte, offset int, v uint64) int {
	offset -= sovConsul(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceConsul) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovConsul(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovConsul(uint64(m.Port))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovConsul(uint64(l))
		}
	}
	l = len(m.CheckName)
	if l > 0 {
		n += 1 + l + sovConsul(uint64(l))
	}
	l = len(m.ServerName)
	if l > 0 {
		n += 1 + l + sovConsul(uint64(l))
	}
	l = len(m.ServerType)
	if l > 0 {
		n += 1 + l + sovConsul(uint64(l))
	}
	if m.Interval != 0 {
		n += 1 + sovConsul(uint64(m.Interval))
	}
	if m.Deregister != 0 {
		n += 1 + sovConsul(uint64(m.Deregister))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConsul(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConsul(x uint64) (n int) {
	return sovConsul(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceConsul) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsul
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
			return fmt.Errorf("proto: ServiceConsul: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceConsul: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
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
				return ErrInvalidLengthConsul
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsul
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
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
				return ErrInvalidLengthConsul
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsul
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
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
				return ErrInvalidLengthConsul
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsul
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CheckName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
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
				return ErrInvalidLengthConsul
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsul
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
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
				return ErrInvalidLengthConsul
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsul
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			m.Interval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Interval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deregister", wireType)
			}
			m.Deregister = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsul
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deregister |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConsul(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsul
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConsul
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConsul(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsul
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
					return 0, ErrIntOverflowConsul
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
					return 0, ErrIntOverflowConsul
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
				return 0, ErrInvalidLengthConsul
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConsul
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConsul
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConsul        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsul          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConsul = fmt.Errorf("proto: unexpected end of group")
)
