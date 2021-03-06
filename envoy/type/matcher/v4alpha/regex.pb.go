// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v4alpha/regex.proto

package envoy_type_matcher_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
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

type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType           isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	Regex                string                    `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b6039e2e80d4e39, []int{0}
}

func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher.Unmarshal(m, b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher.Size(m)
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

type RegexMatcher_GoogleRE2 struct {
	HiddenEnvoyDeprecatedMaxProgramSize *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=hidden_envoy_deprecated_max_program_size,json=hiddenEnvoyDeprecatedMaxProgramSize,proto3" json:"hidden_envoy_deprecated_max_program_size,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral                struct{}              `json:"-"`
	XXX_unrecognized                    []byte                `json:"-"`
	XXX_sizecache                       int32                 `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b6039e2e80d4e39, []int{0, 0}
}

func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Unmarshal(m, b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Size(m)
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *RegexMatcher_GoogleRE2) GetHiddenEnvoyDeprecatedMaxProgramSize() *wrappers.UInt32Value {
	if m != nil {
		return m.HiddenEnvoyDeprecatedMaxProgramSize
	}
	return nil
}

type RegexMatchAndSubstitute struct {
	Pattern              *RegexMatcher `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Substitution         string        `protobuf:"bytes,2,opt,name=substitution,proto3" json:"substitution,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RegexMatchAndSubstitute) Reset()         { *m = RegexMatchAndSubstitute{} }
func (m *RegexMatchAndSubstitute) String() string { return proto.CompactTextString(m) }
func (*RegexMatchAndSubstitute) ProtoMessage()    {}
func (*RegexMatchAndSubstitute) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b6039e2e80d4e39, []int{1}
}

func (m *RegexMatchAndSubstitute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatchAndSubstitute.Unmarshal(m, b)
}
func (m *RegexMatchAndSubstitute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatchAndSubstitute.Marshal(b, m, deterministic)
}
func (m *RegexMatchAndSubstitute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatchAndSubstitute.Merge(m, src)
}
func (m *RegexMatchAndSubstitute) XXX_Size() int {
	return xxx_messageInfo_RegexMatchAndSubstitute.Size(m)
}
func (m *RegexMatchAndSubstitute) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatchAndSubstitute.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatchAndSubstitute proto.InternalMessageInfo

func (m *RegexMatchAndSubstitute) GetPattern() *RegexMatcher {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *RegexMatchAndSubstitute) GetSubstitution() string {
	if m != nil {
		return m.Substitution
	}
	return ""
}

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "envoy.type.matcher.v4alpha.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "envoy.type.matcher.v4alpha.RegexMatcher.GoogleRE2")
	proto.RegisterType((*RegexMatchAndSubstitute)(nil), "envoy.type.matcher.v4alpha.RegexMatchAndSubstitute")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v4alpha/regex.proto", fileDescriptor_6b6039e2e80d4e39)
}

var fileDescriptor_6b6039e2e80d4e39 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x71, 0xca, 0x36, 0xea, 0xed, 0x30, 0xe5, 0xb2, 0xaa, 0x62, 0x53, 0x29, 0x08, 0x05,
	0x09, 0x6c, 0x29, 0xdd, 0xa9, 0x27, 0xb0, 0x98, 0x80, 0xc3, 0xa4, 0xca, 0x13, 0xdc, 0x50, 0xe4,
	0x2e, 0x7f, 0x52, 0x4b, 0xad, 0x6d, 0x39, 0x4e, 0x97, 0xee, 0xc4, 0x0d, 0xc4, 0x47, 0xe0, 0x8b,
	0x20, 0xb8, 0x23, 0x71, 0xe5, 0xdb, 0xa0, 0x9d, 0x50, 0xe2, 0x64, 0x30, 0xc1, 0x28, 0x37, 0xcb,
	0xff, 0xf7, 0x9e, 0x7f, 0xcf, 0x7f, 0x7c, 0x1f, 0xd4, 0x52, 0xaf, 0xa8, 0x5b, 0x19, 0xa0, 0x0b,
	0xe1, 0x4e, 0x67, 0x60, 0xe9, 0xf2, 0x50, 0xcc, 0xcd, 0x4c, 0x50, 0x0b, 0x19, 0x94, 0xc4, 0x58,
	0xed, 0x74, 0xd8, 0xaf, 0x75, 0xa4, 0xd2, 0x91, 0x46, 0x47, 0x1a, 0x5d, 0xff, 0x20, 0xd3, 0x3a,
	0x9b, 0x03, 0xad, 0x95, 0xd3, 0xe2, 0x0d, 0x3d, 0xb3, 0xc2, 0x18, 0xb0, 0xb9, 0xf7, 0xf6, 0xf7,
	0x8b, 0xd4, 0x08, 0x2a, 0x94, 0xd2, 0x4e, 0x38, 0xa9, 0x55, 0x4e, 0x73, 0x27, 0x5c, 0xd1, 0x8e,
	0xef, 0xfc, 0x31, 0x5e, 0x82, 0xcd, 0xa5, 0x56, 0x52, 0x65, 0x8d, 0x64, 0x6f, 0x29, 0xe6, 0x32,
	0x15, 0x0e, 0x68, 0x7b, 0xf0, 0x83, 0xe1, 0xbb, 0x0e, 0xde, 0xe1, 0x15, 0xe6, 0xb1, 0x67, 0x0a,
	0x5f, 0x63, 0xec, 0x69, 0x12, 0x0b, 0x71, 0x0f, 0x0d, 0x50, 0xb4, 0x1d, 0xc7, 0xe4, 0x7a, 0x78,
	0xf2, 0xbb, 0x9b, 0x3c, 0xab, 0xad, 0xfc, 0x28, 0x66, 0xb7, 0x2e, 0xd8, 0xc6, 0x07, 0x14, 0xec,
	0xa2, 0xe7, 0x37, 0x78, 0xd7, 0x27, 0x72, 0x88, 0xc3, 0x7d, 0xbc, 0x51, 0xff, 0x4a, 0x2f, 0x18,
	0xa0, 0xa8, 0xcb, 0xb6, 0x2e, 0xd8, 0x4d, 0x1b, 0x0c, 0x10, 0xf7, 0xb7, 0xfd, 0xcf, 0x08, 0x77,
	0x2f, 0x33, 0xc2, 0x33, 0x1c, 0xcd, 0x64, 0x9a, 0x82, 0x4a, 0xea, 0xf7, 0x93, 0x14, 0x8c, 0x85,
	0x53, 0xe1, 0x20, 0x4d, 0x16, 0xa2, 0x4c, 0x8c, 0xd5, 0x99, 0x15, 0x8b, 0x24, 0x97, 0xe7, 0xd0,
	0x90, 0xde, 0x26, 0xfe, 0x29, 0xd2, 0x7e, 0x25, 0x79, 0xf9, 0x42, 0xb9, 0x51, 0xfc, 0x4a, 0xcc,
	0x0b, 0x60, 0x41, 0x0f, 0xf1, 0xbb, 0x3e, 0xf1, 0xa8, 0x0a, 0x7c, 0x7a, 0x99, 0x77, 0x2c, 0xca,
	0x89, 0x4f, 0x3b, 0x91, 0xe7, 0x30, 0x1e, 0x7d, 0xfc, 0xfa, 0xfe, 0x80, 0xe0, 0x87, 0x7f, 0xab,
	0x3d, 0xba, 0xa6, 0xf1, 0xf8, 0x41, 0x65, 0xba, 0x87, 0x87, 0xeb, 0x4d, 0x2c, 0xc4, 0xdb, 0xa0,
	0x32, 0xa9, 0x20, 0xa9, 0x64, 0x61, 0xe7, 0x07, 0x43, 0xc3, 0x4f, 0x08, 0xef, 0xfd, 0x12, 0x3d,
	0x51, 0xe9, 0x49, 0x31, 0xcd, 0x9d, 0x74, 0x85, 0x83, 0x90, 0xe1, 0x2d, 0x23, 0x9c, 0x03, 0xab,
	0x9a, 0x9e, 0xd1, 0xff, 0x6e, 0x84, 0xb7, 0xc6, 0x70, 0x88, 0x77, 0xf2, 0x36, 0x51, 0x6a, 0xe5,
	0x17, 0xc0, 0xaf, 0xdc, 0x8d, 0x0f, 0xab, 0x0a, 0x14, 0x3f, 0x5a, 0x57, 0xe1, 0x0a, 0x1d, 0x7b,
	0xfc, 0xe5, 0xed, 0xb7, 0xef, 0x9b, 0xc1, 0x6e, 0x07, 0x47, 0x52, 0x7b, 0x30, 0x63, 0x75, 0xb9,
	0xfa, 0x07, 0x23, 0xc3, 0x75, 0xd8, 0xa4, 0xda, 0xd2, 0x04, 0x4d, 0x37, 0xeb, 0x75, 0x8d, 0x7e,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x1c, 0x2f, 0x2f, 0x4d, 0x03, 0x00, 0x00,
}
