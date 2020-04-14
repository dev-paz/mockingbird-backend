// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/merchant_center_link_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [MerchantCenterLinkService.ListMerchantCenterLinks][google.ads.googleads.v2.services.MerchantCenterLinkService.ListMerchantCenterLinks].
type ListMerchantCenterLinksRequest struct {
	// Required. The ID of the customer onto which to apply the Merchant Center link list
	// operation.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMerchantCenterLinksRequest) Reset()         { *m = ListMerchantCenterLinksRequest{} }
func (m *ListMerchantCenterLinksRequest) String() string { return proto.CompactTextString(m) }
func (*ListMerchantCenterLinksRequest) ProtoMessage()    {}
func (*ListMerchantCenterLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{0}
}

func (m *ListMerchantCenterLinksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Unmarshal(m, b)
}
func (m *ListMerchantCenterLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Marshal(b, m, deterministic)
}
func (m *ListMerchantCenterLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMerchantCenterLinksRequest.Merge(m, src)
}
func (m *ListMerchantCenterLinksRequest) XXX_Size() int {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Size(m)
}
func (m *ListMerchantCenterLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMerchantCenterLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMerchantCenterLinksRequest proto.InternalMessageInfo

func (m *ListMerchantCenterLinksRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for [MerchantCenterLinkService.ListMerchantCenterLinks][google.ads.googleads.v2.services.MerchantCenterLinkService.ListMerchantCenterLinks].
type ListMerchantCenterLinksResponse struct {
	// Merchant Center links available for the requested customer
	MerchantCenterLinks  []*resources.MerchantCenterLink `protobuf:"bytes,1,rep,name=merchant_center_links,json=merchantCenterLinks,proto3" json:"merchant_center_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ListMerchantCenterLinksResponse) Reset()         { *m = ListMerchantCenterLinksResponse{} }
func (m *ListMerchantCenterLinksResponse) String() string { return proto.CompactTextString(m) }
func (*ListMerchantCenterLinksResponse) ProtoMessage()    {}
func (*ListMerchantCenterLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{1}
}

func (m *ListMerchantCenterLinksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Unmarshal(m, b)
}
func (m *ListMerchantCenterLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Marshal(b, m, deterministic)
}
func (m *ListMerchantCenterLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMerchantCenterLinksResponse.Merge(m, src)
}
func (m *ListMerchantCenterLinksResponse) XXX_Size() int {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Size(m)
}
func (m *ListMerchantCenterLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMerchantCenterLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMerchantCenterLinksResponse proto.InternalMessageInfo

func (m *ListMerchantCenterLinksResponse) GetMerchantCenterLinks() []*resources.MerchantCenterLink {
	if m != nil {
		return m.MerchantCenterLinks
	}
	return nil
}

// Request message for [MerchantCenterLinkService.GetMerchantCenterLink][google.ads.googleads.v2.services.MerchantCenterLinkService.GetMerchantCenterLink].
type GetMerchantCenterLinkRequest struct {
	// Required. Resource name of the Merchant Center link.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMerchantCenterLinkRequest) Reset()         { *m = GetMerchantCenterLinkRequest{} }
func (m *GetMerchantCenterLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetMerchantCenterLinkRequest) ProtoMessage()    {}
func (*GetMerchantCenterLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{2}
}

func (m *GetMerchantCenterLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Unmarshal(m, b)
}
func (m *GetMerchantCenterLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Marshal(b, m, deterministic)
}
func (m *GetMerchantCenterLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMerchantCenterLinkRequest.Merge(m, src)
}
func (m *GetMerchantCenterLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Size(m)
}
func (m *GetMerchantCenterLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMerchantCenterLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMerchantCenterLinkRequest proto.InternalMessageInfo

func (m *GetMerchantCenterLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MerchantCenterLinkService.MutateMerchantCenterLink][google.ads.googleads.v2.services.MerchantCenterLinkService.MutateMerchantCenterLink].
type MutateMerchantCenterLinkRequest struct {
	// Required. The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform on the link
	Operation            *MerchantCenterLinkOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateMerchantCenterLinkRequest) Reset()         { *m = MutateMerchantCenterLinkRequest{} }
func (m *MutateMerchantCenterLinkRequest) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkRequest) ProtoMessage()    {}
func (*MutateMerchantCenterLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{3}
}

func (m *MutateMerchantCenterLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkRequest.Merge(m, src)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Size(m)
}
func (m *MutateMerchantCenterLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkRequest proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateMerchantCenterLinkRequest) GetOperation() *MerchantCenterLinkOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single update on a Merchant Center link.
type MerchantCenterLinkOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The operation to perform
	//
	// Types that are valid to be assigned to Operation:
	//	*MerchantCenterLinkOperation_Update
	//	*MerchantCenterLinkOperation_Remove
	Operation            isMerchantCenterLinkOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MerchantCenterLinkOperation) Reset()         { *m = MerchantCenterLinkOperation{} }
func (m *MerchantCenterLinkOperation) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLinkOperation) ProtoMessage()    {}
func (*MerchantCenterLinkOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{4}
}

func (m *MerchantCenterLinkOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLinkOperation.Unmarshal(m, b)
}
func (m *MerchantCenterLinkOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLinkOperation.Marshal(b, m, deterministic)
}
func (m *MerchantCenterLinkOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLinkOperation.Merge(m, src)
}
func (m *MerchantCenterLinkOperation) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLinkOperation.Size(m)
}
func (m *MerchantCenterLinkOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLinkOperation.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLinkOperation proto.InternalMessageInfo

func (m *MerchantCenterLinkOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isMerchantCenterLinkOperation_Operation interface {
	isMerchantCenterLinkOperation_Operation()
}

type MerchantCenterLinkOperation_Update struct {
	Update *resources.MerchantCenterLink `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type MerchantCenterLinkOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*MerchantCenterLinkOperation_Update) isMerchantCenterLinkOperation_Operation() {}

func (*MerchantCenterLinkOperation_Remove) isMerchantCenterLinkOperation_Operation() {}

func (m *MerchantCenterLinkOperation) GetOperation() isMerchantCenterLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MerchantCenterLinkOperation) GetUpdate() *resources.MerchantCenterLink {
	if x, ok := m.GetOperation().(*MerchantCenterLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *MerchantCenterLinkOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*MerchantCenterLinkOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MerchantCenterLinkOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MerchantCenterLinkOperation_Update)(nil),
		(*MerchantCenterLinkOperation_Remove)(nil),
	}
}

// Response message for Merchant Center link mutate.
type MutateMerchantCenterLinkResponse struct {
	// Result for the mutate.
	Result               *MutateMerchantCenterLinkResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateMerchantCenterLinkResponse) Reset()         { *m = MutateMerchantCenterLinkResponse{} }
func (m *MutateMerchantCenterLinkResponse) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkResponse) ProtoMessage()    {}
func (*MutateMerchantCenterLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{5}
}

func (m *MutateMerchantCenterLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkResponse.Merge(m, src)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Size(m)
}
func (m *MutateMerchantCenterLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkResponse proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkResponse) GetResult() *MutateMerchantCenterLinkResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the Merchant Center link mutate.
type MutateMerchantCenterLinkResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateMerchantCenterLinkResult) Reset()         { *m = MutateMerchantCenterLinkResult{} }
func (m *MutateMerchantCenterLinkResult) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkResult) ProtoMessage()    {}
func (*MutateMerchantCenterLinkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd0dc4f0ea434cdd, []int{6}
}

func (m *MutateMerchantCenterLinkResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkResult.Merge(m, src)
}
func (m *MutateMerchantCenterLinkResult) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Size(m)
}
func (m *MutateMerchantCenterLinkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkResult proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListMerchantCenterLinksRequest)(nil), "google.ads.googleads.v2.services.ListMerchantCenterLinksRequest")
	proto.RegisterType((*ListMerchantCenterLinksResponse)(nil), "google.ads.googleads.v2.services.ListMerchantCenterLinksResponse")
	proto.RegisterType((*GetMerchantCenterLinkRequest)(nil), "google.ads.googleads.v2.services.GetMerchantCenterLinkRequest")
	proto.RegisterType((*MutateMerchantCenterLinkRequest)(nil), "google.ads.googleads.v2.services.MutateMerchantCenterLinkRequest")
	proto.RegisterType((*MerchantCenterLinkOperation)(nil), "google.ads.googleads.v2.services.MerchantCenterLinkOperation")
	proto.RegisterType((*MutateMerchantCenterLinkResponse)(nil), "google.ads.googleads.v2.services.MutateMerchantCenterLinkResponse")
	proto.RegisterType((*MutateMerchantCenterLinkResult)(nil), "google.ads.googleads.v2.services.MutateMerchantCenterLinkResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/merchant_center_link_service.proto", fileDescriptor_dd0dc4f0ea434cdd)
}

var fileDescriptor_dd0dc4f0ea434cdd = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x41, 0x6b, 0xdb, 0x4a,
	0x10, 0x7e, 0x92, 0xc1, 0x90, 0xf5, 0x7b, 0x97, 0x7d, 0x84, 0x38, 0x4e, 0x48, 0x8c, 0x9a, 0x43,
	0x70, 0x5b, 0x89, 0x2a, 0x84, 0x16, 0xa5, 0x2e, 0x5d, 0x87, 0xc6, 0x29, 0xc4, 0x4d, 0x70, 0x21,
	0x84, 0xd6, 0xe0, 0x2a, 0xd6, 0xc6, 0x11, 0xb1, 0xb4, 0xae, 0x76, 0x65, 0x0a, 0x69, 0x2e, 0x2d,
	0xbd, 0xf5, 0xd6, 0x7f, 0xd0, 0x5b, 0xfb, 0x53, 0x72, 0x6c, 0x6f, 0x81, 0x42, 0x0e, 0x85, 0x42,
	0x7f, 0x42, 0x0f, 0xa5, 0x48, 0xbb, 0x6b, 0x2b, 0x44, 0x72, 0xd2, 0xe4, 0x36, 0xbb, 0x33, 0xfa,
	0x66, 0xbe, 0x99, 0x6f, 0x16, 0x81, 0xd5, 0x2e, 0x21, 0xdd, 0x1e, 0x36, 0x6c, 0x87, 0x1a, 0xdc,
	0x8c, 0xac, 0x81, 0x69, 0x50, 0x1c, 0x0c, 0xdc, 0x0e, 0xa6, 0x86, 0x87, 0x83, 0xce, 0xbe, 0xed,
	0xb3, 0x76, 0x07, 0xfb, 0x0c, 0x07, 0xed, 0x9e, 0xeb, 0x1f, 0xb4, 0x85, 0x57, 0xef, 0x07, 0x84,
	0x11, 0x58, 0xe6, 0x5f, 0xea, 0xb6, 0x43, 0xf5, 0x21, 0x88, 0x3e, 0x30, 0x75, 0x09, 0x52, 0xba,
	0x9f, 0x95, 0x26, 0xc0, 0x94, 0x84, 0x41, 0x56, 0x1e, 0x8e, 0x5f, 0x9a, 0x95, 0x5f, 0xf7, 0x5d,
	0xc3, 0xf6, 0x7d, 0xc2, 0x6c, 0xe6, 0x12, 0x9f, 0x0a, 0xef, 0x54, 0xc2, 0xdb, 0xe9, 0xb9, 0xd8,
	0x67, 0xc2, 0x31, 0x9f, 0x70, 0xec, 0xb9, 0xb8, 0xe7, 0xb4, 0x77, 0xf1, 0xbe, 0x3d, 0x70, 0x49,
	0x20, 0x02, 0xa6, 0x13, 0x01, 0xb2, 0x10, 0xe1, 0x12, 0x94, 0x8c, 0xf8, 0xb4, 0x1b, 0xee, 0x09,
	0x00, 0xcf, 0xa6, 0xa2, 0x28, 0x6d, 0x0d, 0xcc, 0x6d, 0xb8, 0x94, 0x35, 0x44, 0xd9, 0xab, 0x71,
	0xd5, 0x1b, 0xae, 0x7f, 0x40, 0x9b, 0xf8, 0x65, 0x88, 0x29, 0x83, 0x0b, 0xa0, 0xd0, 0x09, 0x29,
	0x23, 0x1e, 0x0e, 0xda, 0xae, 0x53, 0x54, 0xca, 0xca, 0xe2, 0x44, 0x2d, 0x77, 0x8a, 0xd4, 0x26,
	0x90, 0xf7, 0x8f, 0x1d, 0xed, 0xbd, 0x02, 0xe6, 0x33, 0x81, 0x68, 0x9f, 0xf8, 0x14, 0x43, 0x17,
	0x4c, 0xa6, 0xb5, 0x87, 0x16, 0x95, 0x72, 0x6e, 0xb1, 0x60, 0x2e, 0xeb, 0x59, 0x03, 0x18, 0xb6,
	0x57, 0x3f, 0x0f, 0xdf, 0xfc, 0xdf, 0x3b, 0x9f, 0x52, 0x7b, 0x05, 0x66, 0xeb, 0x38, 0xa5, 0x18,
	0x49, 0x6a, 0x07, 0xfc, 0x27, 0x41, 0xdb, 0xbe, 0xed, 0x61, 0x41, 0x6b, 0xe9, 0x14, 0xa9, 0xbf,
	0xd0, 0x6d, 0x70, 0x73, 0x94, 0x5e, 0x58, 0x7d, 0x97, 0xea, 0x1d, 0xe2, 0x19, 0x29, 0x90, 0xff,
	0x4a, 0xa4, 0x27, 0xb6, 0x87, 0xb5, 0x4f, 0x0a, 0x98, 0x6f, 0x84, 0xcc, 0x66, 0x38, 0x3b, 0xfb,
	0xa5, 0x5a, 0x0a, 0x5f, 0x80, 0x09, 0xd2, 0xc7, 0x41, 0xac, 0x92, 0xa2, 0x5a, 0x56, 0x16, 0x0b,
	0x66, 0x55, 0xbf, 0x48, 0xa3, 0x29, 0x1d, 0xda, 0x94, 0x20, 0x3c, 0xc5, 0x08, 0x54, 0xfb, 0xa2,
	0x80, 0x99, 0x31, 0xf1, 0x70, 0x05, 0x14, 0xc2, 0xbe, 0x63, 0x33, 0x1c, 0x2b, 0xa6, 0x98, 0x8b,
	0x6b, 0x28, 0xc9, 0x1a, 0xa4, 0xa8, 0xf4, 0xb5, 0x48, 0x54, 0x0d, 0x9b, 0x1e, 0x34, 0x01, 0x0f,
	0x8f, 0x6c, 0xb8, 0x09, 0xf2, 0xfc, 0x14, 0xf3, 0xbb, 0xea, 0x78, 0xd7, 0xff, 0x69, 0x0a, 0x18,
	0x58, 0x04, 0xf9, 0x00, 0x7b, 0x64, 0x80, 0xe3, 0x66, 0x4c, 0x44, 0x1e, 0x7e, 0xae, 0x15, 0x12,
	0x9d, 0xd2, 0x5e, 0x83, 0x72, 0x76, 0xff, 0x85, 0x12, 0x77, 0x22, 0x28, 0x1a, 0xf6, 0x98, 0xe8,
	0xeb, 0xc3, 0x4b, 0xf4, 0x35, 0x1b, 0x33, 0xec, 0xb1, 0xa6, 0xc0, 0xd3, 0x1e, 0x81, 0xb9, 0xf1,
	0x91, 0xf0, 0x46, 0xaa, 0xf4, 0xce, 0xaa, 0xc8, 0x7c, 0x9b, 0x07, 0xd3, 0xe7, 0x11, 0x9e, 0xf2,
	0x62, 0xe0, 0x0f, 0x05, 0x4c, 0x65, 0x2c, 0x1b, 0xbc, 0x04, 0x95, 0xf1, 0x0b, 0x5f, 0x42, 0xd7,
	0x40, 0xe0, 0xfd, 0xd5, 0xea, 0x27, 0x28, 0xa9, 0xf0, 0x37, 0x5f, 0xbf, 0x7f, 0x50, 0x97, 0xe0,
	0x9d, 0xe8, 0xa9, 0x94, 0xd7, 0xd4, 0x38, 0x4c, 0x44, 0x54, 0x2b, 0x47, 0x46, 0xca, 0x1e, 0xc3,
	0x6f, 0x0a, 0x98, 0x4c, 0x5d, 0x64, 0xf8, 0xe0, 0xe2, 0x2a, 0xc7, 0xbd, 0x00, 0xa5, 0xab, 0xc9,
	0x51, 0x6b, 0x9c, 0xa0, 0xb3, 0xe3, 0x8b, 0xb9, 0xdd, 0x85, 0xcb, 0x11, 0xb7, 0xc3, 0x33, 0x9e,
	0xea, 0x88, 0x6a, 0x25, 0x8d, 0x9c, 0x51, 0x39, 0x82, 0xbf, 0x15, 0x50, 0xcc, 0xd2, 0x0b, 0x44,
	0xd7, 0x51, 0x25, 0x67, 0x59, 0xbb, 0x96, 0xb0, 0xf9, 0x30, 0x9f, 0x9f, 0xa0, 0xc9, 0xc4, 0xa8,
	0x6e, 0x0d, 0x57, 0x2d, 0xa6, 0x5e, 0xd5, 0xee, 0xfd, 0xf5, 0x58, 0x2d, 0x2f, 0xce, 0x69, 0x29,
	0x95, 0xd2, 0xcc, 0x31, 0x2a, 0x66, 0x3d, 0xb6, 0xb5, 0x77, 0x2a, 0x58, 0xe8, 0x10, 0xef, 0x42,
	0x0e, 0xb5, 0xb9, 0xcc, 0x5d, 0xd9, 0x8a, 0x1e, 0xa9, 0x2d, 0xe5, 0xd9, 0xba, 0xc0, 0xe8, 0x92,
	0x9e, 0xed, 0x77, 0x75, 0x12, 0x74, 0x8d, 0x2e, 0xf6, 0xe3, 0x27, 0xcc, 0x18, 0x65, 0xcd, 0xfe,
	0x81, 0x58, 0x91, 0xc6, 0x47, 0x35, 0x57, 0x47, 0xe8, 0xb3, 0x5a, 0xae, 0x73, 0x40, 0xe4, 0x50,
	0x9d, 0x9b, 0x91, 0xb5, 0x6d, 0xea, 0x22, 0x31, 0x3d, 0x96, 0x21, 0x2d, 0xe4, 0xd0, 0xd6, 0x30,
	0xa4, 0xb5, 0x6d, 0xb6, 0x64, 0xc8, 0x4f, 0x75, 0x81, 0xdf, 0x5b, 0x16, 0x72, 0xa8, 0x65, 0x0d,
	0x83, 0x2c, 0x6b, 0xdb, 0xb4, 0x2c, 0x19, 0xb6, 0x9b, 0x8f, 0xeb, 0x5c, 0xfa, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0x59, 0x57, 0x8a, 0xe7, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MerchantCenterLinkServiceClient is the client API for MerchantCenterLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchantCenterLinkServiceClient interface {
	// Returns Merchant Center links available for this customer.
	ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error)
}

type merchantCenterLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantCenterLinkServiceClient(cc grpc.ClientConnInterface) MerchantCenterLinkServiceClient {
	return &merchantCenterLinkServiceClient{cc}
}

func (c *merchantCenterLinkServiceClient) ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error) {
	out := new(ListMerchantCenterLinksResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MerchantCenterLinkService/ListMerchantCenterLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error) {
	out := new(resources.MerchantCenterLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MerchantCenterLinkService/GetMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error) {
	out := new(MutateMerchantCenterLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MerchantCenterLinkService/MutateMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantCenterLinkServiceServer is the server API for MerchantCenterLinkService service.
type MerchantCenterLinkServiceServer interface {
	// Returns Merchant Center links available for this customer.
	ListMerchantCenterLinks(context.Context, *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	GetMerchantCenterLink(context.Context, *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	MutateMerchantCenterLink(context.Context, *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error)
}

// UnimplementedMerchantCenterLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMerchantCenterLinkServiceServer struct {
}

func (*UnimplementedMerchantCenterLinkServiceServer) ListMerchantCenterLinks(ctx context.Context, req *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchantCenterLinks not implemented")
}
func (*UnimplementedMerchantCenterLinkServiceServer) GetMerchantCenterLink(ctx context.Context, req *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantCenterLink not implemented")
}
func (*UnimplementedMerchantCenterLinkServiceServer) MutateMerchantCenterLink(ctx context.Context, req *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateMerchantCenterLink not implemented")
}

func RegisterMerchantCenterLinkServiceServer(s *grpc.Server, srv MerchantCenterLinkServiceServer) {
	s.RegisterService(&_MerchantCenterLinkService_serviceDesc, srv)
}

func _MerchantCenterLinkService_ListMerchantCenterLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantCenterLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MerchantCenterLinkService/ListMerchantCenterLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, req.(*ListMerchantCenterLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_GetMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MerchantCenterLinkService/GetMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, req.(*GetMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_MutateMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MerchantCenterLinkService/MutateMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, req.(*MutateMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MerchantCenterLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.MerchantCenterLinkService",
	HandlerType: (*MerchantCenterLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMerchantCenterLinks",
			Handler:    _MerchantCenterLinkService_ListMerchantCenterLinks_Handler,
		},
		{
			MethodName: "GetMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_GetMerchantCenterLink_Handler,
		},
		{
			MethodName: "MutateMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_MutateMerchantCenterLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/merchant_center_link_service.proto",
}
