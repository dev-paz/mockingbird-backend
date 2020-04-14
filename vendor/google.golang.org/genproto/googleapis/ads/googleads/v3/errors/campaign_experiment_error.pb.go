// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/campaign_experiment_error.proto

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

// Enum describing possible campaign experiment errors.
type CampaignExperimentErrorEnum_CampaignExperimentError int32

const (
	// Enum unspecified.
	CampaignExperimentErrorEnum_UNSPECIFIED CampaignExperimentErrorEnum_CampaignExperimentError = 0
	// The received error code is not known in this version.
	CampaignExperimentErrorEnum_UNKNOWN CampaignExperimentErrorEnum_CampaignExperimentError = 1
	// An active campaign or experiment with this name already exists.
	CampaignExperimentErrorEnum_DUPLICATE_NAME CampaignExperimentErrorEnum_CampaignExperimentError = 2
	// Experiment cannot be updated from the current state to the
	// requested target state. For example, an experiment can only graduate
	// if its status is ENABLED.
	CampaignExperimentErrorEnum_INVALID_TRANSITION CampaignExperimentErrorEnum_CampaignExperimentError = 3
	// Cannot create an experiment from a campaign using an explicitly shared
	// budget.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET CampaignExperimentErrorEnum_CampaignExperimentError = 4
	// Cannot create an experiment for a removed base campaign.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN CampaignExperimentErrorEnum_CampaignExperimentError = 5
	// Cannot create an experiment from a draft, which has a status other than
	// proposed.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT CampaignExperimentErrorEnum_CampaignExperimentError = 6
	// This customer is not allowed to create an experiment.
	CampaignExperimentErrorEnum_CUSTOMER_CANNOT_CREATE_EXPERIMENT CampaignExperimentErrorEnum_CampaignExperimentError = 7
	// This campaign is not allowed to create an experiment.
	CampaignExperimentErrorEnum_CAMPAIGN_CANNOT_CREATE_EXPERIMENT CampaignExperimentErrorEnum_CampaignExperimentError = 8
	// Trying to set an experiment duration which overlaps with another
	// experiment.
	CampaignExperimentErrorEnum_EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP CampaignExperimentErrorEnum_CampaignExperimentError = 9
	// All non-removed experiments must start and end within their campaign's
	// duration.
	CampaignExperimentErrorEnum_EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION CampaignExperimentErrorEnum_CampaignExperimentError = 10
	// The experiment cannot be modified because its status is in a terminal
	// state, such as REMOVED.
	CampaignExperimentErrorEnum_CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS CampaignExperimentErrorEnum_CampaignExperimentError = 11
)

var CampaignExperimentErrorEnum_CampaignExperimentError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DUPLICATE_NAME",
	3:  "INVALID_TRANSITION",
	4:  "CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET",
	5:  "CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN",
	6:  "CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT",
	7:  "CUSTOMER_CANNOT_CREATE_EXPERIMENT",
	8:  "CAMPAIGN_CANNOT_CREATE_EXPERIMENT",
	9:  "EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP",
	10: "EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION",
	11: "CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS",
}

var CampaignExperimentErrorEnum_CampaignExperimentError_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"DUPLICATE_NAME":     2,
	"INVALID_TRANSITION": 3,
	"CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET":          4,
	"CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN":   5,
	"CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT":      6,
	"CUSTOMER_CANNOT_CREATE_EXPERIMENT":                    7,
	"CAMPAIGN_CANNOT_CREATE_EXPERIMENT":                    8,
	"EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP":                9,
	"EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION": 10,
	"CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS":               11,
}

func (x CampaignExperimentErrorEnum_CampaignExperimentError) String() string {
	return proto.EnumName(CampaignExperimentErrorEnum_CampaignExperimentError_name, int32(x))
}

func (CampaignExperimentErrorEnum_CampaignExperimentError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9f22a1dd7457e60, []int{0, 0}
}

// Container for enum describing possible campaign experiment errors.
type CampaignExperimentErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignExperimentErrorEnum) Reset()         { *m = CampaignExperimentErrorEnum{} }
func (m *CampaignExperimentErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignExperimentErrorEnum) ProtoMessage()    {}
func (*CampaignExperimentErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9f22a1dd7457e60, []int{0}
}

func (m *CampaignExperimentErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignExperimentErrorEnum.Unmarshal(m, b)
}
func (m *CampaignExperimentErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignExperimentErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignExperimentErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignExperimentErrorEnum.Merge(m, src)
}
func (m *CampaignExperimentErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignExperimentErrorEnum.Size(m)
}
func (m *CampaignExperimentErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignExperimentErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignExperimentErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.CampaignExperimentErrorEnum_CampaignExperimentError", CampaignExperimentErrorEnum_CampaignExperimentError_name, CampaignExperimentErrorEnum_CampaignExperimentError_value)
	proto.RegisterType((*CampaignExperimentErrorEnum)(nil), "google.ads.googleads.v3.errors.CampaignExperimentErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/campaign_experiment_error.proto", fileDescriptor_f9f22a1dd7457e60)
}

var fileDescriptor_f9f22a1dd7457e60 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x8b, 0xd3, 0x30,
	0x1c, 0xc7, 0xbd, 0xed, 0xbc, 0xd3, 0x0c, 0x34, 0xe4, 0x41, 0x41, 0x8f, 0x03, 0x07, 0x27, 0xa8,
	0xd0, 0x82, 0x15, 0x91, 0x0a, 0x42, 0xd6, 0xfe, 0xb6, 0x0b, 0xae, 0x49, 0x49, 0xd3, 0x9e, 0xc8,
	0x20, 0xd4, 0x5b, 0x29, 0x83, 0x5b, 0x3b, 0xda, 0x79, 0xf8, 0xe8, 0xdf, 0xe2, 0xa3, 0x7f, 0x8a,
	0x7f, 0x85, 0xcf, 0xfe, 0x03, 0xbe, 0x4a, 0x9b, 0xad, 0x88, 0xd8, 0x7b, 0xea, 0x8f, 0xe6, 0xf3,
	0xc9, 0x37, 0x09, 0x5f, 0xf4, 0x2e, 0x2f, 0xcb, 0xfc, 0x2a, 0xb3, 0xd3, 0x65, 0x6d, 0x9b, 0xb1,
	0x99, 0xae, 0x1d, 0x3b, 0xab, 0xaa, 0xb2, 0xaa, 0xed, 0xcb, 0x74, 0xbd, 0x49, 0x57, 0x79, 0xa1,
	0xb3, 0x2f, 0x9b, 0xac, 0x5a, 0xad, 0xb3, 0x62, 0xab, 0xdb, 0x25, 0x6b, 0x53, 0x95, 0xdb, 0x92,
	0x9c, 0x1a, 0xc9, 0x4a, 0x97, 0xb5, 0xd5, 0xf9, 0xd6, 0xb5, 0x63, 0x19, 0xff, 0xd1, 0xc9, 0x7e,
	0xff, 0xcd, 0xca, 0x4e, 0x8b, 0xa2, 0xdc, 0xa6, 0xdb, 0x55, 0x59, 0xd4, 0xc6, 0x1e, 0x7f, 0x3d,
	0x44, 0x8f, 0xbd, 0x5d, 0x02, 0x74, 0x01, 0xd0, 0xa8, 0x50, 0x7c, 0x5e, 0x8f, 0x7f, 0x0e, 0xd1,
	0xc3, 0x9e, 0x75, 0x72, 0x1f, 0x8d, 0x62, 0x1e, 0x85, 0xe0, 0xb1, 0x29, 0x03, 0x1f, 0xdf, 0x22,
	0x23, 0x74, 0x1c, 0xf3, 0xf7, 0x5c, 0x5c, 0x70, 0x7c, 0x40, 0x08, 0xba, 0xe7, 0xc7, 0xe1, 0x9c,
	0x79, 0x54, 0x81, 0xe6, 0x34, 0x00, 0x3c, 0x20, 0x0f, 0x10, 0x61, 0x3c, 0xa1, 0x73, 0xe6, 0x6b,
	0x25, 0x29, 0x8f, 0x98, 0x62, 0x82, 0xe3, 0x21, 0xb1, 0xd1, 0x0b, 0x8f, 0x72, 0x2e, 0x94, 0xf6,
	0x24, 0x34, 0x3c, 0x7c, 0x08, 0x41, 0xb2, 0x00, 0xb8, 0xd2, 0x17, 0x4c, 0x9d, 0xeb, 0xe8, 0x9c,
	0x4a, 0xf0, 0xf5, 0x24, 0xf6, 0x67, 0xa0, 0xf0, 0x21, 0x79, 0x8d, 0x5e, 0xf6, 0x0a, 0x53, 0x21,
	0xb5, 0x84, 0x40, 0x24, 0x8d, 0x40, 0x23, 0xd0, 0x1e, 0x0d, 0x42, 0xca, 0x66, 0x1c, 0xdf, 0x26,
	0x0e, 0xb2, 0x6f, 0xf4, 0xb8, 0xe0, 0x3a, 0x94, 0x22, 0x14, 0x11, 0xf8, 0xda, 0x97, 0x74, 0xaa,
	0xf0, 0x11, 0x39, 0x43, 0x4f, 0xbc, 0x38, 0x52, 0x22, 0x00, 0xa9, 0xfb, 0x6c, 0x7c, 0xdc, 0x62,
	0xbb, 0xa4, 0x7e, 0xec, 0x0e, 0x79, 0x86, 0xce, 0xfe, 0x0a, 0xf5, 0x63, 0x49, 0x9b, 0x47, 0x88,
	0x74, 0x10, 0x47, 0x4a, 0x37, 0x92, 0x48, 0x40, 0xce, 0x69, 0x88, 0xef, 0x92, 0x37, 0xe8, 0xd5,
	0x7f, 0x50, 0x43, 0x4e, 0xa0, 0x7d, 0x19, 0xc6, 0xbb, 0xeb, 0x75, 0x00, 0x46, 0xe4, 0x39, 0x7a,
	0xba, 0x3b, 0x42, 0x10, 0xab, 0x7f, 0xee, 0xe9, 0xc7, 0xa0, 0x95, 0xd0, 0x91, 0xa2, 0x2a, 0x8e,
	0xf0, 0x68, 0xf2, 0xfb, 0x00, 0x8d, 0x2f, 0xcb, 0xb5, 0x75, 0x73, 0x8f, 0x26, 0x27, 0x3d, 0x35,
	0x08, 0x9b, 0x1e, 0x85, 0x07, 0x1f, 0xfd, 0x9d, 0x9f, 0x97, 0x57, 0x69, 0x91, 0x5b, 0x65, 0x95,
	0xdb, 0x79, 0x56, 0xb4, 0x2d, 0xdb, 0xf7, 0x7a, 0xb3, 0xaa, 0xfb, 0x6a, 0xfe, 0xd6, 0x7c, 0xbe,
	0x0d, 0x86, 0x33, 0x4a, 0xbf, 0x0f, 0x4e, 0x67, 0x66, 0x33, 0xba, 0xac, 0x2d, 0x33, 0x36, 0x53,
	0xe2, 0x58, 0x6d, 0x64, 0xfd, 0x63, 0x0f, 0x2c, 0xe8, 0xb2, 0x5e, 0x74, 0xc0, 0x22, 0x71, 0x16,
	0x06, 0xf8, 0x35, 0x18, 0x9b, 0xbf, 0xae, 0x4b, 0x97, 0xb5, 0xeb, 0x76, 0x88, 0xeb, 0x26, 0x8e,
	0xeb, 0x1a, 0xe8, 0xd3, 0x51, 0x7b, 0x3a, 0xe7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x5e,
	0x9e, 0xd9, 0x83, 0x03, 0x00, 0x00,
}
