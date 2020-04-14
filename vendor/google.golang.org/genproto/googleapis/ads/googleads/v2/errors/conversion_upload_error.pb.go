// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/conversion_upload_error.proto

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

// Enum describing possible conversion upload errors.
type ConversionUploadErrorEnum_ConversionUploadError int32

const (
	// Enum unspecified.
	ConversionUploadErrorEnum_UNSPECIFIED ConversionUploadErrorEnum_ConversionUploadError = 0
	// The received error code is not known in this version.
	ConversionUploadErrorEnum_UNKNOWN ConversionUploadErrorEnum_ConversionUploadError = 1
	// The request contained more than 2000 conversions.
	ConversionUploadErrorEnum_TOO_MANY_CONVERSIONS_IN_REQUEST ConversionUploadErrorEnum_ConversionUploadError = 2
	// The specified gclid could not be decoded.
	ConversionUploadErrorEnum_UNPARSEABLE_GCLID ConversionUploadErrorEnum_ConversionUploadError = 3
	// The specified conversion_date_time is before the event time
	// associated with the given gclid.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_GCLID ConversionUploadErrorEnum_ConversionUploadError = 4
	// The click associated with the given gclid is either too old to be
	// imported or occurred outside of the click through lookback window for the
	// specified conversion action.
	ConversionUploadErrorEnum_EXPIRED_GCLID ConversionUploadErrorEnum_ConversionUploadError = 5
	// The click associated with the given gclid occurred too recently. Please
	// try uploading again after 6 hours have passed since the click occurred.
	ConversionUploadErrorEnum_TOO_RECENT_GCLID ConversionUploadErrorEnum_ConversionUploadError = 6
	// The click associated with the given gclid could not be found in the
	// system. This can happen if Google Click IDs are collected for non Google
	// Ads clicks.
	ConversionUploadErrorEnum_GCLID_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 7
	// The click associated with the given gclid is owned by a customer
	// account that the uploading customer does not manage.
	ConversionUploadErrorEnum_UNAUTHORIZED_CUSTOMER ConversionUploadErrorEnum_ConversionUploadError = 8
	// No upload eligible conversion action that matches the provided
	// information can be found for the customer.
	ConversionUploadErrorEnum_INVALID_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 9
	// The specified conversion action was created too recently.
	// Please try the upload again after 4-6 hours have passed since the
	// conversion action was created.
	ConversionUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 10
	// The click associated with the given gclid does not contain conversion
	// tracking information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME ConversionUploadErrorEnum_ConversionUploadError = 11
	// The specified conversion action does not use an external attribution
	// model, but external_attribution_data was set.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 12
	// The specified conversion action uses an external attribution model, but
	// external_attribution_data or one of its contained fields was not set.
	// Both external_attribution_credit and external_attribution_model must be
	// set for externally attributed conversion actions.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 13
	// Order IDs are not supported for conversion actions which use an external
	// attribution model.
	ConversionUploadErrorEnum_ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 14
	// A conversion with the same order id and conversion action combination
	// already exists in our system.
	ConversionUploadErrorEnum_ORDER_ID_ALREADY_IN_USE ConversionUploadErrorEnum_ConversionUploadError = 15
	// The request contained two or more conversions with the same order id and
	// conversion action combination.
	ConversionUploadErrorEnum_DUPLICATE_ORDER_ID ConversionUploadErrorEnum_ConversionUploadError = 16
	// The call occurred too recently. Please try uploading again after 6 hours
	// have passed since the call occurred.
	ConversionUploadErrorEnum_TOO_RECENT_CALL ConversionUploadErrorEnum_ConversionUploadError = 17
	// The click that initiated the call is too old for this conversion to be
	// imported.
	ConversionUploadErrorEnum_EXPIRED_CALL ConversionUploadErrorEnum_ConversionUploadError = 18
	// The call or the click leading to the call was not found.
	ConversionUploadErrorEnum_CALL_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 19
	// The specified conversion_date_time is before the call_start_date_time.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_CALL ConversionUploadErrorEnum_ConversionUploadError = 20
	// The click associated with the call does not contain conversion tracking
	// information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME ConversionUploadErrorEnum_ConversionUploadError = 21
	// The caller’s phone number cannot be parsed. It should be formatted either
	// as E.164 "+16502531234", International "+64 3-331 6005" or US national
	// number "6502531234".
	ConversionUploadErrorEnum_UNPARSEABLE_CALLERS_PHONE_NUMBER ConversionUploadErrorEnum_ConversionUploadError = 22
)

var ConversionUploadErrorEnum_ConversionUploadError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "TOO_MANY_CONVERSIONS_IN_REQUEST",
	3:  "UNPARSEABLE_GCLID",
	4:  "CONVERSION_PRECEDES_GCLID",
	5:  "EXPIRED_GCLID",
	6:  "TOO_RECENT_GCLID",
	7:  "GCLID_NOT_FOUND",
	8:  "UNAUTHORIZED_CUSTOMER",
	9:  "INVALID_CONVERSION_ACTION",
	10: "TOO_RECENT_CONVERSION_ACTION",
	11: "CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME",
	12: "EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	13: "EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	14: "ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	15: "ORDER_ID_ALREADY_IN_USE",
	16: "DUPLICATE_ORDER_ID",
	17: "TOO_RECENT_CALL",
	18: "EXPIRED_CALL",
	19: "CALL_NOT_FOUND",
	20: "CONVERSION_PRECEDES_CALL",
	21: "CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME",
	22: "UNPARSEABLE_CALLERS_PHONE_NUMBER",
}

var ConversionUploadErrorEnum_ConversionUploadError_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"TOO_MANY_CONVERSIONS_IN_REQUEST": 2,
	"UNPARSEABLE_GCLID":               3,
	"CONVERSION_PRECEDES_GCLID":       4,
	"EXPIRED_GCLID":                   5,
	"TOO_RECENT_GCLID":                6,
	"GCLID_NOT_FOUND":                 7,
	"UNAUTHORIZED_CUSTOMER":           8,
	"INVALID_CONVERSION_ACTION":       9,
	"TOO_RECENT_CONVERSION_ACTION":    10,
	"CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME":                            11,
	"EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 12,
	"EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 13,
	"ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION":            14,
	"ORDER_ID_ALREADY_IN_USE":                      15,
	"DUPLICATE_ORDER_ID":                           16,
	"TOO_RECENT_CALL":                              17,
	"EXPIRED_CALL":                                 18,
	"CALL_NOT_FOUND":                               19,
	"CONVERSION_PRECEDES_CALL":                     20,
	"CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME": 21,
	"UNPARSEABLE_CALLERS_PHONE_NUMBER":             22,
}

func (x ConversionUploadErrorEnum_ConversionUploadError) String() string {
	return proto.EnumName(ConversionUploadErrorEnum_ConversionUploadError_name, int32(x))
}

func (ConversionUploadErrorEnum_ConversionUploadError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0f2c8da0320cd9a9, []int{0, 0}
}

// Container for enum describing possible conversion upload errors.
type ConversionUploadErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionUploadErrorEnum) Reset()         { *m = ConversionUploadErrorEnum{} }
func (m *ConversionUploadErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionUploadErrorEnum) ProtoMessage()    {}
func (*ConversionUploadErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f2c8da0320cd9a9, []int{0}
}

func (m *ConversionUploadErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionUploadErrorEnum.Unmarshal(m, b)
}
func (m *ConversionUploadErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionUploadErrorEnum.Marshal(b, m, deterministic)
}
func (m *ConversionUploadErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionUploadErrorEnum.Merge(m, src)
}
func (m *ConversionUploadErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionUploadErrorEnum.Size(m)
}
func (m *ConversionUploadErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionUploadErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionUploadErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.ConversionUploadErrorEnum_ConversionUploadError", ConversionUploadErrorEnum_ConversionUploadError_name, ConversionUploadErrorEnum_ConversionUploadError_value)
	proto.RegisterType((*ConversionUploadErrorEnum)(nil), "google.ads.googleads.v2.errors.ConversionUploadErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/conversion_upload_error.proto", fileDescriptor_0f2c8da0320cd9a9)
}

var fileDescriptor_0f2c8da0320cd9a9 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xa6, 0x29, 0xa4, 0xb0, 0xfd, 0xdb, 0x6e, 0x9b, 0x42, 0x4b, 0x29, 0x55, 0xe0, 0x88, 0x1c,
	0x14, 0x24, 0x0e, 0x86, 0xcb, 0xc6, 0x9e, 0xa6, 0xab, 0x3a, 0x6b, 0x77, 0xbd, 0x1b, 0xda, 0x2a,
	0xd2, 0x2a, 0x34, 0x51, 0x14, 0xa9, 0xf5, 0x46, 0x71, 0xdb, 0x07, 0xe2, 0xc8, 0x03, 0xf0, 0x10,
	0x9c, 0x79, 0x0a, 0xee, 0xdc, 0xd1, 0xda, 0x71, 0x88, 0x44, 0xa8, 0xca, 0x29, 0xa3, 0x99, 0xef,
	0x67, 0xe2, 0x99, 0x59, 0xf4, 0x71, 0x60, 0xcc, 0xe0, 0xb2, 0x5f, 0xeb, 0xf6, 0xd2, 0x5a, 0x1e,
	0xda, 0xe8, 0xb6, 0x5e, 0xeb, 0x8f, 0xc7, 0x66, 0x9c, 0xd6, 0x2e, 0x4c, 0x72, 0xdb, 0x1f, 0xa7,
	0x43, 0x93, 0xe8, 0x9b, 0xd1, 0xa5, 0xe9, 0xf6, 0x74, 0x56, 0x70, 0x46, 0x63, 0x73, 0x6d, 0xc8,
	0x7e, 0x4e, 0x71, 0xba, 0xbd, 0xd4, 0x99, 0xb2, 0x9d, 0xdb, 0xba, 0x93, 0xb3, 0x77, 0xf7, 0x0a,
	0xf5, 0xd1, 0xb0, 0xd6, 0x4d, 0x12, 0x73, 0xdd, 0xbd, 0x1e, 0x9a, 0x24, 0xcd, 0xd9, 0xd5, 0x1f,
	0x65, 0xb4, 0xe3, 0x4d, 0xf5, 0x55, 0x26, 0x0f, 0x96, 0x08, 0xc9, 0xcd, 0x55, 0xf5, 0x5b, 0x19,
	0x55, 0xe6, 0x56, 0xc9, 0x3a, 0x5a, 0x56, 0x3c, 0x8e, 0xc0, 0x63, 0x87, 0x0c, 0x7c, 0xfc, 0x80,
	0x2c, 0xa3, 0x25, 0xc5, 0x8f, 0x79, 0xf8, 0x89, 0xe3, 0x05, 0xf2, 0x0a, 0xbd, 0x94, 0x61, 0xa8,
	0x5b, 0x94, 0x9f, 0x69, 0x2f, 0xe4, 0x6d, 0x10, 0x31, 0x0b, 0x79, 0xac, 0x19, 0xd7, 0x02, 0x4e,
	0x14, 0xc4, 0x12, 0x97, 0x48, 0x05, 0x6d, 0x28, 0x1e, 0x51, 0x11, 0x03, 0x6d, 0x04, 0xa0, 0x9b,
	0x5e, 0xc0, 0x7c, 0xbc, 0x48, 0x5e, 0xa0, 0x9d, 0x3f, 0x14, 0x1d, 0x09, 0xf0, 0xc0, 0x87, 0x78,
	0x52, 0x7e, 0x48, 0x36, 0xd0, 0x2a, 0x9c, 0x46, 0x4c, 0x80, 0x3f, 0x49, 0x3d, 0x22, 0x5b, 0x08,
	0x5b, 0x37, 0x8b, 0xe4, 0x72, 0x92, 0x2d, 0x93, 0x4d, 0xb4, 0x9e, 0x85, 0x9a, 0x87, 0x52, 0x1f,
	0x86, 0x8a, 0xfb, 0x78, 0x89, 0xec, 0xa0, 0x8a, 0xe2, 0x54, 0xc9, 0xa3, 0x50, 0xb0, 0x73, 0xf0,
	0xb5, 0xa7, 0x62, 0x19, 0xb6, 0x40, 0xe0, 0xc7, 0xd6, 0x97, 0xf1, 0x36, 0xb5, 0x8c, 0x19, 0x7f,
	0xea, 0x49, 0x16, 0x72, 0xfc, 0x84, 0x1c, 0xa0, 0xbd, 0x19, 0x93, 0xbf, 0x11, 0x88, 0xbc, 0x47,
	0xf5, 0x99, 0xb4, 0x14, 0xd4, 0x3b, 0x66, 0xbc, 0x99, 0xd9, 0x03, 0xb7, 0x7f, 0xd1, 0xd7, 0x54,
	0x6a, 0xd6, 0x8a, 0x04, 0xc4, 0x39, 0x84, 0xb5, 0x00, 0x2f, 0x93, 0x13, 0xd4, 0x82, 0x53, 0x09,
	0x82, 0xd3, 0x40, 0x53, 0x29, 0x05, 0x6b, 0x28, 0xab, 0xa8, 0x7d, 0x2a, 0xa9, 0x8e, 0xc1, 0x36,
	0x2f, 0x34, 0x0f, 0xb9, 0x2e, 0x50, 0xc1, 0xd9, 0x14, 0x07, 0xf3, 0x9a, 0x5d, 0xb9, 0x5b, 0xd2,
	0x36, 0x54, 0xc8, 0xde, 0x57, 0x72, 0x95, 0x1c, 0xa2, 0x46, 0x28, 0x7c, 0x10, 0x7a, 0xf2, 0x45,
	0x23, 0x10, 0x2d, 0x26, 0x2d, 0xfa, 0x7f, 0x74, 0xd6, 0xc8, 0x73, 0xf4, 0x74, 0xaa, 0x43, 0x03,
	0x01, 0xd4, 0x3f, 0xb3, 0x6b, 0xa1, 0x62, 0xc0, 0xeb, 0x64, 0x1b, 0x11, 0x5f, 0x45, 0x01, 0xf3,
	0xa8, 0x04, 0x5d, 0xc0, 0x30, 0xb6, 0xb3, 0x9c, 0xfd, 0xf8, 0x34, 0x08, 0xf0, 0x06, 0xc1, 0x68,
	0xa5, 0xd8, 0x84, 0x2c, 0x43, 0x08, 0x41, 0x6b, 0x36, 0x9a, 0x99, 0xf8, 0x26, 0xd9, 0x43, 0xcf,
	0xe6, 0xad, 0x53, 0xc6, 0xd8, 0x22, 0x6f, 0xd1, 0x9b, 0x7b, 0xcc, 0x2c, 0x13, 0xcd, 0xa6, 0x55,
	0x21, 0xaf, 0xd1, 0xc1, 0xec, 0xd6, 0xda, 0x12, 0x88, 0x58, 0x47, 0x47, 0x21, 0x07, 0xcd, 0x55,
	0xab, 0x01, 0x02, 0x6f, 0x37, 0x7e, 0x2d, 0xa0, 0xea, 0x85, 0xb9, 0x72, 0xee, 0xbe, 0xcd, 0xc6,
	0xee, 0xdc, 0xe3, 0x8a, 0xec, 0x65, 0x46, 0x0b, 0xe7, 0xfe, 0x84, 0x3d, 0x30, 0x97, 0xdd, 0x64,
	0xe0, 0x98, 0xf1, 0xa0, 0x36, 0xe8, 0x27, 0xd9, 0xdd, 0x16, 0xef, 0xc4, 0x68, 0x98, 0xfe, 0xeb,
	0xd9, 0xf8, 0x90, 0xff, 0x7c, 0x29, 0x2d, 0x36, 0x29, 0xfd, 0x5a, 0xda, 0x6f, 0xe6, 0x62, 0xb4,
	0x97, 0x3a, 0x79, 0x68, 0xa3, 0x76, 0xdd, 0xc9, 0x2c, 0xd3, 0xef, 0x05, 0xa0, 0x43, 0x7b, 0x69,
	0x67, 0x0a, 0xe8, 0xb4, 0xeb, 0x9d, 0x1c, 0xf0, 0xb3, 0x54, 0xcd, 0xb3, 0xae, 0x4b, 0x7b, 0xa9,
	0xeb, 0x4e, 0x21, 0xae, 0xdb, 0xae, 0xbb, 0x6e, 0x0e, 0xfa, 0x5c, 0xce, 0xba, 0x7b, 0xf7, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x03, 0xc0, 0x5d, 0x29, 0xd3, 0x04, 0x00, 0x00,
}
