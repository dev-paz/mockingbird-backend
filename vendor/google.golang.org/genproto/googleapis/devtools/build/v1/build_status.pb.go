// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/devtools/build/v1/build_status.proto

package build

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The end result of the Build.
type BuildStatus_Result int32

const (
	// Unspecified or unknown.
	BuildStatus_UNKNOWN_STATUS BuildStatus_Result = 0
	// Build was successful and tests (if requested) all pass.
	BuildStatus_COMMAND_SUCCEEDED BuildStatus_Result = 1
	// Build error and/or test failure.
	BuildStatus_COMMAND_FAILED BuildStatus_Result = 2
	// Unable to obtain a result due to input provided by the user.
	BuildStatus_USER_ERROR BuildStatus_Result = 3
	// Unable to obtain a result due to a failure within the build system.
	BuildStatus_SYSTEM_ERROR BuildStatus_Result = 4
	// Build required too many resources, such as build tool RAM.
	BuildStatus_RESOURCE_EXHAUSTED BuildStatus_Result = 5
	// An invocation attempt time exceeded its deadline.
	BuildStatus_INVOCATION_DEADLINE_EXCEEDED BuildStatus_Result = 6
	// Build request time exceeded the request_deadline
	BuildStatus_REQUEST_DEADLINE_EXCEEDED BuildStatus_Result = 8
	// The build was cancelled by a call to CancelBuild.
	BuildStatus_CANCELLED BuildStatus_Result = 7
)

// Enum value maps for BuildStatus_Result.
var (
	BuildStatus_Result_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "COMMAND_SUCCEEDED",
		2: "COMMAND_FAILED",
		3: "USER_ERROR",
		4: "SYSTEM_ERROR",
		5: "RESOURCE_EXHAUSTED",
		6: "INVOCATION_DEADLINE_EXCEEDED",
		8: "REQUEST_DEADLINE_EXCEEDED",
		7: "CANCELLED",
	}
	BuildStatus_Result_value = map[string]int32{
		"UNKNOWN_STATUS":               0,
		"COMMAND_SUCCEEDED":            1,
		"COMMAND_FAILED":               2,
		"USER_ERROR":                   3,
		"SYSTEM_ERROR":                 4,
		"RESOURCE_EXHAUSTED":           5,
		"INVOCATION_DEADLINE_EXCEEDED": 6,
		"REQUEST_DEADLINE_EXCEEDED":    8,
		"CANCELLED":                    7,
	}
)

func (x BuildStatus_Result) Enum() *BuildStatus_Result {
	p := new(BuildStatus_Result)
	*p = x
	return p
}

func (x BuildStatus_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildStatus_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_build_v1_build_status_proto_enumTypes[0].Descriptor()
}

func (BuildStatus_Result) Type() protoreflect.EnumType {
	return &file_google_devtools_build_v1_build_status_proto_enumTypes[0]
}

func (x BuildStatus_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildStatus_Result.Descriptor instead.
func (BuildStatus_Result) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_build_v1_build_status_proto_rawDescGZIP(), []int{0, 0}
}

// Status used for both invocation attempt and overall build completion.
type BuildStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The end result.
	Result BuildStatus_Result `protobuf:"varint,1,opt,name=result,proto3,enum=google.devtools.build.v1.BuildStatus_Result" json:"result,omitempty"`
	// Final invocation ID of the build, if there was one.
	// This field is only set on a status in BuildFinished event.
	FinalInvocationId string `protobuf:"bytes,3,opt,name=final_invocation_id,json=finalInvocationId,proto3" json:"final_invocation_id,omitempty"`
	// Build tool exit code. Integer value returned by the executed build tool.
	// Might not be available in some cases, e.g., a build timeout.
	BuildToolExitCode *wrappers.Int32Value `protobuf:"bytes,4,opt,name=build_tool_exit_code,json=buildToolExitCode,proto3" json:"build_tool_exit_code,omitempty"`
	// Fine-grained diagnostic information to complement the status.
	Details *any.Any `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *BuildStatus) Reset() {
	*x = BuildStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_build_v1_build_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildStatus) ProtoMessage() {}

func (x *BuildStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_build_v1_build_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildStatus.ProtoReflect.Descriptor instead.
func (*BuildStatus) Descriptor() ([]byte, []int) {
	return file_google_devtools_build_v1_build_status_proto_rawDescGZIP(), []int{0}
}

func (x *BuildStatus) GetResult() BuildStatus_Result {
	if x != nil {
		return x.Result
	}
	return BuildStatus_UNKNOWN_STATUS
}

func (x *BuildStatus) GetFinalInvocationId() string {
	if x != nil {
		return x.FinalInvocationId
	}
	return ""
}

func (x *BuildStatus) GetBuildToolExitCode() *wrappers.Int32Value {
	if x != nil {
		return x.BuildToolExitCode
	}
	return nil
}

func (x *BuildStatus) GetDetails() *any.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_google_devtools_build_v1_build_status_proto protoreflect.FileDescriptor

var file_google_devtools_build_v1_build_status_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd5, 0x03, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x14, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x11, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x78,
	0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45,
	0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e,
	0x56, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e,
	0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x42, 0x74, 0x0a, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0xf8, 0x01, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_build_v1_build_status_proto_rawDescOnce sync.Once
	file_google_devtools_build_v1_build_status_proto_rawDescData = file_google_devtools_build_v1_build_status_proto_rawDesc
)

func file_google_devtools_build_v1_build_status_proto_rawDescGZIP() []byte {
	file_google_devtools_build_v1_build_status_proto_rawDescOnce.Do(func() {
		file_google_devtools_build_v1_build_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_build_v1_build_status_proto_rawDescData)
	})
	return file_google_devtools_build_v1_build_status_proto_rawDescData
}

var file_google_devtools_build_v1_build_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_build_v1_build_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_devtools_build_v1_build_status_proto_goTypes = []interface{}{
	(BuildStatus_Result)(0),     // 0: google.devtools.build.v1.BuildStatus.Result
	(*BuildStatus)(nil),         // 1: google.devtools.build.v1.BuildStatus
	(*wrappers.Int32Value)(nil), // 2: google.protobuf.Int32Value
	(*any.Any)(nil),             // 3: google.protobuf.Any
}
var file_google_devtools_build_v1_build_status_proto_depIdxs = []int32{
	0, // 0: google.devtools.build.v1.BuildStatus.result:type_name -> google.devtools.build.v1.BuildStatus.Result
	2, // 1: google.devtools.build.v1.BuildStatus.build_tool_exit_code:type_name -> google.protobuf.Int32Value
	3, // 2: google.devtools.build.v1.BuildStatus.details:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_devtools_build_v1_build_status_proto_init() }
func file_google_devtools_build_v1_build_status_proto_init() {
	if File_google_devtools_build_v1_build_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_build_v1_build_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_build_v1_build_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_build_v1_build_status_proto_goTypes,
		DependencyIndexes: file_google_devtools_build_v1_build_status_proto_depIdxs,
		EnumInfos:         file_google_devtools_build_v1_build_status_proto_enumTypes,
		MessageInfos:      file_google_devtools_build_v1_build_status_proto_msgTypes,
	}.Build()
	File_google_devtools_build_v1_build_status_proto = out.File
	file_google_devtools_build_v1_build_status_proto_rawDesc = nil
	file_google_devtools_build_v1_build_status_proto_goTypes = nil
	file_google_devtools_build_v1_build_status_proto_depIdxs = nil
}
