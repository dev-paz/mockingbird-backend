// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_negative_criterion.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A negative criterion for exclusions at the customer level.
type CustomerNegativeCriterion struct {
	// Immutable. The resource name of the customer negative criterion.
	// Customer negative criterion resource names have the form:
	//
	// `customers/{customer_id}/customerNegativeCriteria/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the criterion.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The customer negative criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CustomerNegativeCriterion_ContentLabel
	//	*CustomerNegativeCriterion_MobileApplication
	//	*CustomerNegativeCriterion_MobileAppCategory
	//	*CustomerNegativeCriterion_Placement
	//	*CustomerNegativeCriterion_YoutubeVideo
	//	*CustomerNegativeCriterion_YoutubeChannel
	Criterion            isCustomerNegativeCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *CustomerNegativeCriterion) Reset()         { *m = CustomerNegativeCriterion{} }
func (m *CustomerNegativeCriterion) String() string { return proto.CompactTextString(m) }
func (*CustomerNegativeCriterion) ProtoMessage()    {}
func (*CustomerNegativeCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_761ad3c7a7659e34, []int{0}
}

func (m *CustomerNegativeCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerNegativeCriterion.Unmarshal(m, b)
}
func (m *CustomerNegativeCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerNegativeCriterion.Marshal(b, m, deterministic)
}
func (m *CustomerNegativeCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerNegativeCriterion.Merge(m, src)
}
func (m *CustomerNegativeCriterion) XXX_Size() int {
	return xxx_messageInfo_CustomerNegativeCriterion.Size(m)
}
func (m *CustomerNegativeCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerNegativeCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerNegativeCriterion proto.InternalMessageInfo

func (m *CustomerNegativeCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerNegativeCriterion) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

type isCustomerNegativeCriterion_Criterion interface {
	isCustomerNegativeCriterion_Criterion()
}

type CustomerNegativeCriterion_ContentLabel struct {
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,4,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CustomerNegativeCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,5,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

type CustomerNegativeCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,6,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CustomerNegativeCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,7,opt,name=placement,proto3,oneof"`
}

type CustomerNegativeCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,8,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CustomerNegativeCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,9,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

func (*CustomerNegativeCriterion_ContentLabel) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_MobileApplication) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_MobileAppCategory) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_Placement) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_YoutubeVideo) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_YoutubeChannel) isCustomerNegativeCriterion_Criterion() {}

func (m *CustomerNegativeCriterion) GetCriterion() isCustomerNegativeCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerNegativeCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerNegativeCriterion_ContentLabel)(nil),
		(*CustomerNegativeCriterion_MobileApplication)(nil),
		(*CustomerNegativeCriterion_MobileAppCategory)(nil),
		(*CustomerNegativeCriterion_Placement)(nil),
		(*CustomerNegativeCriterion_YoutubeVideo)(nil),
		(*CustomerNegativeCriterion_YoutubeChannel)(nil),
	}
}

func init() {
	proto.RegisterType((*CustomerNegativeCriterion)(nil), "google.ads.googleads.v1.resources.CustomerNegativeCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_negative_criterion.proto", fileDescriptor_761ad3c7a7659e34)
}

var fileDescriptor_761ad3c7a7659e34 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x4e, 0x7f, 0xc8, 0xf6, 0x07, 0x61, 0x2e, 0x6e, 0x41, 0xa5, 0x45, 0xaa, 0xd4, 0x4b,
	0xd7, 0x4d, 0x28, 0x1c, 0x8c, 0x84, 0xe4, 0x44, 0xa8, 0x94, 0x9f, 0xaa, 0x8a, 0xaa, 0x20, 0x50,
	0x24, 0x6b, 0x6d, 0x4f, 0xdd, 0x15, 0xf6, 0xae, 0x65, 0xaf, 0x83, 0x22, 0xd4, 0x03, 0x07, 0x4e,
	0xbc, 0x05, 0x47, 0x1e, 0x85, 0xa7, 0xe8, 0xb9, 0xbc, 0x01, 0x27, 0x64, 0x7b, 0xd7, 0x49, 0xab,
	0x86, 0x16, 0x6e, 0x93, 0x7c, 0x7f, 0xe3, 0x19, 0xed, 0xa0, 0x6e, 0xc8, 0x79, 0x18, 0x81, 0x45,
	0x82, 0xcc, 0xaa, 0xca, 0xa2, 0x1a, 0xb6, 0xac, 0x14, 0x32, 0x9e, 0xa7, 0x3e, 0x64, 0x96, 0x9f,
	0x67, 0x82, 0xc7, 0x90, 0xba, 0x0c, 0x42, 0x22, 0xe8, 0x10, 0x5c, 0x3f, 0xa5, 0x02, 0x52, 0xca,
	0x19, 0x4e, 0x52, 0x2e, 0xb8, 0xb1, 0x51, 0x29, 0x31, 0x09, 0x32, 0x5c, 0x9b, 0xe0, 0x61, 0x0b,
	0xd7, 0x26, 0xab, 0xdb, 0xd3, 0x72, 0x7c, 0x1e, 0xc7, 0x9c, 0x59, 0xd2, 0x92, 0x54, 0x8e, 0xab,
	0xed, 0x69, 0x74, 0x60, 0x79, 0x9c, 0x59, 0x75, 0x03, 0xae, 0x18, 0x25, 0x20, 0x35, 0x0f, 0x95,
	0x26, 0xa1, 0xd6, 0x31, 0x85, 0x28, 0x70, 0x3d, 0x38, 0x21, 0x43, 0xca, 0x53, 0x49, 0x58, 0x99,
	0x20, 0xa8, 0xce, 0x24, 0xb4, 0x26, 0xa1, 0xf2, 0x97, 0x97, 0x1f, 0x5b, 0x9f, 0x52, 0x92, 0x24,
	0x90, 0x66, 0x12, 0x7f, 0x30, 0x21, 0x25, 0x8c, 0x71, 0x41, 0x04, 0xe5, 0x4c, 0xa2, 0x8f, 0x7e,
	0xcd, 0xa3, 0x95, 0xae, 0x9c, 0xd2, 0x81, 0x1c, 0x52, 0x57, 0xb5, 0x68, 0xb8, 0x68, 0x49, 0xa5,
	0xb9, 0x8c, 0xc4, 0x60, 0x6a, 0xeb, 0xda, 0x56, 0xb3, 0x63, 0x9f, 0x39, 0xb3, 0xbf, 0x9d, 0x5d,
	0xd4, 0x1e, 0x4f, 0x4c, 0x56, 0x09, 0xcd, 0xb0, 0xcf, 0x63, 0x6b, 0xaa, 0x65, 0x6f, 0x51, 0x19,
	0x1e, 0x90, 0x18, 0x8c, 0x1d, 0xa4, 0xd3, 0xc0, 0xd4, 0xd7, 0xb5, 0xad, 0x85, 0xf6, 0x7d, 0x69,
	0x82, 0xd5, 0x97, 0xe0, 0x7d, 0x26, 0x9e, 0xee, 0xf6, 0x49, 0x94, 0x43, 0xa7, 0x71, 0xe6, 0x34,
	0x7a, 0x3a, 0x0d, 0x8c, 0x77, 0x68, 0xa6, 0x18, 0x9c, 0xd9, 0x58, 0xd7, 0xb6, 0x96, 0xdb, 0xcf,
	0xf1, 0xb4, 0xfd, 0x95, 0xd3, 0xc6, 0x75, 0xee, 0xd1, 0x28, 0x81, 0x17, 0x2c, 0x8f, 0x2f, 0xfe,
	0x53, 0xd9, 0x96, 0x86, 0xc6, 0x00, 0x2d, 0xf9, 0x9c, 0x09, 0x60, 0xc2, 0x8d, 0x88, 0x07, 0x91,
	0x39, 0x53, 0x76, 0xb5, 0x33, 0x35, 0xa1, 0x5a, 0x3f, 0xee, 0x56, 0xa2, 0x37, 0x85, 0x66, 0x9f,
	0x1d, 0xf3, 0xc2, 0x73, 0xf6, 0xe5, 0xad, 0xde, 0xa2, 0x3f, 0x01, 0x18, 0x1f, 0x91, 0x11, 0x73,
	0x8f, 0x46, 0xe0, 0x92, 0x24, 0x89, 0xa8, 0x5f, 0x2e, 0xc1, 0x9c, 0x2d, 0x23, 0x9e, 0x5c, 0x17,
	0xf1, 0xb6, 0x54, 0x3a, 0x63, 0xe1, 0x64, 0xce, 0xdd, 0xf8, 0x32, 0x6a, 0x44, 0xe8, 0xde, 0x38,
	0xcc, 0xf5, 0x89, 0x80, 0x90, 0xa7, 0x23, 0x73, 0xee, 0x1f, 0xd3, 0xba, 0x52, 0x78, 0x75, 0x9a,
	0x42, 0x8d, 0x1e, 0x6a, 0x26, 0x11, 0xf1, 0x21, 0x06, 0x26, 0xcc, 0xf9, 0x32, 0x63, 0xfb, 0xba,
	0x8c, 0x43, 0x25, 0x98, 0xf4, 0x1e, 0xdb, 0x14, 0xcb, 0x18, 0xf1, 0x5c, 0xe4, 0x1e, 0xb8, 0x43,
	0x1a, 0x00, 0x37, 0x6f, 0xdf, 0x6c, 0x19, 0xef, 0x79, 0x7e, 0x94, 0x7b, 0xd0, 0x2f, 0x34, 0x17,
	0x96, 0x21, 0xdd, 0x4a, 0xc0, 0xf0, 0xd0, 0x1d, 0xe5, 0xee, 0x9f, 0x10, 0xc6, 0x20, 0x32, 0x9b,
	0xa5, 0x7f, 0xfb, 0x86, 0xfe, 0xdd, 0x4a, 0x35, 0x99, 0xb0, 0x2c, 0x1d, 0x25, 0x64, 0x7f, 0xd5,
	0xce, 0x9d, 0x2f, 0xda, 0xff, 0x3c, 0x11, 0xe3, 0xb5, 0x3a, 0x5b, 0x99, 0xf5, 0x59, 0x95, 0xa7,
	0xf5, 0x2d, 0xbb, 0xc4, 0x27, 0x63, 0xce, 0x15, 0x57, 0xee, 0xb4, 0xb3, 0x80, 0x9a, 0xf5, 0xaf,
	0xce, 0x37, 0x1d, 0x6d, 0xfa, 0x3c, 0xc6, 0xd7, 0x1e, 0xbd, 0xce, 0xda, 0xd4, 0xf6, 0x0e, 0x8b,
	0xd7, 0x79, 0xa8, 0x7d, 0x78, 0x25, 0x4d, 0x42, 0x1e, 0x11, 0x16, 0x62, 0x9e, 0x86, 0x56, 0x08,
	0xac, 0x7c, 0xbb, 0xd6, 0xf8, 0x53, 0xff, 0x72, 0x9d, 0x9f, 0xd5, 0xd5, 0x77, 0xbd, 0xb1, 0xe7,
	0x38, 0x3f, 0xf4, 0x8d, 0xbd, 0xca, 0xd2, 0x09, 0x32, 0x5c, 0x95, 0x45, 0xd5, 0x6f, 0xe1, 0x9e,
	0x62, 0xfe, 0x54, 0x9c, 0x81, 0x13, 0x64, 0x83, 0x9a, 0x33, 0xe8, 0xb7, 0x06, 0x35, 0xe7, 0x5c,
	0xdf, 0xac, 0x00, 0xdb, 0x76, 0x82, 0xcc, 0xb6, 0x6b, 0x96, 0x6d, 0xf7, 0x5b, 0xb6, 0x5d, 0xf3,
	0xbc, 0xb9, 0xb2, 0xd9, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0x7d, 0x4c, 0x28, 0x49,
	0x06, 0x00, 0x00,
}
