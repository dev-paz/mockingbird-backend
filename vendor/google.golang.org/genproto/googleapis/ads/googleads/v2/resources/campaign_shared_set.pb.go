// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_shared_set.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// CampaignSharedSets are used for managing the shared sets associated with a
// campaign.
type CampaignSharedSet struct {
	// Immutable. The resource name of the campaign shared set.
	// Campaign shared set resource names have the form:
	//
	// `customers/{customer_id}/campaignSharedSets/{campaign_id}~{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The campaign to which the campaign shared set belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Immutable. The shared set associated with the campaign. This may be a negative keyword
	// shared set of another customer. This customer should be a manager of the
	// other customer, otherwise the campaign shared set will exist but have no
	// serving effect. Only negative keyword shared sets can be associated with
	// Shopping campaigns. Only negative placement shared sets can be associated
	// with Display mobile app campaigns.
	SharedSet *wrappers.StringValue `protobuf:"bytes,4,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// Output only. The status of this campaign shared set. Read only.
	Status               enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *CampaignSharedSet) Reset()         { *m = CampaignSharedSet{} }
func (m *CampaignSharedSet) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSet) ProtoMessage()    {}
func (*CampaignSharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_799fb2d64e95762f, []int{0}
}

func (m *CampaignSharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSet.Unmarshal(m, b)
}
func (m *CampaignSharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSet.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSet.Merge(m, src)
}
func (m *CampaignSharedSet) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSet.Size(m)
}
func (m *CampaignSharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSet proto.InternalMessageInfo

func (m *CampaignSharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignSharedSet) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignSharedSet) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *CampaignSharedSet) GetStatus() enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignSharedSetStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignSharedSet)(nil), "google.ads.googleads.v2.resources.CampaignSharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_shared_set.proto", fileDescriptor_799fb2d64e95762f)
}

var fileDescriptor_799fb2d64e95762f = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0x66, 0x37, 0xb5, 0xd8, 0xf1, 0x03, 0xdc, 0x83, 0xc4, 0x52, 0x34, 0xa9, 0x16, 0x62, 0x91,
	0x19, 0x58, 0x6f, 0x53, 0x14, 0x66, 0x45, 0x0a, 0x1e, 0xa4, 0x6c, 0x20, 0x82, 0x04, 0x96, 0xc9,
	0xee, 0x74, 0xba, 0x98, 0x9d, 0x59, 0x66, 0x66, 0xe3, 0x41, 0x7a, 0xf2, 0x67, 0x78, 0xf3, 0xe8,
	0x4f, 0xf1, 0x57, 0xf4, 0xdc, 0x9f, 0xe0, 0x49, 0xb2, 0x3b, 0x33, 0x29, 0xc4, 0x18, 0xbd, 0x3d,
	0xe1, 0x7d, 0x3e, 0xde, 0x3c, 0xfb, 0x0e, 0x38, 0xe1, 0x52, 0xf2, 0x39, 0x43, 0xb4, 0xd0, 0xa8,
	0x83, 0x4b, 0xb4, 0x88, 0x91, 0x62, 0x5a, 0x36, 0x2a, 0x67, 0x1a, 0xe5, 0xb4, 0xaa, 0x69, 0xc9,
	0x45, 0xa6, 0x2f, 0xa8, 0x62, 0x45, 0xa6, 0x99, 0x81, 0xb5, 0x92, 0x46, 0x46, 0xc3, 0x4e, 0x01,
	0x69, 0xa1, 0xa1, 0x17, 0xc3, 0x45, 0x0c, 0xbd, 0x78, 0xff, 0xf5, 0x26, 0x7f, 0x26, 0x9a, 0xea,
	0x8f, 0xde, 0x99, 0x36, 0xd4, 0x34, 0xba, 0x8b, 0xd8, 0x7f, 0xe2, 0xf4, 0x75, 0x89, 0xce, 0x4b,
	0x36, 0x2f, 0xb2, 0x19, 0xbb, 0xa0, 0x8b, 0x52, 0x2a, 0x4b, 0x78, 0x74, 0x83, 0xe0, 0x62, 0xed,
	0xe8, 0xb1, 0x1d, 0xb5, 0xbf, 0x66, 0xcd, 0x39, 0xfa, 0xac, 0x68, 0x5d, 0x33, 0xe5, 0xbc, 0x0f,
	0x6e, 0x48, 0xa9, 0x10, 0xd2, 0x50, 0x53, 0x4a, 0x61, 0xa7, 0x87, 0xdf, 0x76, 0xc0, 0x83, 0x37,
	0x76, 0xbd, 0x71, 0xbb, 0xdd, 0x98, 0x99, 0xe8, 0x03, 0xb8, 0xe7, 0x52, 0x32, 0x41, 0x2b, 0xd6,
	0x0f, 0x06, 0xc1, 0x68, 0x2f, 0x89, 0xaf, 0xc8, 0xad, 0x5f, 0xe4, 0x05, 0x38, 0x5e, 0xd5, 0x60,
	0x51, 0x5d, 0x6a, 0x98, 0xcb, 0x0a, 0xad, 0x59, 0xa5, 0x77, 0x9d, 0xd1, 0x7b, 0x5a, 0xb1, 0x28,
	0x07, 0xb7, 0x5d, 0x19, 0xfd, 0xde, 0x20, 0x18, 0xdd, 0x89, 0x0f, 0xac, 0x05, 0x74, 0xfb, 0xc3,
	0xb1, 0x51, 0xa5, 0xe0, 0x13, 0x3a, 0x6f, 0x58, 0xf2, 0xbc, 0x4d, 0x7c, 0x0a, 0x86, 0x5b, 0x13,
	0x53, 0x6f, 0x1c, 0x71, 0x00, 0x56, 0x45, 0xf7, 0x77, 0xfe, 0x21, 0xe6, 0xb8, 0x8d, 0x79, 0x06,
	0x0e, 0x37, 0xc6, 0xac, 0xfe, 0xd0, 0x9e, 0xf6, 0x35, 0x49, 0xb0, 0xdb, 0x7d, 0xc6, 0x7e, 0x38,
	0x08, 0x46, 0xf7, 0xe3, 0x14, 0x6e, 0x3a, 0x95, 0xf6, 0x0e, 0xe0, 0x5a, 0x3b, 0xe3, 0x56, 0xfd,
	0x56, 0x34, 0xd5, 0xa6, 0x59, 0xd2, 0xbb, 0x22, 0xbd, 0xd4, 0xc6, 0x60, 0x71, 0x4d, 0x3e, 0xfd,
	0x4f, 0xfb, 0xd1, 0xab, 0xbc, 0xd1, 0x46, 0x56, 0x4c, 0x69, 0xf4, 0xc5, 0xc1, 0x4b, 0x7f, 0x8f,
	0x9e, 0xb7, 0x9c, 0xae, 0xdf, 0xe8, 0x65, 0xf2, 0x35, 0x04, 0x47, 0xb9, 0xac, 0xe0, 0xd6, 0x17,
	0x90, 0x3c, 0x5c, 0xcb, 0x3e, 0x5b, 0x16, 0x7d, 0x16, 0x7c, 0x7c, 0x67, 0xc5, 0x5c, 0xce, 0xa9,
	0xe0, 0x50, 0x2a, 0x8e, 0x38, 0x13, 0xed, 0x67, 0x40, 0xab, 0xfd, 0xff, 0xf2, 0x34, 0x4f, 0x3c,
	0xfa, 0x1e, 0xf6, 0x4e, 0x09, 0xf9, 0x11, 0x0e, 0x4f, 0x3b, 0x4b, 0x52, 0x68, 0xd8, 0xc1, 0x25,
	0x9a, 0xc4, 0x30, 0x75, 0xcc, 0x9f, 0x8e, 0x33, 0x25, 0x85, 0x9e, 0x7a, 0xce, 0x74, 0x12, 0x4f,
	0x3d, 0xe7, 0x3a, 0x3c, 0xea, 0x06, 0x18, 0x93, 0x42, 0x63, 0xec, 0x59, 0x18, 0x4f, 0x62, 0x8c,
	0x3d, 0x6f, 0xb6, 0xdb, 0x2e, 0xfb, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x1c, 0x2f,
	0x5c, 0x46, 0x04, 0x00, 0x00,
}
