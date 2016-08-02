// Code generated by protoc-gen-gogo.
// source: proto3.proto
// DO NOT EDIT!

/*
Package vanity is a generated protocol buffer package.

It is generated from these files:
	proto3.proto

It has these top-level messages:
	Aproto3
*/
package vanity

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Aproto3 struct {
	B string `protobuf:"bytes,1,opt,name=B,json=b,proto3" json:"B,omitempty"`
}

func (m *Aproto3) Reset()                    { *m = Aproto3{} }
func (*Aproto3) ProtoMessage()               {}
func (*Aproto3) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0} }

func init() {
	proto.RegisterType((*Aproto3)(nil), "vanity.Aproto3")
}
func (this *Aproto3) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Aproto3)
	if !ok {
		that2, ok := that.(Aproto3)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.B != that1.B {
		return false
	}
	return true
}
func (this *Aproto3) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&vanity.Aproto3{")
	s = append(s, "B: "+fmt.Sprintf("%#v", this.B)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProto3(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringProto3(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *Aproto3) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Aproto3) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.B) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintProto3(data, i, uint64(len(m.B)))
		i += copy(data[i:], m.B)
	}
	return i, nil
}

func encodeFixed64Proto3(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Proto3(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProto3(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Aproto3) Size() (n int) {
	var l int
	_ = l
	l = len(m.B)
	if l > 0 {
		n += 1 + l + sovProto3(uint64(l))
	}
	return n
}

func sovProto3(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProto3(x uint64) (n int) {
	return sovProto3(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Aproto3) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Aproto3{`,
		`B:` + fmt.Sprintf("%v", this.B) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProto3(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Aproto3) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto3
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Aproto3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Aproto3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto3
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProto3
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto3(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto3
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
func skipProto3(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto3
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowProto3
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProto3
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProto3
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProto3
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProto3(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProto3 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto3   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto3.proto", fileDescriptorProto3) }

var fileDescriptorProto3 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd6, 0x03, 0x53, 0x42, 0x6c, 0x65, 0x89, 0x79, 0x99, 0x25, 0x95, 0x4a, 0xe2, 0x5c,
	0xec, 0x8e, 0x10, 0x09, 0x21, 0x1e, 0x2e, 0x46, 0x27, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0xc6, 0x24, 0x27, 0x9d, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x00, 0xf1, 0x87, 0x87, 0x72, 0x8c, 0x0d,
	0x8f, 0xe4, 0x18, 0x57, 0x00, 0xf1, 0x09, 0x20, 0xbe, 0x00, 0xc4, 0x0f, 0x80, 0xf8, 0xc5, 0x23,
	0xa0, 0x1c, 0x90, 0x9e, 0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x62, 0x06, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x37, 0xd7, 0x45, 0xf0, 0x65, 0x00, 0x00, 0x00,
}
