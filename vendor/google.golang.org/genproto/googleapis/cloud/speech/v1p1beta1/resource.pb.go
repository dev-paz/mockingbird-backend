// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/speech/v1p1beta1/resource.proto

package speech

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

// A set of words or phrases that represents a common concept likely to appear
// in your audio, for example a list of passenger ship names. CustomClass items
// can be substituted into placeholders that you set in PhraseSet phrases.
type CustomClass struct {
	// The resource name of the custom class.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If this custom class is a resource, the custom_class_id is the resource id
	// of the CustomClass.
	CustomClassId string `protobuf:"bytes,2,opt,name=custom_class_id,json=customClassId,proto3" json:"custom_class_id,omitempty"`
	// A collection of class items.
	Items                []*CustomClass_ClassItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CustomClass) Reset()         { *m = CustomClass{} }
func (m *CustomClass) String() string { return proto.CompactTextString(m) }
func (*CustomClass) ProtoMessage()    {}
func (*CustomClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_0411f669716aa30c, []int{0}
}

func (m *CustomClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomClass.Unmarshal(m, b)
}
func (m *CustomClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomClass.Marshal(b, m, deterministic)
}
func (m *CustomClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomClass.Merge(m, src)
}
func (m *CustomClass) XXX_Size() int {
	return xxx_messageInfo_CustomClass.Size(m)
}
func (m *CustomClass) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomClass.DiscardUnknown(m)
}

var xxx_messageInfo_CustomClass proto.InternalMessageInfo

func (m *CustomClass) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomClass) GetCustomClassId() string {
	if m != nil {
		return m.CustomClassId
	}
	return ""
}

func (m *CustomClass) GetItems() []*CustomClass_ClassItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// An item of the class.
type CustomClass_ClassItem struct {
	// The class item's value.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomClass_ClassItem) Reset()         { *m = CustomClass_ClassItem{} }
func (m *CustomClass_ClassItem) String() string { return proto.CompactTextString(m) }
func (*CustomClass_ClassItem) ProtoMessage()    {}
func (*CustomClass_ClassItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0411f669716aa30c, []int{0, 0}
}

func (m *CustomClass_ClassItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomClass_ClassItem.Unmarshal(m, b)
}
func (m *CustomClass_ClassItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomClass_ClassItem.Marshal(b, m, deterministic)
}
func (m *CustomClass_ClassItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomClass_ClassItem.Merge(m, src)
}
func (m *CustomClass_ClassItem) XXX_Size() int {
	return xxx_messageInfo_CustomClass_ClassItem.Size(m)
}
func (m *CustomClass_ClassItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomClass_ClassItem.DiscardUnknown(m)
}

var xxx_messageInfo_CustomClass_ClassItem proto.InternalMessageInfo

func (m *CustomClass_ClassItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Provides "hints" to the speech recognizer to favor specific words and phrases
// in the results.
type PhraseSet struct {
	// The resource name of the phrase set.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of word and phrases.
	Phrases []*PhraseSet_Phrase `protobuf:"bytes,2,rep,name=phrases,proto3" json:"phrases,omitempty"`
	// Hint Boost. Positive value will increase the probability that a specific
	// phrase will be recognized over other similar sounding phrases. The higher
	// the boost, the higher the chance of false positive recognition as well.
	// Negative boost values would correspond to anti-biasing. Anti-biasing is not
	// enabled, so negative boost will simply be ignored. Though `boost` can
	// accept a wide range of positive values, most use cases are best served with
	// values between 0 (exclusive) and 20. We recommend using a binary search
	// approach to finding the optimal value for your use case. Speech recognition
	// will skip PhraseSets with a boost value of 0.
	Boost                float32  `protobuf:"fixed32,4,opt,name=boost,proto3" json:"boost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhraseSet) Reset()         { *m = PhraseSet{} }
func (m *PhraseSet) String() string { return proto.CompactTextString(m) }
func (*PhraseSet) ProtoMessage()    {}
func (*PhraseSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0411f669716aa30c, []int{1}
}

func (m *PhraseSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhraseSet.Unmarshal(m, b)
}
func (m *PhraseSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhraseSet.Marshal(b, m, deterministic)
}
func (m *PhraseSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhraseSet.Merge(m, src)
}
func (m *PhraseSet) XXX_Size() int {
	return xxx_messageInfo_PhraseSet.Size(m)
}
func (m *PhraseSet) XXX_DiscardUnknown() {
	xxx_messageInfo_PhraseSet.DiscardUnknown(m)
}

var xxx_messageInfo_PhraseSet proto.InternalMessageInfo

func (m *PhraseSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PhraseSet) GetPhrases() []*PhraseSet_Phrase {
	if m != nil {
		return m.Phrases
	}
	return nil
}

func (m *PhraseSet) GetBoost() float32 {
	if m != nil {
		return m.Boost
	}
	return 0
}

// A phrases containing words and phrase "hints" so that
// the speech recognition is more likely to recognize them. This can be used
// to improve the accuracy for specific words and phrases, for example, if
// specific commands are typically spoken by the user. This can also be used
// to add additional words to the vocabulary of the recognizer. See
// [usage limits](https://cloud.google.com/speech-to-text/quotas#content).
//
// List items can also include pre-built or custom classes containing groups
// of words that represent common concepts that occur in natural language. For
// example, rather than providing a phrase hint for every month of the
// year (e.g. "i was born in january", "i was born in febuary", ...), use the
// pre-built `$MONTH` class improves the likelihood of correctly transcribing
// audio that includes months (e.g. "i was born in $month").
// To refer to pre-built classes, use the class' symbol prepended with `$`
// e.g. `$MONTH`. To refer to custom classes that were defined inline in the
// request, set the class's `custom_class_id` to a string unique to all class
// resources and inline classes. Then use the class' id wrapped in $`{...}`
// e.g. "${my-months}". To refer to custom classes resources, use the class'
// id wrapped in `${}` (e.g. `${my-months}`).
type PhraseSet_Phrase struct {
	// The phrase itself.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Hint Boost. Overrides the boost set at the phrase set level.
	// Positive value will increase the probability that a specific phrase will
	// be recognized over other similar sounding phrases. The higher the boost,
	// the higher the chance of false positive recognition as well. Negative
	// boost values would correspond to anti-biasing. Anti-biasing is not
	// enabled, so negative boost will simply be ignored. Though `boost` can
	// accept a wide range of positive values, most use cases are best served
	// with values between 0 and 20. We recommend using a binary search approach
	// to finding the optimal value for your use case. Speech recognition
	// will skip PhraseSets with a boost value of 0.
	Boost                float32  `protobuf:"fixed32,2,opt,name=boost,proto3" json:"boost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhraseSet_Phrase) Reset()         { *m = PhraseSet_Phrase{} }
func (m *PhraseSet_Phrase) String() string { return proto.CompactTextString(m) }
func (*PhraseSet_Phrase) ProtoMessage()    {}
func (*PhraseSet_Phrase) Descriptor() ([]byte, []int) {
	return fileDescriptor_0411f669716aa30c, []int{1, 0}
}

func (m *PhraseSet_Phrase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhraseSet_Phrase.Unmarshal(m, b)
}
func (m *PhraseSet_Phrase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhraseSet_Phrase.Marshal(b, m, deterministic)
}
func (m *PhraseSet_Phrase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhraseSet_Phrase.Merge(m, src)
}
func (m *PhraseSet_Phrase) XXX_Size() int {
	return xxx_messageInfo_PhraseSet_Phrase.Size(m)
}
func (m *PhraseSet_Phrase) XXX_DiscardUnknown() {
	xxx_messageInfo_PhraseSet_Phrase.DiscardUnknown(m)
}

var xxx_messageInfo_PhraseSet_Phrase proto.InternalMessageInfo

func (m *PhraseSet_Phrase) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PhraseSet_Phrase) GetBoost() float32 {
	if m != nil {
		return m.Boost
	}
	return 0
}

// Speech adaptation configuration.
type SpeechAdaptation struct {
	// A collection of phrase sets. To specify the hints inline, leave the
	// phrase set's `name` blank and fill in the rest of its fields. Any
	// phrase set can use any custom class.
	PhraseSets []*PhraseSet `protobuf:"bytes,1,rep,name=phrase_sets,json=phraseSets,proto3" json:"phrase_sets,omitempty"`
	// A collection of custom classes. To specify the classes inline, leave the
	// class' `name` blank and fill in the rest of its fields, giving it a unique
	// `custom_class_id`. Refer to the inline defined class in phrase hints by its
	// `custom_class_id`.
	CustomClasses        []*CustomClass `protobuf:"bytes,2,rep,name=custom_classes,json=customClasses,proto3" json:"custom_classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SpeechAdaptation) Reset()         { *m = SpeechAdaptation{} }
func (m *SpeechAdaptation) String() string { return proto.CompactTextString(m) }
func (*SpeechAdaptation) ProtoMessage()    {}
func (*SpeechAdaptation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0411f669716aa30c, []int{2}
}

func (m *SpeechAdaptation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpeechAdaptation.Unmarshal(m, b)
}
func (m *SpeechAdaptation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpeechAdaptation.Marshal(b, m, deterministic)
}
func (m *SpeechAdaptation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpeechAdaptation.Merge(m, src)
}
func (m *SpeechAdaptation) XXX_Size() int {
	return xxx_messageInfo_SpeechAdaptation.Size(m)
}
func (m *SpeechAdaptation) XXX_DiscardUnknown() {
	xxx_messageInfo_SpeechAdaptation.DiscardUnknown(m)
}

var xxx_messageInfo_SpeechAdaptation proto.InternalMessageInfo

func (m *SpeechAdaptation) GetPhraseSets() []*PhraseSet {
	if m != nil {
		return m.PhraseSets
	}
	return nil
}

func (m *SpeechAdaptation) GetCustomClasses() []*CustomClass {
	if m != nil {
		return m.CustomClasses
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomClass)(nil), "google.cloud.speech.v1p1beta1.CustomClass")
	proto.RegisterType((*CustomClass_ClassItem)(nil), "google.cloud.speech.v1p1beta1.CustomClass.ClassItem")
	proto.RegisterType((*PhraseSet)(nil), "google.cloud.speech.v1p1beta1.PhraseSet")
	proto.RegisterType((*PhraseSet_Phrase)(nil), "google.cloud.speech.v1p1beta1.PhraseSet.Phrase")
	proto.RegisterType((*SpeechAdaptation)(nil), "google.cloud.speech.v1p1beta1.SpeechAdaptation")
}

func init() {
	proto.RegisterFile("google/cloud/speech/v1p1beta1/resource.proto", fileDescriptor_0411f669716aa30c)
}

var fileDescriptor_0411f669716aa30c = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0xbb, 0x43, 0x99, 0x88, 0x0f, 0x19, 0x0a, 0x13, 0x81, 0x48, 0x52, 0xa0, 0x08,
	0xa1, 0x5d, 0x05, 0xae, 0x3a, 0x0a, 0x94, 0x0b, 0x12, 0x0a, 0x55, 0x70, 0x3a, 0x9a, 0x68, 0xb3,
	0x19, 0xf9, 0x8c, 0x6c, 0xcf, 0xca, 0xbb, 0xb9, 0x26, 0x4a, 0xcf, 0x9f, 0xa0, 0xa1, 0xa0, 0xe6,
	0x77, 0xdd, 0x2f, 0xa0, 0x44, 0xd9, 0x75, 0x6c, 0x9f, 0x14, 0xee, 0xd2, 0x8d, 0x77, 0xdf, 0x7b,
	0xb3, 0xef, 0x79, 0x06, 0xde, 0xc6, 0x44, 0x71, 0x8a, 0x5c, 0xa6, 0xb4, 0x5e, 0x71, 0xad, 0x10,
	0xe5, 0x15, 0xbf, 0x1e, 0xa9, 0xd1, 0x12, 0x8d, 0x18, 0xf1, 0x02, 0x35, 0xad, 0x0b, 0x89, 0x4c,
	0x15, 0x64, 0x28, 0x78, 0xe9, 0xd0, 0xcc, 0xa2, 0x99, 0x43, 0xb3, 0x0a, 0xdd, 0x7d, 0x51, 0x8a,
	0x09, 0x95, 0x70, 0x91, 0xe7, 0x64, 0x84, 0x49, 0x28, 0xd7, 0x8e, 0xdc, 0x7d, 0xde, 0xb8, 0xbd,
	0xad, 0x3b, 0xf8, 0xed, 0x43, 0x67, 0xb2, 0xd6, 0x86, 0xb2, 0x49, 0x2a, 0xb4, 0x0e, 0x02, 0x38,
	0xc9, 0x45, 0x86, 0xa1, 0xd7, 0xf3, 0x86, 0xed, 0xc8, 0xd6, 0xc1, 0x6b, 0x78, 0x2c, 0x2d, 0x64,
	0x21, 0x77, 0x98, 0x45, 0xb2, 0x0a, 0x7d, 0x7b, 0xfd, 0x50, 0xd6, 0xcc, 0xe9, 0x2a, 0xf8, 0x02,
	0xa7, 0x89, 0xc1, 0x4c, 0x87, 0xad, 0x5e, 0x6b, 0xd8, 0x79, 0x77, 0xce, 0xee, 0x7c, 0x33, 0x6b,
	0xb4, 0x65, 0x4e, 0xc2, 0x60, 0x16, 0x39, 0x89, 0x6e, 0x1f, 0xda, 0xd5, 0x59, 0xf0, 0x0c, 0x4e,
	0xaf, 0x45, 0xba, 0xde, 0xbf, 0xca, 0x7d, 0x5c, 0xa4, 0x37, 0xe3, 0x04, 0xfa, 0xa5, 0xac, 0x6b,
	0x25, 0x54, 0xa2, 0x99, 0xa4, 0x8c, 0x37, 0x2d, 0x7d, 0x52, 0x05, 0x7d, 0x47, 0x69, 0x34, 0xdf,
	0x94, 0xd5, 0x96, 0xa7, 0x24, 0x5d, 0x42, 0x7c, 0xb3, 0x2f, 0xb7, 0xbc, 0xe1, 0x06, 0x35, 0xdf,
	0x34, 0x3d, 0x6f, 0x07, 0x3f, 0x7d, 0x68, 0xcf, 0xae, 0x0a, 0xa1, 0x71, 0x8e, 0xe6, 0x60, 0x4c,
	0x53, 0x78, 0xa0, 0x2c, 0x40, 0x87, 0xbe, 0x0d, 0x80, 0xdf, 0x13, 0x40, 0x25, 0x57, 0x56, 0xd1,
	0x9e, 0xbf, 0x33, 0xbc, 0x24, 0xd2, 0x26, 0x3c, 0xe9, 0x79, 0x43, 0x3f, 0x72, 0x1f, 0xdd, 0x73,
	0x38, 0x73, 0xc0, 0xc3, 0x81, 0xd4, 0x2c, 0xbf, 0xc1, 0xba, 0xc0, 0x9b, 0xf1, 0x12, 0x5e, 0x1d,
	0x8e, 0xa9, 0x36, 0xf4, 0xf1, 0xd8, 0x90, 0xd4, 0x9e, 0xb2, 0x43, 0xda, 0x7a, 0xa1, 0xd1, 0x6c,
	0x07, 0x7f, 0x3c, 0x78, 0x32, 0xb7, 0x4d, 0xc6, 0x2b, 0xa1, 0xdc, 0xfc, 0x05, 0x53, 0xe8, 0xd4,
	0x18, 0x1d, 0x7a, 0x36, 0x96, 0xe1, 0xb1, 0xb1, 0x44, 0x50, 0x37, 0x0b, 0xbe, 0xc2, 0xa3, 0xe6,
	0x0f, 0xa9, 0x42, 0x7e, 0x73, 0xfc, 0x94, 0xdd, 0x9a, 0x57, 0xd4, 0x97, 0x3f, 0x3c, 0xe8, 0x4b,
	0xca, 0xee, 0x16, 0xb8, 0x7c, 0xea, 0x5c, 0x45, 0xe5, 0xde, 0xcc, 0x76, 0x6b, 0x33, 0xf3, 0xbe,
	0x4d, 0x4a, 0x56, 0x4c, 0xa9, 0xc8, 0x63, 0x46, 0x45, 0xcc, 0x63, 0xcc, 0xed, 0x52, 0xf1, 0x3a,
	0xe5, 0xff, 0x6c, 0xf7, 0x07, 0x77, 0xf0, 0xd7, 0xf3, 0x7e, 0xf9, 0xad, 0xcf, 0x93, 0xf9, 0xf2,
	0xcc, 0x12, 0xdf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x21, 0x55, 0x91, 0x69, 0x15, 0x04, 0x00,
	0x00,
}
