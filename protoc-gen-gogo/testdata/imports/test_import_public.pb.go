// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_import_public.proto

package imports // import "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import test_a "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports/test_a_1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// M1 from public import imports/test_a_1/m1.proto
type M1 test_a.M1

func (m *M1) Reset()                         { (*test_a.M1)(m).Reset() }
func (m *M1) String() string                 { return (*test_a.M1)(m).String() }
func (*M1) ProtoMessage()                    {}
func (m *M1) XXX_Unmarshal(buf []byte) error { return (*test_a.M1)(m).XXX_Unmarshal(buf) }
func (m *M1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return (*test_a.M1)(m).XXX_Marshal(b, deterministic)
}
func (m *M1) XXX_Size() int       { return (*test_a.M1)(m).XXX_Size() }
func (m *M1) XXX_DiscardUnknown() { (*test_a.M1)(m).XXX_DiscardUnknown() }

type Public struct {
	F                    *test_a.M1 `protobuf:"bytes,1,opt,name=f" json:"f,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Public) Reset()         { *m = Public{} }
func (m *Public) String() string { return proto.CompactTextString(m) }
func (*Public) ProtoMessage()    {}
func (*Public) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_import_public_1d234712ec8ca1b8, []int{0}
}
func (m *Public) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Public.Unmarshal(m, b)
}
func (m *Public) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Public.Marshal(b, m, deterministic)
}
func (dst *Public) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Public.Merge(dst, src)
}
func (m *Public) XXX_Size() int {
	return xxx_messageInfo_Public.Size(m)
}
func (m *Public) XXX_DiscardUnknown() {
	xxx_messageInfo_Public.DiscardUnknown(m)
}

var xxx_messageInfo_Public proto.InternalMessageInfo

func (m *Public) GetF() *test_a.M1 {
	if m != nil {
		return m.F
	}
	return nil
}

func init() {
	proto.RegisterType((*Public)(nil), "test.Public")
}

func init() {
	proto.RegisterFile("imports/test_import_public.proto", fileDescriptor_test_import_public_1d234712ec8ca1b8)
}

var fileDescriptor_test_import_public_1d234712ec8ca1b8 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x87, 0x70, 0xe2, 0x0b, 0x4a, 0x93, 0x72,
	0x32, 0x93, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0x40, 0x32, 0x52, 0x92, 0x28, 0xea,
	0x12, 0xe3, 0x0d, 0xf5, 0x73, 0x0d, 0x21, 0x0a, 0x94, 0x94, 0xb8, 0xd8, 0x02, 0xc0, 0x1a, 0x84,
	0x24, 0xb8, 0x18, 0xd3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0x40, 0x0a, 0xf5,
	0x12, 0xf5, 0x7c, 0x0d, 0x83, 0x18, 0xd3, 0x9c, 0xac, 0xa3, 0x2c, 0xd3, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x9a, 0x93, 0x4a, 0xd3,
	0x20, 0x8c, 0x64, 0xdd, 0xf4, 0xd4, 0x3c, 0x5d, 0xb0, 0x04, 0x48, 0x63, 0x4a, 0x62, 0x49, 0xa2,
	0x3e, 0xd4, 0xca, 0x00, 0x86, 0x24, 0x36, 0xb0, 0x1a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x71, 0x00, 0x80, 0xfe, 0xae, 0x00, 0x00, 0x00,
}