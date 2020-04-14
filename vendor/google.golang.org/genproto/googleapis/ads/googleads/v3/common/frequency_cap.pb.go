// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/common/frequency_cap.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// A rule specifying the maximum number of times an ad (or some set of ads) can
// be shown to a user over a particular time period.
type FrequencyCapEntry struct {
	// The key of a particular frequency cap. There can be no more
	// than one frequency cap with the same key.
	Key *FrequencyCapKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Maximum number of events allowed during the time range by this cap.
	Cap                  *wrappers.Int32Value `protobuf:"bytes,2,opt,name=cap,proto3" json:"cap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapEntry) Reset()         { *m = FrequencyCapEntry{} }
func (m *FrequencyCapEntry) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapEntry) ProtoMessage()    {}
func (*FrequencyCapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5d792890a02906, []int{0}
}

func (m *FrequencyCapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapEntry.Unmarshal(m, b)
}
func (m *FrequencyCapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapEntry.Marshal(b, m, deterministic)
}
func (m *FrequencyCapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapEntry.Merge(m, src)
}
func (m *FrequencyCapEntry) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapEntry.Size(m)
}
func (m *FrequencyCapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapEntry proto.InternalMessageInfo

func (m *FrequencyCapEntry) GetKey() *FrequencyCapKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FrequencyCapEntry) GetCap() *wrappers.Int32Value {
	if m != nil {
		return m.Cap
	}
	return nil
}

// A group of fields used as keys for a frequency cap.
// There can be no more than one frequency cap with the same key.
type FrequencyCapKey struct {
	// The level on which the cap is to be applied (e.g. ad group ad, ad group).
	// The cap is applied to all the entities of this level.
	Level enums.FrequencyCapLevelEnum_FrequencyCapLevel `protobuf:"varint,1,opt,name=level,proto3,enum=google.ads.googleads.v3.enums.FrequencyCapLevelEnum_FrequencyCapLevel" json:"level,omitempty"`
	// The type of event that the cap applies to (e.g. impression).
	EventType enums.FrequencyCapEventTypeEnum_FrequencyCapEventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=google.ads.googleads.v3.enums.FrequencyCapEventTypeEnum_FrequencyCapEventType" json:"event_type,omitempty"`
	// Unit of time the cap is defined at (e.g. day, week).
	TimeUnit enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit `protobuf:"varint,2,opt,name=time_unit,json=timeUnit,proto3,enum=google.ads.googleads.v3.enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit" json:"time_unit,omitempty"`
	// Number of time units the cap lasts.
	TimeLength           *wrappers.Int32Value `protobuf:"bytes,4,opt,name=time_length,json=timeLength,proto3" json:"time_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapKey) Reset()         { *m = FrequencyCapKey{} }
func (m *FrequencyCapKey) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapKey) ProtoMessage()    {}
func (*FrequencyCapKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5d792890a02906, []int{1}
}

func (m *FrequencyCapKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapKey.Unmarshal(m, b)
}
func (m *FrequencyCapKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapKey.Marshal(b, m, deterministic)
}
func (m *FrequencyCapKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapKey.Merge(m, src)
}
func (m *FrequencyCapKey) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapKey.Size(m)
}
func (m *FrequencyCapKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapKey.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapKey proto.InternalMessageInfo

func (m *FrequencyCapKey) GetLevel() enums.FrequencyCapLevelEnum_FrequencyCapLevel {
	if m != nil {
		return m.Level
	}
	return enums.FrequencyCapLevelEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetEventType() enums.FrequencyCapEventTypeEnum_FrequencyCapEventType {
	if m != nil {
		return m.EventType
	}
	return enums.FrequencyCapEventTypeEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeUnit() enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit {
	if m != nil {
		return m.TimeUnit
	}
	return enums.FrequencyCapTimeUnitEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeLength() *wrappers.Int32Value {
	if m != nil {
		return m.TimeLength
	}
	return nil
}

func init() {
	proto.RegisterType((*FrequencyCapEntry)(nil), "google.ads.googleads.v3.common.FrequencyCapEntry")
	proto.RegisterType((*FrequencyCapKey)(nil), "google.ads.googleads.v3.common.FrequencyCapKey")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/common/frequency_cap.proto", fileDescriptor_9d5d792890a02906)
}

var fileDescriptor_9d5d792890a02906 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6a, 0xd4, 0x40,
	0x18, 0x27, 0x1b, 0x15, 0x3b, 0x85, 0x8a, 0x39, 0x2d, 0x55, 0x8a, 0xec, 0xc9, 0x8b, 0x33, 0x90,
	0x1c, 0x84, 0xb4, 0x97, 0xb4, 0x6e, 0x8b, 0xb8, 0x48, 0x09, 0x75, 0x0f, 0x12, 0x58, 0xa6, 0xc9,
	0xd7, 0x18, 0x4c, 0x66, 0xc6, 0x64, 0xb2, 0x92, 0x07, 0x10, 0xdf, 0xc3, 0xa3, 0x8f, 0xe2, 0xa3,
	0x88, 0x0f, 0x21, 0x33, 0x93, 0x89, 0xae, 0xcb, 0xba, 0xe4, 0x94, 0x5f, 0xe6, 0xfb, 0xfd, 0xf9,
	0xf2, 0x7d, 0x13, 0xe4, 0xe7, 0x9c, 0xe7, 0x25, 0x10, 0x9a, 0x35, 0xc4, 0x40, 0x85, 0xd6, 0x01,
	0x49, 0x79, 0x55, 0x71, 0x46, 0xee, 0x6a, 0xf8, 0xd4, 0x02, 0x4b, 0xbb, 0x55, 0x4a, 0x05, 0x16,
	0x35, 0x97, 0xdc, 0x3b, 0x31, 0x44, 0x4c, 0xb3, 0x06, 0x0f, 0x1a, 0xbc, 0x0e, 0xb0, 0xd1, 0x1c,
	0x9f, 0xed, 0xf2, 0x04, 0xd6, 0x56, 0xcd, 0xa6, 0xe5, 0x0a, 0xd6, 0xc0, 0xe4, 0x4a, 0x76, 0x02,
	0x8c, 0xfb, 0xf1, 0xcb, 0x31, 0xea, 0x12, 0xd6, 0x50, 0xf6, 0xc2, 0xd3, 0x31, 0x42, 0x59, 0x54,
	0xb0, 0x6a, 0x59, 0x21, 0x7b, 0x71, 0xff, 0x4d, 0x44, 0xbf, 0xdd, 0xb6, 0x77, 0xe4, 0x73, 0x4d,
	0x85, 0x80, 0xba, 0xe9, 0xeb, 0x4f, 0xad, 0xb9, 0x28, 0x08, 0x65, 0x8c, 0x4b, 0x2a, 0x0b, 0xce,
	0xfa, 0xea, 0xec, 0x8b, 0x83, 0x1e, 0x5f, 0x5a, 0xff, 0x0b, 0x2a, 0xe6, 0x4c, 0xd6, 0x9d, 0x17,
	0x21, 0xf7, 0x23, 0x74, 0x53, 0xe7, 0x99, 0xf3, 0xfc, 0xd0, 0x27, 0xf8, 0xff, 0x53, 0xc3, 0x7f,
	0xeb, 0xdf, 0x40, 0x17, 0x2b, 0xad, 0xf7, 0x02, 0xb9, 0x29, 0x15, 0xd3, 0x89, 0xb6, 0x78, 0x62,
	0x2d, 0x6c, 0x93, 0xf8, 0x35, 0x93, 0x81, 0xbf, 0xa4, 0x65, 0x0b, 0xb1, 0xe2, 0xcd, 0xbe, 0xba,
	0xe8, 0xd1, 0x3f, 0x3e, 0x5e, 0x82, 0xee, 0xeb, 0x29, 0xe9, 0x3e, 0x8e, 0xfc, 0xcb, 0x9d, 0x7d,
	0xe8, 0x31, 0x6d, 0xb4, 0xb1, 0x50, 0xba, 0x39, 0x6b, 0xab, 0xed, 0xd3, 0xd8, 0x98, 0x7a, 0x15,
	0x42, 0x7f, 0x36, 0x38, 0x75, 0x75, 0xc4, 0xdb, 0x11, 0x11, 0x73, 0x25, 0xbe, 0xe9, 0x04, 0x6c,
	0xc5, 0x0c, 0x95, 0xf8, 0x00, 0x2c, 0xf4, 0x0a, 0x74, 0x30, 0x6c, 0x4e, 0x4f, 0xe5, 0xc8, 0x5f,
	0x8c, 0x48, 0xbb, 0x29, 0x2a, 0x78, 0xc7, 0x0a, 0xb9, 0x15, 0x66, 0x0b, 0xf1, 0x43, 0xd9, 0x23,
	0xef, 0x0c, 0x1d, 0xea, 0xa8, 0x12, 0x58, 0x2e, 0x3f, 0x4c, 0xef, 0xed, 0x5f, 0x01, 0x52, 0xfc,
	0x85, 0xa6, 0x9f, 0xff, 0x72, 0xd0, 0x2c, 0xe5, 0xd5, 0x9e, 0xa5, 0x9f, 0x6f, 0xdc, 0x9a, 0x6b,
	0xe5, 0x79, 0xed, 0xbc, 0x7f, 0xd5, 0x8b, 0x72, 0x5e, 0x52, 0x96, 0x63, 0x5e, 0xe7, 0x24, 0x07,
	0xa6, 0x13, 0xed, 0xc5, 0x16, 0x45, 0xb3, 0xeb, 0x97, 0x3d, 0x35, 0x8f, 0x6f, 0x13, 0xf7, 0x2a,
	0x8a, 0xbe, 0x4f, 0x4e, 0xae, 0x8c, 0x59, 0x94, 0x35, 0xd8, 0x40, 0x85, 0x96, 0x01, 0xbe, 0xd0,
	0xb4, 0x1f, 0x96, 0x90, 0x44, 0x59, 0x93, 0x0c, 0x84, 0x64, 0x19, 0x24, 0x86, 0xf0, 0x73, 0x32,
	0x33, 0xa7, 0x61, 0x18, 0x65, 0x4d, 0x18, 0x0e, 0x94, 0x30, 0x5c, 0x06, 0x61, 0x68, 0x48, 0xb7,
	0x0f, 0x74, 0x77, 0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0x8d, 0xaf, 0xa2, 0x4f, 0x04,
	0x00, 0x00,
}
