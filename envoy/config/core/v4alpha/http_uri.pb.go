// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v4alpha/http_uri.proto

package envoy_config_core_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type HttpUri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType     isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_238f6e98372d5daa, []int{0}
}

func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpUri.Unmarshal(m, b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return xxx_messageInfo_HttpUri.Size(m)
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.config.core.v4alpha.HttpUri")
}

func init() {
	proto.RegisterFile("envoy/config/core/v4alpha/http_uri.proto", fileDescriptor_238f6e98372d5daa)
}

var fileDescriptor_238f6e98372d5daa = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4e, 0x2a, 0x41,
	0x14, 0xc6, 0x19, 0x96, 0xb0, 0xf7, 0xce, 0xbd, 0xc5, 0xcd, 0x36, 0x97, 0x25, 0x4a, 0x50, 0x0a,
	0xa9, 0x66, 0x12, 0xb0, 0xb2, 0x31, 0x4e, 0x2c, 0x28, 0x09, 0x89, 0x35, 0x19, 0x60, 0x58, 0x26,
	0x59, 0xe6, 0x4c, 0x66, 0xcf, 0x6c, 0xa4, 0xb3, 0xf4, 0x19, 0x7c, 0x04, 0x1f, 0xc1, 0xde, 0xc4,
	0xd8, 0xf9, 0x36, 0x86, 0xca, 0xb0, 0x7f, 0x0a, 0xa3, 0x76, 0x27, 0xe7, 0xfb, 0x9d, 0x93, 0xef,
	0xfb, 0xe8, 0x50, 0x99, 0x1c, 0x76, 0x7c, 0x09, 0x66, 0xad, 0x13, 0xbe, 0x04, 0xa7, 0x78, 0x7e,
	0x2e, 0x53, 0xbb, 0x91, 0x7c, 0x83, 0x68, 0xe7, 0xde, 0x69, 0x66, 0x1d, 0x20, 0x44, 0x71, 0x41,
	0xb2, 0x92, 0x64, 0x07, 0x92, 0x55, 0x64, 0xb7, 0x97, 0x00, 0x24, 0xa9, 0xe2, 0x05, 0xb8, 0xf0,
	0x6b, 0xbe, 0xf2, 0x4e, 0xa2, 0x06, 0x53, 0x9e, 0x76, 0x8f, 0xfd, 0xca, 0x4a, 0x2e, 0x8d, 0x01,
	0x2c, 0xd6, 0x19, 0xcf, 0x50, 0xa2, 0xcf, 0x2a, 0xf9, 0xe4, 0x8b, 0x9c, 0x2b, 0x97, 0x69, 0x30,
	0xda, 0x24, 0x15, 0xf2, 0x3f, 0x97, 0xa9, 0x5e, 0x49, 0x54, 0xbc, 0x1e, 0x4a, 0xe1, 0xf4, 0x95,
	0xd0, 0x70, 0x82, 0x68, 0x6f, 0x9c, 0x8e, 0x62, 0x1a, 0x78, 0xa7, 0x3b, 0xa4, 0x4f, 0x86, 0xbf,
	0x45, 0xb8, 0x17, 0x2d, 0xd7, 0xec, 0x93, 0xd9, 0x61, 0x17, 0x0d, 0x68, 0xb8, 0x4c, 0x7d, 0x86,
	0xca, 0x75, 0x9a, 0x9f, 0xe4, 0x49, 0x63, 0x56, 0x2b, 0xd1, 0x25, 0x0d, 0x51, 0x6f, 0x15, 0x78,
	0xec, 0x04, 0x7d, 0x32, 0xfc, 0x33, 0x8a, 0x59, 0x19, 0x8c, 0xd5, 0xc1, 0xd8, 0x75, 0x15, 0x4c,
	0xd0, 0xbd, 0x08, 0x1f, 0x49, 0xeb, 0x17, 0x19, 0x35, 0x66, 0xf5, 0xd5, 0xc5, 0xe0, 0xe1, 0xf9,
	0xbe, 0xd7, 0xa3, 0x47, 0xdf, 0x34, 0x35, 0x66, 0x95, 0x4b, 0x11, 0xd3, 0xa8, 0x6c, 0xd6, 0x66,
	0xe8, 0x94, 0xdc, 0xce, 0x71, 0x67, 0x55, 0x14, 0xbc, 0x0b, 0x22, 0xae, 0x9e, 0xee, 0x5e, 0xde,
	0xda, 0xcd, 0x7f, 0x01, 0x3d, 0xd3, 0xc0, 0x8a, 0x2f, 0xd6, 0xc1, 0xed, 0x8e, 0xfd, 0x58, 0xbd,
	0xf8, 0x5b, 0xbd, 0x9d, 0x1e, 0x1c, 0x4e, 0xc9, 0xa2, 0x5d, 0x58, 0x1d, 0x7f, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x74, 0x7f, 0x5c, 0x0b, 0xd8, 0x01, 0x00, 0x00,
}
