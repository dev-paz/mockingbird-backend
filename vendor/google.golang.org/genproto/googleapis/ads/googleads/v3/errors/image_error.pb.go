// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/image_error.proto

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

// Enum describing possible image errors.
type ImageErrorEnum_ImageError int32

const (
	// Enum unspecified.
	ImageErrorEnum_UNSPECIFIED ImageErrorEnum_ImageError = 0
	// The received error code is not known in this version.
	ImageErrorEnum_UNKNOWN ImageErrorEnum_ImageError = 1
	// The image is not valid.
	ImageErrorEnum_INVALID_IMAGE ImageErrorEnum_ImageError = 2
	// The image could not be stored.
	ImageErrorEnum_STORAGE_ERROR ImageErrorEnum_ImageError = 3
	// There was a problem with the request.
	ImageErrorEnum_BAD_REQUEST ImageErrorEnum_ImageError = 4
	// The image is not of legal dimensions.
	ImageErrorEnum_UNEXPECTED_SIZE ImageErrorEnum_ImageError = 5
	// Animated image are not permitted.
	ImageErrorEnum_ANIMATED_NOT_ALLOWED ImageErrorEnum_ImageError = 6
	// Animation is too long.
	ImageErrorEnum_ANIMATION_TOO_LONG ImageErrorEnum_ImageError = 7
	// There was an error on the server.
	ImageErrorEnum_SERVER_ERROR ImageErrorEnum_ImageError = 8
	// Image cannot be in CMYK color format.
	ImageErrorEnum_CMYK_JPEG_NOT_ALLOWED ImageErrorEnum_ImageError = 9
	// Flash images are not permitted.
	ImageErrorEnum_FLASH_NOT_ALLOWED ImageErrorEnum_ImageError = 10
	// Flash images must support clickTag.
	ImageErrorEnum_FLASH_WITHOUT_CLICKTAG ImageErrorEnum_ImageError = 11
	// A flash error has occurred after fixing the click tag.
	ImageErrorEnum_FLASH_ERROR_AFTER_FIXING_CLICK_TAG ImageErrorEnum_ImageError = 12
	// Unacceptable visual effects.
	ImageErrorEnum_ANIMATED_VISUAL_EFFECT ImageErrorEnum_ImageError = 13
	// There was a problem with the flash image.
	ImageErrorEnum_FLASH_ERROR ImageErrorEnum_ImageError = 14
	// Incorrect image layout.
	ImageErrorEnum_LAYOUT_PROBLEM ImageErrorEnum_ImageError = 15
	// There was a problem reading the image file.
	ImageErrorEnum_PROBLEM_READING_IMAGE_FILE ImageErrorEnum_ImageError = 16
	// There was an error storing the image.
	ImageErrorEnum_ERROR_STORING_IMAGE ImageErrorEnum_ImageError = 17
	// The aspect ratio of the image is not allowed.
	ImageErrorEnum_ASPECT_RATIO_NOT_ALLOWED ImageErrorEnum_ImageError = 18
	// Flash cannot have network objects.
	ImageErrorEnum_FLASH_HAS_NETWORK_OBJECTS ImageErrorEnum_ImageError = 19
	// Flash cannot have network methods.
	ImageErrorEnum_FLASH_HAS_NETWORK_METHODS ImageErrorEnum_ImageError = 20
	// Flash cannot have a Url.
	ImageErrorEnum_FLASH_HAS_URL ImageErrorEnum_ImageError = 21
	// Flash cannot use mouse tracking.
	ImageErrorEnum_FLASH_HAS_MOUSE_TRACKING ImageErrorEnum_ImageError = 22
	// Flash cannot have a random number.
	ImageErrorEnum_FLASH_HAS_RANDOM_NUM ImageErrorEnum_ImageError = 23
	// Ad click target cannot be '_self'.
	ImageErrorEnum_FLASH_SELF_TARGETS ImageErrorEnum_ImageError = 24
	// GetUrl method should only use '_blank'.
	ImageErrorEnum_FLASH_BAD_GETURL_TARGET ImageErrorEnum_ImageError = 25
	// Flash version is not supported.
	ImageErrorEnum_FLASH_VERSION_NOT_SUPPORTED ImageErrorEnum_ImageError = 26
	// Flash movies need to have hard coded click URL or clickTAG
	ImageErrorEnum_FLASH_WITHOUT_HARD_CODED_CLICK_URL ImageErrorEnum_ImageError = 27
	// Uploaded flash file is corrupted.
	ImageErrorEnum_INVALID_FLASH_FILE ImageErrorEnum_ImageError = 28
	// Uploaded flash file can be parsed, but the click tag can not be fixed
	// properly.
	ImageErrorEnum_FAILED_TO_FIX_CLICK_TAG_IN_FLASH ImageErrorEnum_ImageError = 29
	// Flash movie accesses network resources
	ImageErrorEnum_FLASH_ACCESSES_NETWORK_RESOURCES ImageErrorEnum_ImageError = 30
	// Flash movie attempts to call external javascript code
	ImageErrorEnum_FLASH_EXTERNAL_JS_CALL ImageErrorEnum_ImageError = 31
	// Flash movie attempts to call flash system commands
	ImageErrorEnum_FLASH_EXTERNAL_FS_CALL ImageErrorEnum_ImageError = 32
	// Image file is too large.
	ImageErrorEnum_FILE_TOO_LARGE ImageErrorEnum_ImageError = 33
	// Image data is too large.
	ImageErrorEnum_IMAGE_DATA_TOO_LARGE ImageErrorEnum_ImageError = 34
	// Error while processing the image.
	ImageErrorEnum_IMAGE_PROCESSING_ERROR ImageErrorEnum_ImageError = 35
	// Image is too small.
	ImageErrorEnum_IMAGE_TOO_SMALL ImageErrorEnum_ImageError = 36
	// Input was invalid.
	ImageErrorEnum_INVALID_INPUT ImageErrorEnum_ImageError = 37
	// There was a problem reading the image file.
	ImageErrorEnum_PROBLEM_READING_FILE ImageErrorEnum_ImageError = 38
	// Image constraints are violated, but details like ASPECT_RATIO_NOT_ALLOWED
	// can't be provided. This happens when asset spec contains more than one
	// constraint and different criteria of different constraints are violated.
	ImageErrorEnum_IMAGE_CONSTRAINTS_VIOLATED ImageErrorEnum_ImageError = 39
)

var ImageErrorEnum_ImageError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INVALID_IMAGE",
	3:  "STORAGE_ERROR",
	4:  "BAD_REQUEST",
	5:  "UNEXPECTED_SIZE",
	6:  "ANIMATED_NOT_ALLOWED",
	7:  "ANIMATION_TOO_LONG",
	8:  "SERVER_ERROR",
	9:  "CMYK_JPEG_NOT_ALLOWED",
	10: "FLASH_NOT_ALLOWED",
	11: "FLASH_WITHOUT_CLICKTAG",
	12: "FLASH_ERROR_AFTER_FIXING_CLICK_TAG",
	13: "ANIMATED_VISUAL_EFFECT",
	14: "FLASH_ERROR",
	15: "LAYOUT_PROBLEM",
	16: "PROBLEM_READING_IMAGE_FILE",
	17: "ERROR_STORING_IMAGE",
	18: "ASPECT_RATIO_NOT_ALLOWED",
	19: "FLASH_HAS_NETWORK_OBJECTS",
	20: "FLASH_HAS_NETWORK_METHODS",
	21: "FLASH_HAS_URL",
	22: "FLASH_HAS_MOUSE_TRACKING",
	23: "FLASH_HAS_RANDOM_NUM",
	24: "FLASH_SELF_TARGETS",
	25: "FLASH_BAD_GETURL_TARGET",
	26: "FLASH_VERSION_NOT_SUPPORTED",
	27: "FLASH_WITHOUT_HARD_CODED_CLICK_URL",
	28: "INVALID_FLASH_FILE",
	29: "FAILED_TO_FIX_CLICK_TAG_IN_FLASH",
	30: "FLASH_ACCESSES_NETWORK_RESOURCES",
	31: "FLASH_EXTERNAL_JS_CALL",
	32: "FLASH_EXTERNAL_FS_CALL",
	33: "FILE_TOO_LARGE",
	34: "IMAGE_DATA_TOO_LARGE",
	35: "IMAGE_PROCESSING_ERROR",
	36: "IMAGE_TOO_SMALL",
	37: "INVALID_INPUT",
	38: "PROBLEM_READING_FILE",
	39: "IMAGE_CONSTRAINTS_VIOLATED",
}

var ImageErrorEnum_ImageError_value = map[string]int32{
	"UNSPECIFIED":                        0,
	"UNKNOWN":                            1,
	"INVALID_IMAGE":                      2,
	"STORAGE_ERROR":                      3,
	"BAD_REQUEST":                        4,
	"UNEXPECTED_SIZE":                    5,
	"ANIMATED_NOT_ALLOWED":               6,
	"ANIMATION_TOO_LONG":                 7,
	"SERVER_ERROR":                       8,
	"CMYK_JPEG_NOT_ALLOWED":              9,
	"FLASH_NOT_ALLOWED":                  10,
	"FLASH_WITHOUT_CLICKTAG":             11,
	"FLASH_ERROR_AFTER_FIXING_CLICK_TAG": 12,
	"ANIMATED_VISUAL_EFFECT":             13,
	"FLASH_ERROR":                        14,
	"LAYOUT_PROBLEM":                     15,
	"PROBLEM_READING_IMAGE_FILE":         16,
	"ERROR_STORING_IMAGE":                17,
	"ASPECT_RATIO_NOT_ALLOWED":           18,
	"FLASH_HAS_NETWORK_OBJECTS":          19,
	"FLASH_HAS_NETWORK_METHODS":          20,
	"FLASH_HAS_URL":                      21,
	"FLASH_HAS_MOUSE_TRACKING":           22,
	"FLASH_HAS_RANDOM_NUM":               23,
	"FLASH_SELF_TARGETS":                 24,
	"FLASH_BAD_GETURL_TARGET":            25,
	"FLASH_VERSION_NOT_SUPPORTED":        26,
	"FLASH_WITHOUT_HARD_CODED_CLICK_URL": 27,
	"INVALID_FLASH_FILE":                 28,
	"FAILED_TO_FIX_CLICK_TAG_IN_FLASH":   29,
	"FLASH_ACCESSES_NETWORK_RESOURCES":   30,
	"FLASH_EXTERNAL_JS_CALL":             31,
	"FLASH_EXTERNAL_FS_CALL":             32,
	"FILE_TOO_LARGE":                     33,
	"IMAGE_DATA_TOO_LARGE":               34,
	"IMAGE_PROCESSING_ERROR":             35,
	"IMAGE_TOO_SMALL":                    36,
	"INVALID_INPUT":                      37,
	"PROBLEM_READING_FILE":               38,
	"IMAGE_CONSTRAINTS_VIOLATED":         39,
}

func (x ImageErrorEnum_ImageError) String() string {
	return proto.EnumName(ImageErrorEnum_ImageError_name, int32(x))
}

func (ImageErrorEnum_ImageError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_276eefc416298097, []int{0, 0}
}

// Container for enum describing possible image errors.
type ImageErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageErrorEnum) Reset()         { *m = ImageErrorEnum{} }
func (m *ImageErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ImageErrorEnum) ProtoMessage()    {}
func (*ImageErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_276eefc416298097, []int{0}
}

func (m *ImageErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageErrorEnum.Unmarshal(m, b)
}
func (m *ImageErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageErrorEnum.Marshal(b, m, deterministic)
}
func (m *ImageErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageErrorEnum.Merge(m, src)
}
func (m *ImageErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ImageErrorEnum.Size(m)
}
func (m *ImageErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ImageErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.ImageErrorEnum_ImageError", ImageErrorEnum_ImageError_name, ImageErrorEnum_ImageError_value)
	proto.RegisterType((*ImageErrorEnum)(nil), "google.ads.googleads.v3.errors.ImageErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/image_error.proto", fileDescriptor_276eefc416298097)
}

var fileDescriptor_276eefc416298097 = []byte{
	// 801 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xae, 0x9d, 0x36, 0x71, 0xe9, 0xc4, 0xa6, 0xe9, 0xf8, 0xdf, 0x71, 0x52, 0x35, 0x4d, 0x6f,
	0x52, 0x01, 0xdd, 0xd4, 0x13, 0xb5, 0x3b, 0xbb, 0xa2, 0xc5, 0x25, 0xb7, 0x24, 0x57, 0x76, 0x02,
	0x03, 0x84, 0x5a, 0x19, 0x82, 0x81, 0x58, 0x6b, 0x68, 0xdd, 0xbc, 0x4a, 0xef, 0x3d, 0xe6, 0x51,
	0xfa, 0x28, 0x79, 0x8a, 0x62, 0x96, 0x92, 0x36, 0x2e, 0xd2, 0x9c, 0x44, 0xcd, 0xf7, 0xf1, 0x9b,
	0xe1, 0x7c, 0xb3, 0x43, 0x7e, 0x99, 0x96, 0xe5, 0xf4, 0xfd, 0x75, 0x67, 0x3c, 0xa9, 0x3a, 0xe1,
	0x88, 0xa7, 0x0f, 0xdd, 0xce, 0xf5, 0x7c, 0x5e, 0xce, 0xab, 0xce, 0xcd, 0xed, 0x78, 0x7a, 0xed,
	0xeb, 0x3f, 0xed, 0xbb, 0x79, 0x79, 0x5f, 0xb2, 0xb3, 0x40, 0x6b, 0x8f, 0x27, 0x55, 0x7b, 0x75,
	0xa3, 0xfd, 0xa1, 0xdb, 0x0e, 0x37, 0x8e, 0x4f, 0x97, 0x8a, 0x77, 0x37, 0x9d, 0xf1, 0x6c, 0x56,
	0xde, 0x8f, 0xef, 0x6f, 0xca, 0x59, 0x15, 0x6e, 0xb7, 0x3e, 0x6e, 0x90, 0x2d, 0x81, 0x9a, 0x80,
	0x6c, 0x98, 0xfd, 0x79, 0xdb, 0xfa, 0x6b, 0x83, 0x90, 0x26, 0xc4, 0xb6, 0xc9, 0x66, 0xa1, 0x6c,
	0x0e, 0x91, 0x48, 0x04, 0xc4, 0xf4, 0x1b, 0xb6, 0x49, 0x9e, 0x14, 0x6a, 0xa8, 0xf4, 0x85, 0xa2,
	0x6b, 0x6c, 0x87, 0x3c, 0x13, 0x6a, 0xc4, 0xa5, 0x88, 0xbd, 0xc8, 0x78, 0x0a, 0x74, 0x1d, 0x43,
	0xd6, 0x69, 0xc3, 0x53, 0xf0, 0x60, 0x8c, 0x36, 0xf4, 0x11, 0x6a, 0xf4, 0x79, 0xec, 0x0d, 0xfc,
	0x56, 0x80, 0x75, 0xf4, 0x5b, 0xb6, 0x4b, 0xb6, 0x0b, 0x05, 0x97, 0x39, 0x44, 0x0e, 0x62, 0x6f,
	0xc5, 0x3b, 0xa0, 0xdf, 0xb1, 0x43, 0xf2, 0x9c, 0x2b, 0x91, 0x71, 0x0c, 0x29, 0xed, 0x3c, 0x97,
	0x52, 0x5f, 0x40, 0x4c, 0x1f, 0xb3, 0x7d, 0xc2, 0x02, 0x22, 0xb4, 0xf2, 0x4e, 0x6b, 0x2f, 0xb5,
	0x4a, 0xe9, 0x13, 0x46, 0xc9, 0x53, 0x0b, 0x66, 0x04, 0x66, 0x91, 0x69, 0x83, 0x1d, 0x91, 0xbd,
	0x28, 0x7b, 0x3b, 0xf4, 0xe7, 0x39, 0xa4, 0x0f, 0x44, 0xbe, 0x67, 0x7b, 0x64, 0x27, 0x91, 0xdc,
	0x0e, 0x1e, 0x84, 0x09, 0x3b, 0x26, 0xfb, 0x21, 0x7c, 0x21, 0xdc, 0x40, 0x17, 0xce, 0x47, 0x52,
	0x44, 0x43, 0xc7, 0x53, 0xba, 0xc9, 0xde, 0x90, 0x56, 0xc0, 0x6a, 0x79, 0xcf, 0x13, 0x07, 0xc6,
	0x27, 0xe2, 0x52, 0xa8, 0x34, 0xd0, 0x3c, 0xf2, 0x9e, 0xa2, 0xc6, 0xaa, 0xf2, 0x91, 0xb0, 0x05,
	0x97, 0x1e, 0x92, 0x04, 0x22, 0x47, 0x9f, 0xe1, 0xdb, 0x3f, 0xd3, 0xa0, 0x5b, 0x8c, 0x91, 0x2d,
	0xc9, 0xdf, 0x62, 0xa6, 0xdc, 0xe8, 0xbe, 0x84, 0x8c, 0x6e, 0xb3, 0x33, 0x72, 0xbc, 0xf8, 0xe3,
	0x0d, 0xf0, 0x18, 0xf5, 0xeb, 0x76, 0xfa, 0x44, 0x48, 0xa0, 0x94, 0x1d, 0x90, 0xdd, 0x50, 0x02,
	0x76, 0x76, 0x85, 0xd2, 0x1d, 0x76, 0x4a, 0x0e, 0x39, 0x9a, 0xe3, 0xbc, 0xc1, 0xee, 0x3c, 0x78,
	0x1b, 0x63, 0x2f, 0xc8, 0x51, 0xc8, 0x3d, 0xe0, 0xd6, 0x2b, 0x70, 0x17, 0xda, 0x0c, 0xbd, 0xee,
	0x9f, 0x43, 0xe4, 0x2c, 0xdd, 0xfd, 0x32, 0x9c, 0x81, 0x1b, 0xe8, 0xd8, 0xd2, 0xe7, 0x68, 0x64,
	0x03, 0x17, 0x46, 0xd2, 0x3d, 0x4c, 0xd7, 0x84, 0x32, 0x5d, 0x58, 0xf0, 0xce, 0xf0, 0x68, 0x28,
	0x54, 0x4a, 0xf7, 0xd1, 0xc0, 0x06, 0x35, 0x5c, 0xc5, 0x3a, 0xf3, 0xaa, 0xc8, 0xe8, 0x01, 0x1a,
	0x18, 0x10, 0x0b, 0x32, 0xf1, 0x8e, 0x9b, 0x14, 0x9c, 0xa5, 0x87, 0xec, 0x84, 0x1c, 0x84, 0x38,
	0x8e, 0x47, 0x0a, 0xae, 0x30, 0x72, 0x81, 0xd2, 0x23, 0xf6, 0x92, 0x9c, 0x04, 0x70, 0x04, 0xc6,
	0xa2, 0xf3, 0xf8, 0x38, 0x5b, 0xe4, 0xb9, 0x36, 0x0e, 0x62, 0x7a, 0xdc, 0xd8, 0xb3, 0xb4, 0x6e,
	0xc0, 0x4d, 0xec, 0x23, 0x1d, 0x43, 0xbc, 0xb0, 0x07, 0xab, 0x3e, 0xc1, 0xec, 0xcb, 0x21, 0x0d,
	0xfc, 0xba, 0xab, 0xa7, 0xec, 0x35, 0x79, 0x95, 0x70, 0x21, 0x21, 0xf6, 0x4e, 0xa3, 0xad, 0x8d,
	0xa7, 0x5e, 0xa8, 0x40, 0xa5, 0x2f, 0x6a, 0x56, 0x7d, 0x8b, 0x47, 0x11, 0x58, 0x0b, 0x4d, 0xab,
	0x0c, 0x58, 0x5d, 0x98, 0x08, 0x2c, 0x3d, 0x6b, 0xc6, 0x08, 0x2e, 0x1d, 0x18, 0xc5, 0xa5, 0x3f,
	0xb7, 0x3e, 0xe2, 0x52, 0xd2, 0x97, 0x5f, 0xc0, 0x92, 0x05, 0xf6, 0x0a, 0xa7, 0x01, 0xab, 0x09,
	0x53, 0x8d, 0x2f, 0xa7, 0x3f, 0x60, 0x1f, 0x83, 0xfb, 0x31, 0x77, 0xfc, 0x33, 0xa4, 0x85, 0x4a,
	0x01, 0xc9, 0x8d, 0xc6, 0x62, 0x70, 0x14, 0xc2, 0x5c, 0xfd, 0x88, 0xdf, 0x54, 0xc0, 0xf0, 0x82,
	0xcd, 0x50, 0xfe, 0xf5, 0x83, 0xef, 0x53, 0xe5, 0x85, 0xa3, 0x3f, 0xa1, 0xfa, 0x7f, 0x67, 0xad,
	0xee, 0xc7, 0x1b, 0x9c, 0xc2, 0xa0, 0x10, 0x69, 0x65, 0x9d, 0xe1, 0x42, 0x39, 0xeb, 0x47, 0x42,
	0x4b, 0x1c, 0x6c, 0xfa, 0x73, 0xff, 0xd3, 0x1a, 0x69, 0xfd, 0x51, 0xde, 0xb6, 0xbf, 0xbe, 0x71,
	0xfa, 0xdb, 0xcd, 0xf6, 0xc8, 0x71, 0xc9, 0xe4, 0x6b, 0xef, 0xe2, 0xc5, 0x95, 0x69, 0xf9, 0x7e,
	0x3c, 0x9b, 0xb6, 0xcb, 0xf9, 0xb4, 0x33, 0xbd, 0x9e, 0xd5, 0x2b, 0x68, 0xb9, 0xe6, 0xee, 0x6e,
	0xaa, 0xff, 0xdb, 0x7a, 0xbf, 0x86, 0x9f, 0xbf, 0xd7, 0x1f, 0xa5, 0x9c, 0x7f, 0x5c, 0x3f, 0x4b,
	0x83, 0x18, 0x9f, 0x54, 0xed, 0x70, 0xc4, 0xd3, 0xa8, 0xdb, 0xae, 0x53, 0x56, 0xff, 0x2c, 0x09,
	0x57, 0x7c, 0x52, 0x5d, 0xad, 0x08, 0x57, 0xa3, 0xee, 0x55, 0x20, 0x7c, 0x5a, 0x6f, 0x85, 0x68,
	0xaf, 0xc7, 0x27, 0x55, 0xaf, 0xb7, 0xa2, 0xf4, 0x7a, 0xa3, 0x6e, 0xaf, 0x17, 0x48, 0xbf, 0x3f,
	0xae, 0xab, 0xeb, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xeb, 0x84, 0xf4, 0x92, 0x05, 0x00,
	0x00,
}
