// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v4alpha/tap.proto

package envoy_admin_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/tap/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type TapRequest struct {
	ConfigId             string             `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	TapConfig            *v4alpha.TapConfig `protobuf:"bytes,2,opt,name=tap_config,json=tapConfig,proto3" json:"tap_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TapRequest) Reset()         { *m = TapRequest{} }
func (m *TapRequest) String() string { return proto.CompactTextString(m) }
func (*TapRequest) ProtoMessage()    {}
func (*TapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4783f37990388c8, []int{0}
}

func (m *TapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapRequest.Unmarshal(m, b)
}
func (m *TapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapRequest.Marshal(b, m, deterministic)
}
func (m *TapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapRequest.Merge(m, src)
}
func (m *TapRequest) XXX_Size() int {
	return xxx_messageInfo_TapRequest.Size(m)
}
func (m *TapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TapRequest proto.InternalMessageInfo

func (m *TapRequest) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func (m *TapRequest) GetTapConfig() *v4alpha.TapConfig {
	if m != nil {
		return m.TapConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*TapRequest)(nil), "envoy.admin.v4alpha.TapRequest")
}

func init() { proto.RegisterFile("envoy/admin/v4alpha/tap.proto", fileDescriptor_e4783f37990388c8) }

var fileDescriptor_e4783f37990388c8 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0x33, 0x49, 0xcc, 0x29, 0xc8, 0x48, 0xd4,
	0x2f, 0x49, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06, 0x4b, 0xeb, 0x81, 0xa5,
	0xf5, 0xa0, 0xd2, 0x52, 0xaa, 0x10, 0x3d, 0xc9, 0xf9, 0x79, 0x69, 0x99, 0xe9, 0x20, 0xc5, 0x70,
	0x8d, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x10, 0xbd, 0x52, 0xb2, 0xa5, 0x29, 0x05, 0x89, 0xfa,
	0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xfa, 0xc5, 0x25, 0x89, 0x25,
	0xa5, 0xc5, 0x50, 0x69, 0x45, 0x0c, 0xe9, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0xbc, 0xcc, 0xbc,
	0x74, 0xa8, 0x12, 0xf1, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92, 0x54, 0x7d, 0x18, 0x03, 0x22,
	0xa1, 0xb4, 0x88, 0x91, 0x8b, 0x2b, 0x24, 0xb1, 0x20, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0x48, 0x85, 0x8b, 0x13, 0xe2, 0x98, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27,
	0xf6, 0x5f, 0x4e, 0x2c, 0x45, 0x4c, 0x0a, 0x8c, 0x41, 0x1c, 0x10, 0x19, 0xcf, 0x14, 0x21, 0x1f,
	0x2e, 0xae, 0x92, 0xc4, 0x82, 0x78, 0x08, 0x5f, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x59,
	0x0f, 0xe2, 0x41, 0x88, 0xa0, 0x1e, 0xc8, 0xe3, 0x50, 0xbf, 0xe8, 0x85, 0x24, 0x16, 0x38, 0x83,
	0x45, 0x9d, 0x38, 0x7e, 0x39, 0xb1, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0x71, 0x96, 0xc0, 0x04,
	0xad, 0x14, 0x66, 0x1d, 0xed, 0x90, 0x93, 0xe6, 0x92, 0x44, 0x09, 0x20, 0x63, 0x3d, 0x84, 0xab,
	0x9c, 0xcc, 0x77, 0x35, 0x9c, 0xb8, 0xc8, 0xc6, 0x24, 0xc0, 0xcc, 0xa5, 0x98, 0x99, 0x0f, 0xb1,
	0xa7, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x0f, 0x4b, 0x98, 0x3a, 0x71, 0x84, 0x24, 0x16, 0x04, 0x80,
	0xfc, 0x16, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0xa4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x81, 0x3d,
	0x1b, 0x7c, 0x9c, 0x01, 0x00, 0x00,
}
