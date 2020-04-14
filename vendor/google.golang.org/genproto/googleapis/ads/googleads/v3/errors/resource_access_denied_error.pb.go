// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/resource_access_denied_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError int32

const (
	// Enum unspecified.
	ResourceAccessDeniedErrorEnum_UNSPECIFIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 0
	// The received error code is not known in this version.
	ResourceAccessDeniedErrorEnum_UNKNOWN ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 1
	// User did not have write access.
	ResourceAccessDeniedErrorEnum_WRITE_ACCESS_DENIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 3
)

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "WRITE_ACCESS_DENIED",
}

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"WRITE_ACCESS_DENIED": 3,
}

func (x ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) String() string {
	return proto.EnumName(ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, int32(x))
}

func (ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2995c18210d516d0, []int{0, 0}
}

// Container for enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceAccessDeniedErrorEnum) Reset()         { *m = ResourceAccessDeniedErrorEnum{} }
func (m *ResourceAccessDeniedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ResourceAccessDeniedErrorEnum) ProtoMessage()    {}
func (*ResourceAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2995c18210d516d0, []int{0}
}

func (m *ResourceAccessDeniedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Unmarshal(m, b)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Marshal(b, m, deterministic)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.Merge(m, src)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Size(m)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAccessDeniedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError", ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value)
	proto.RegisterType((*ResourceAccessDeniedErrorEnum)(nil), "google.ads.googleads.v3.errors.ResourceAccessDeniedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/resource_access_denied_error.proto", fileDescriptor_2995c18210d516d0)
}

var fileDescriptor_2995c18210d516d0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x6a, 0x02, 0x31,
	0x18, 0x85, 0xab, 0x42, 0x0b, 0x71, 0x51, 0x99, 0x2e, 0x4a, 0x4b, 0xeb, 0x62, 0x0e, 0x90, 0x2c,
	0xb2, 0x4b, 0x57, 0xd1, 0x49, 0x45, 0x0a, 0x53, 0xd1, 0xaa, 0x50, 0x06, 0x86, 0x74, 0x12, 0xc2,
	0x80, 0x26, 0x92, 0x5f, 0x3d, 0x50, 0x97, 0x3d, 0x4a, 0x8f, 0xd2, 0x1b, 0x74, 0x57, 0x26, 0x51,
	0x77, 0x76, 0x95, 0x47, 0xf2, 0xfd, 0xef, 0xbd, 0xfc, 0x88, 0x1b, 0xe7, 0xcc, 0x4a, 0x13, 0xa9,
	0x80, 0x44, 0xd9, 0xa8, 0x3d, 0x25, 0xda, 0x7b, 0xe7, 0x81, 0x78, 0x0d, 0x6e, 0xe7, 0x2b, 0x5d,
	0xca, 0xaa, 0xd2, 0x00, 0xa5, 0xd2, 0xb6, 0xd6, 0xaa, 0x0c, 0xaf, 0x78, 0xe3, 0xdd, 0xd6, 0x25,
	0xfd, 0x38, 0x87, 0xa5, 0x02, 0x7c, 0xb2, 0xc0, 0x7b, 0x8a, 0xa3, 0xc5, 0xfd, 0xc3, 0x31, 0x62,
	0x53, 0x13, 0x69, 0xad, 0xdb, 0xca, 0x6d, 0xed, 0x2c, 0xc4, 0xe9, 0x14, 0xd0, 0xe3, 0xf4, 0x90,
	0xc1, 0x43, 0x44, 0x16, 0x12, 0x44, 0x33, 0x2b, 0xec, 0x6e, 0x9d, 0x4e, 0xd1, 0xdd, 0x59, 0x20,
	0xb9, 0x46, 0xdd, 0x79, 0x3e, 0x9b, 0x88, 0xe1, 0xf8, 0x79, 0x2c, 0xb2, 0xde, 0x45, 0xd2, 0x45,
	0x57, 0xf3, 0xfc, 0x25, 0x7f, 0x5d, 0xe6, 0xbd, 0x56, 0x72, 0x8b, 0x6e, 0x96, 0xd3, 0xf1, 0x9b,
	0x28, 0xf9, 0x70, 0x28, 0x66, 0xb3, 0x32, 0x13, 0x79, 0x43, 0x75, 0x06, 0xbf, 0x2d, 0x94, 0x56,
	0x6e, 0x8d, 0xff, 0x6f, 0x3e, 0xe8, 0x9f, 0x0d, 0x9e, 0x34, 0xdd, 0x27, 0xad, 0xf7, 0xec, 0xe0,
	0x60, 0xdc, 0x4a, 0x5a, 0x83, 0x9d, 0x37, 0xc4, 0x68, 0x1b, 0x7e, 0x76, 0x5c, 0xe7, 0xa6, 0x86,
	0x73, 0xdb, 0x7d, 0x8a, 0xc7, 0x67, 0xbb, 0x33, 0xe2, 0xfc, 0xab, 0xdd, 0x1f, 0x45, 0x33, 0xae,
	0x00, 0x47, 0xd9, 0xa8, 0x05, 0xc5, 0x21, 0x12, 0xbe, 0x8f, 0x40, 0xc1, 0x15, 0x14, 0x27, 0xa0,
	0x58, 0xd0, 0x22, 0x02, 0x3f, 0xed, 0x34, 0xde, 0x32, 0xc6, 0x15, 0x30, 0x76, 0x42, 0x18, 0x5b,
	0x50, 0xc6, 0x22, 0xf4, 0x71, 0x19, 0xda, 0xd1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x24,
	0x4d, 0x25, 0xfa, 0x01, 0x00, 0x00,
}
