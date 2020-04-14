// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/actions/type/datetime_range.proto

package date_time_range

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

// Represents a date and time range. This can represent:
//
// * A range between points in time with time zone or offset, e.g. the duration
//   of a flight which starts in the "America/New_York" time zone and ends in
//   the "Australia/Sydney" time zone
// * A range between points in time without time zone/offset info, e.g. an
//   appointment in local time
// * A range starting at a specific date and time, e.g. the range of time which
//   can be measured in milliseconds since the Unix epoch (period starting with
//   1970-01-01T00:00:00Z)
// * A range ending at a specific date and time, e.g. range of time before
//   a deadline
//
// When considering whether a DateTime falls within a DateTimeRange, the start
// of the range is inclusive and the end is exclusive.
//
// While [google.type.DateTime][google.type.DateTime] allows zero years, DateTimeRange does not.
// Year must always be non-zero.
//
// When both start and end are set, either both or neither must have a
// time_offset. When set, time_offset can be specified by either utc_offset or
// time_zone and must match for start and end, that is if start has utc_offset
// set then end must also have utc_offset set. The values of utc_offset or
// time_zone need not be the same for start and end.
//
// When both start and end are set, start must be chronologically less than or
// equal to end. When start and end are equal, the range is empty.
//
// The semantics of start and end are the same as those of
// [google.type.DateTime][google.type.DateTime].
type DateTimeRange struct {
	// DateTime at which the date range begins. If unset, the range has no
	// beginning bound.
	Start *datetime.DateTime `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// DateTime at which the date range ends. If unset, the range has no ending
	// bound.
	End                  *datetime.DateTime `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DateTimeRange) Reset()         { *m = DateTimeRange{} }
func (m *DateTimeRange) String() string { return proto.CompactTextString(m) }
func (*DateTimeRange) ProtoMessage()    {}
func (*DateTimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b47322960d743f, []int{0}
}

func (m *DateTimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateTimeRange.Unmarshal(m, b)
}
func (m *DateTimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateTimeRange.Marshal(b, m, deterministic)
}
func (m *DateTimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateTimeRange.Merge(m, src)
}
func (m *DateTimeRange) XXX_Size() int {
	return xxx_messageInfo_DateTimeRange.Size(m)
}
func (m *DateTimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DateTimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_DateTimeRange proto.InternalMessageInfo

func (m *DateTimeRange) GetStart() *datetime.DateTime {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *DateTimeRange) GetEnd() *datetime.DateTime {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*DateTimeRange)(nil), "google.actions.type.DateTimeRange")
}

func init() {
	proto.RegisterFile("google/actions/type/datetime_range.proto", fileDescriptor_55b47322960d743f)
}

var fileDescriptor_55b47322960d743f = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0xd5,
	0x4f, 0x49, 0x2c, 0x49, 0x2d, 0xc9, 0xcc, 0x4d, 0x8d, 0x2f, 0x4a, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0xa8, 0xd4, 0x83, 0xaa, 0xd4, 0x03, 0xa9, 0x94, 0x92,
	0x82, 0x6a, 0x47, 0xd1, 0x06, 0xd1, 0xa0, 0x94, 0xca, 0xc5, 0xeb, 0x92, 0x58, 0x92, 0x1a, 0x92,
	0x99, 0x9b, 0x1a, 0x04, 0x32, 0x47, 0x48, 0x9b, 0x8b, 0xb5, 0xb8, 0x24, 0xb1, 0xa8, 0x44, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x54, 0x0f, 0x6a, 0x22, 0x48, 0xb3, 0x1e, 0x5c, 0x29, 0x44,
	0x8d, 0x90, 0x3a, 0x17, 0x73, 0x6a, 0x5e, 0x8a, 0x04, 0x13, 0x3e, 0xa5, 0x20, 0x15, 0x4e, 0xcd,
	0x8c, 0x5c, 0xe2, 0xc9, 0xf9, 0xb9, 0x7a, 0x58, 0x9c, 0xe7, 0x24, 0x84, 0xe2, 0x80, 0x00, 0x90,
	0xb3, 0x02, 0x18, 0xa3, 0xbc, 0xa0, 0x4a, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b,
	0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0x8e, 0xd6, 0x87, 0x48, 0x25, 0x16, 0x64, 0x22, 0x05, 0x45,
	0x3c, 0x22, 0x2c, 0xac, 0xd1, 0xf8, 0x8b, 0x98, 0x58, 0x1d, 0xfd, 0xdd, 0x43, 0x02, 0x92, 0xd8,
	0xc0, 0xda, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x51, 0xa0, 0x92, 0x50, 0x01, 0x00,
	0x00,
}
