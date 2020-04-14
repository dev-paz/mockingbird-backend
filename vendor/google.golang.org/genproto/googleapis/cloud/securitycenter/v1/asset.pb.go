// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/asset.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Cloud Security Command Center's (Cloud SCC) representation of a Google Cloud
// Platform (GCP) resource.
//
// The Asset is a Cloud SCC resource that captures information about a single
// GCP resource. All modifications to an Asset are only within the context of
// Cloud SCC and don't affect the referenced GCP resource.
type Asset struct {
	// The relative resource name of this asset. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/assets/{asset_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Cloud SCC managed properties. These properties are managed by
	// Cloud SCC and cannot be modified by the user.
	SecurityCenterProperties *Asset_SecurityCenterProperties `protobuf:"bytes,2,opt,name=security_center_properties,json=securityCenterProperties,proto3" json:"security_center_properties,omitempty"`
	// Resource managed properties. These properties are managed and defined by
	// the GCP resource and cannot be modified by the user.
	ResourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=resource_properties,json=resourceProperties,proto3" json:"resource_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User specified security marks. These marks are entirely managed by the user
	// and come from the SecurityMarks resource that belongs to the asset.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the asset was created in Cloud SCC.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the asset was last updated, added, or deleted in Cloud
	// SCC.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// IAM Policy information associated with the GCP resource described by the
	// Cloud SCC asset. This information is managed and defined by the GCP
	// resource and cannot be modified by the user.
	IamPolicy            *Asset_IamPolicy `protobuf:"bytes,11,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_8afac45fcf79e390, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetSecurityCenterProperties() *Asset_SecurityCenterProperties {
	if m != nil {
		return m.SecurityCenterProperties
	}
	return nil
}

func (m *Asset) GetResourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.ResourceProperties
	}
	return nil
}

func (m *Asset) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Asset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Asset) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Asset) GetIamPolicy() *Asset_IamPolicy {
	if m != nil {
		return m.IamPolicy
	}
	return nil
}

// Cloud SCC managed properties. These properties are managed by Cloud SCC and
// cannot be modified by the user.
type Asset_SecurityCenterProperties struct {
	// The full resource name of the GCP resource this asset
	// represents. This field is immutable after create time. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The type of the GCP resource. Examples include: APPLICATION,
	// PROJECT, and ORGANIZATION. This is a case insensitive field defined by
	// Cloud SCC and/or the producer of the resource and is immutable
	// after create time.
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The full resource name of the immediate parent of the resource. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceParent string `protobuf:"bytes,3,opt,name=resource_parent,json=resourceParent,proto3" json:"resource_parent,omitempty"`
	// The full resource name of the project the resource belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceProject string `protobuf:"bytes,4,opt,name=resource_project,json=resourceProject,proto3" json:"resource_project,omitempty"`
	// Owners of the Google Cloud resource.
	ResourceOwners []string `protobuf:"bytes,5,rep,name=resource_owners,json=resourceOwners,proto3" json:"resource_owners,omitempty"`
	// The user defined display name for this resource.
	ResourceDisplayName string `protobuf:"bytes,6,opt,name=resource_display_name,json=resourceDisplayName,proto3" json:"resource_display_name,omitempty"`
	// The user defined display name for the parent of this resource.
	ResourceParentDisplayName string `protobuf:"bytes,7,opt,name=resource_parent_display_name,json=resourceParentDisplayName,proto3" json:"resource_parent_display_name,omitempty"`
	// The user defined display name for the project of this resource.
	ResourceProjectDisplayName string   `protobuf:"bytes,8,opt,name=resource_project_display_name,json=resourceProjectDisplayName,proto3" json:"resource_project_display_name,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Asset_SecurityCenterProperties) Reset()         { *m = Asset_SecurityCenterProperties{} }
func (m *Asset_SecurityCenterProperties) String() string { return proto.CompactTextString(m) }
func (*Asset_SecurityCenterProperties) ProtoMessage()    {}
func (*Asset_SecurityCenterProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_8afac45fcf79e390, []int{0, 0}
}

func (m *Asset_SecurityCenterProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Unmarshal(m, b)
}
func (m *Asset_SecurityCenterProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Marshal(b, m, deterministic)
}
func (m *Asset_SecurityCenterProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset_SecurityCenterProperties.Merge(m, src)
}
func (m *Asset_SecurityCenterProperties) XXX_Size() int {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Size(m)
}
func (m *Asset_SecurityCenterProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset_SecurityCenterProperties.DiscardUnknown(m)
}

var xxx_messageInfo_Asset_SecurityCenterProperties proto.InternalMessageInfo

func (m *Asset_SecurityCenterProperties) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceParent() string {
	if m != nil {
		return m.ResourceParent
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceProject() string {
	if m != nil {
		return m.ResourceProject
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceOwners() []string {
	if m != nil {
		return m.ResourceOwners
	}
	return nil
}

func (m *Asset_SecurityCenterProperties) GetResourceDisplayName() string {
	if m != nil {
		return m.ResourceDisplayName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceParentDisplayName() string {
	if m != nil {
		return m.ResourceParentDisplayName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceProjectDisplayName() string {
	if m != nil {
		return m.ResourceProjectDisplayName
	}
	return ""
}

// IAM Policy information associated with the GCP resource described by the
// Cloud SCC asset. This information is managed and defined by the GCP
// resource and cannot be modified by the user.
type Asset_IamPolicy struct {
	// The JSON representation of the Policy associated with the asset.
	// See https://cloud.google.com/iam/reference/rest/v1/Policy for format
	// details.
	PolicyBlob           string   `protobuf:"bytes,1,opt,name=policy_blob,json=policyBlob,proto3" json:"policy_blob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset_IamPolicy) Reset()         { *m = Asset_IamPolicy{} }
func (m *Asset_IamPolicy) String() string { return proto.CompactTextString(m) }
func (*Asset_IamPolicy) ProtoMessage()    {}
func (*Asset_IamPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_8afac45fcf79e390, []int{0, 1}
}

func (m *Asset_IamPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset_IamPolicy.Unmarshal(m, b)
}
func (m *Asset_IamPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset_IamPolicy.Marshal(b, m, deterministic)
}
func (m *Asset_IamPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset_IamPolicy.Merge(m, src)
}
func (m *Asset_IamPolicy) XXX_Size() int {
	return xxx_messageInfo_Asset_IamPolicy.Size(m)
}
func (m *Asset_IamPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset_IamPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_Asset_IamPolicy proto.InternalMessageInfo

func (m *Asset_IamPolicy) GetPolicyBlob() string {
	if m != nil {
		return m.PolicyBlob
	}
	return ""
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.cloud.securitycenter.v1.Asset")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1.Asset.ResourcePropertiesEntry")
	proto.RegisterType((*Asset_SecurityCenterProperties)(nil), "google.cloud.securitycenter.v1.Asset.SecurityCenterProperties")
	proto.RegisterType((*Asset_IamPolicy)(nil), "google.cloud.securitycenter.v1.Asset.IamPolicy")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/asset.proto", fileDescriptor_8afac45fcf79e390)
}

var fileDescriptor_8afac45fcf79e390 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x4e, 0xd4, 0x4c,
	0x18, 0x4e, 0x77, 0xf9, 0xdb, 0x77, 0x3f, 0x7e, 0x32, 0x5f, 0xd4, 0xd2, 0x20, 0xac, 0x70, 0x20,
	0x2a, 0xb6, 0x59, 0x38, 0x31, 0x25, 0x6a, 0x00, 0x8d, 0x31, 0x51, 0x24, 0x15, 0xf7, 0xc0, 0x60,
	0x36, 0xb3, 0x65, 0x6c, 0x2a, 0x6d, 0xa7, 0x99, 0x99, 0x62, 0x2a, 0x72, 0x21, 0x5e, 0x82, 0x5e,
	0x8a, 0x97, 0xe0, 0x25, 0x70, 0x15, 0xa6, 0x33, 0x6d, 0xb7, 0x5d, 0xb2, 0x2e, 0x47, 0xed, 0xbc,
	0xef, 0xf3, 0x3c, 0xef, 0xef, 0x0c, 0x3c, 0xf4, 0x28, 0xf5, 0x02, 0x62, 0xb9, 0x01, 0x4d, 0x4e,
	0x2d, 0x4e, 0xdc, 0x84, 0xf9, 0x22, 0x75, 0x49, 0x24, 0x08, 0xb3, 0xce, 0xbb, 0x16, 0xe6, 0x9c,
	0x08, 0x33, 0x66, 0x54, 0x50, 0xb4, 0xaa, 0xb0, 0xa6, 0xc4, 0x9a, 0x75, 0xac, 0x79, 0xde, 0x35,
	0x56, 0x72, 0x2d, 0x1c, 0xfb, 0x16, 0x8e, 0x22, 0x2a, 0xb0, 0xf0, 0x69, 0xc4, 0x15, 0xdb, 0x58,
	0xae, 0x78, 0x19, 0xe1, 0x34, 0x61, 0x2e, 0xc9, 0x5d, 0x3b, 0x13, 0x92, 0x28, 0x2c, 0xfd, 0x10,
	0xb3, 0xb3, 0x42, 0xaf, 0x88, 0x26, 0x4f, 0x83, 0xe4, 0xb3, 0xc5, 0x05, 0x4b, 0xdc, 0x3c, 0x57,
	0x63, 0x6d, 0xd4, 0x2b, 0xfc, 0x90, 0x70, 0x81, 0xc3, 0x58, 0x01, 0xd6, 0x7f, 0xb6, 0x60, 0x7a,
	0x2f, 0x2b, 0x0e, 0x21, 0x98, 0x8a, 0x70, 0x48, 0x74, 0xad, 0xa3, 0x6d, 0xb6, 0x1c, 0xf9, 0x8f,
	0xbe, 0x83, 0x51, 0x06, 0x55, 0x79, 0xf4, 0x63, 0x46, 0x63, 0xc2, 0x84, 0x4f, 0xb8, 0xde, 0xe8,
	0x68, 0x9b, 0xed, 0xed, 0x67, 0xe6, 0xbf, 0xfb, 0x61, 0x4a, 0x79, 0xf3, 0x7d, 0x6e, 0x3f, 0x90,
	0xf6, 0xa3, 0x52, 0xc5, 0xd1, 0xf9, 0x18, 0x0f, 0x8a, 0xe0, 0xff, 0xa2, 0x43, 0xd5, 0xb0, 0xb3,
	0x9d, 0xe6, 0x66, 0x7b, 0xfb, 0xe9, 0xcd, 0xc2, 0x3a, 0xb9, 0xc0, 0x50, 0xf6, 0x65, 0x24, 0x58,
	0xea, 0x20, 0x76, 0xcd, 0x81, 0x8e, 0x61, 0xa1, 0xde, 0x62, 0x7d, 0x4e, 0x56, 0xf8, 0x78, 0x52,
	0xa8, 0xa2, 0xb6, 0xb7, 0x19, 0xc9, 0x99, 0xe7, 0xd5, 0x23, 0xda, 0x85, 0xb6, 0xcb, 0x08, 0x16,
	0xa4, 0x9f, 0xf5, 0x5e, 0x6f, 0x49, 0x49, 0xa3, 0x90, 0x2c, 0x06, 0x63, 0x1e, 0x17, 0x83, 0x71,
	0x40, 0xc1, 0x33, 0x43, 0x46, 0x4e, 0xe2, 0xd3, 0x92, 0x0c, 0x93, 0xc9, 0x0a, 0x2e, 0xc9, 0x87,
	0x00, 0x3e, 0x0e, 0xfb, 0x31, 0x0d, 0x7c, 0x37, 0xd5, 0xdb, 0x92, 0x6b, 0xdd, 0xac, 0x6d, 0xaf,
	0x71, 0x78, 0x24, 0x69, 0x4e, 0xcb, 0x2f, 0x7e, 0x8d, 0x1f, 0x4d, 0xd0, 0xc7, 0x8d, 0x11, 0x6d,
	0xc0, 0x7c, 0x39, 0xac, 0xca, 0x1e, 0xfd, 0x57, 0x18, 0x0f, 0xb3, 0x7d, 0xaa, 0x82, 0x44, 0x1a,
	0x13, 0xb9, 0x42, 0x15, 0xd0, 0x71, 0x1a, 0x13, 0x74, 0x1f, 0x16, 0x87, 0x63, 0xc7, 0x8c, 0x44,
	0x42, 0x6f, 0x4a, 0xd8, 0x42, 0x39, 0x33, 0x69, 0x45, 0x0f, 0x60, 0xa9, 0xba, 0x1f, 0x5f, 0x88,
	0x2b, 0xf4, 0x29, 0x89, 0x5c, 0xac, 0x4c, 0x37, 0x33, 0xd7, 0x34, 0xe9, 0xd7, 0x88, 0x30, 0xae,
	0x4f, 0x77, 0x9a, 0x55, 0xcd, 0x77, 0xd2, 0x8a, 0xb6, 0xe1, 0x56, 0x09, 0x3c, 0xf5, 0x79, 0x1c,
	0xe0, 0x54, 0x95, 0x33, 0x23, 0x85, 0xcb, 0x85, 0x7c, 0xa1, 0x7c, 0xb2, 0xaa, 0xe7, 0xb0, 0x32,
	0x92, 0x70, 0x9d, 0x3a, 0x2b, 0xa9, 0xcb, 0xf5, 0xec, 0xab, 0x02, 0x7b, 0x70, 0x77, 0xb4, 0x90,
	0xba, 0xc2, 0x9c, 0x54, 0x30, 0x46, 0xaa, 0xaa, 0x48, 0x18, 0x5b, 0xd0, 0x2a, 0x67, 0x86, 0xd6,
	0xa0, 0xad, 0x86, 0xde, 0x1f, 0x04, 0x74, 0x90, 0x4f, 0x02, 0x94, 0x69, 0x3f, 0xa0, 0x03, 0xe3,
	0x13, 0xdc, 0x19, 0x73, 0x31, 0xd0, 0x12, 0x34, 0xcf, 0x48, 0x9a, 0x73, 0xb2, 0x5f, 0xb4, 0x05,
	0xd3, 0xe7, 0x38, 0x48, 0x48, 0x7e, 0xdf, 0x6f, 0x5f, 0xdb, 0xbe, 0x5e, 0xe6, 0x75, 0x14, 0xc8,
	0x6e, 0x3c, 0xd1, 0xec, 0x0f, 0x57, 0x7b, 0x0e, 0x6c, 0x8c, 0xec, 0x96, 0x62, 0xe1, 0xd8, 0xe7,
	0xa6, 0x4b, 0x43, 0x4b, 0x3d, 0x3a, 0x8f, 0x28, 0xf3, 0x70, 0xe4, 0x7f, 0x53, 0x4f, 0xa4, 0x75,
	0x51, 0x3d, 0x5e, 0xaa, 0x57, 0x97, 0x5b, 0x17, 0xf2, 0x7b, 0xb9, 0xff, 0x47, 0x83, 0x75, 0x97,
	0x86, 0x13, 0x36, 0xf8, 0x48, 0xfb, 0xf8, 0x26, 0x47, 0x78, 0x34, 0xc0, 0x91, 0x67, 0x52, 0xe6,
	0x59, 0x1e, 0x89, 0x64, 0xbe, 0xd6, 0x30, 0x8d, 0x71, 0xef, 0xec, 0x6e, 0xdd, 0xf2, 0xab, 0xb1,
	0xfa, 0x4a, 0xc9, 0x1d, 0xc8, 0x80, 0xf5, 0x1b, 0x60, 0xf6, 0xba, 0xbf, 0x0b, 0xc0, 0x89, 0x04,
	0x9c, 0xd4, 0x01, 0x27, 0xbd, 0xee, 0x55, 0xe3, 0x9e, 0x02, 0xd8, 0xb6, 0x44, 0xd8, 0x76, 0x1d,
	0x62, 0xdb, 0xbd, 0xee, 0x60, 0x46, 0xa6, 0xb7, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0x81, 0x34,
	0x4a, 0x76, 0x8a, 0x06, 0x00, 0x00,
}
