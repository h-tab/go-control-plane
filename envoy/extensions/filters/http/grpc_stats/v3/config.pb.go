// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/grpc_stats/v3/config.proto

package envoy_extensions_filters_http_grpc_stats_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

type FilterConfig struct {
	EmitFilterState bool `protobuf:"varint,1,opt,name=emit_filter_state,json=emitFilterState,proto3" json:"emit_filter_state,omitempty"`
	// Types that are valid to be assigned to PerMethodStatSpecifier:
	//	*FilterConfig_IndividualMethodStatsAllowlist
	//	*FilterConfig_StatsForAllMethods
	PerMethodStatSpecifier isFilterConfig_PerMethodStatSpecifier `protobuf_oneof:"per_method_stat_specifier"`
	EnableUpstreamStats    bool                                  `protobuf:"varint,4,opt,name=enable_upstream_stats,json=enableUpstreamStats,proto3" json:"enable_upstream_stats,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                              `json:"-"`
	XXX_unrecognized       []byte                                `json:"-"`
	XXX_sizecache          int32                                 `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a99092111d6420a4, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetEmitFilterState() bool {
	if m != nil {
		return m.EmitFilterState
	}
	return false
}

type isFilterConfig_PerMethodStatSpecifier interface {
	isFilterConfig_PerMethodStatSpecifier()
}

type FilterConfig_IndividualMethodStatsAllowlist struct {
	IndividualMethodStatsAllowlist *v3.GrpcMethodList `protobuf:"bytes,2,opt,name=individual_method_stats_allowlist,json=individualMethodStatsAllowlist,proto3,oneof"`
}

type FilterConfig_StatsForAllMethods struct {
	StatsForAllMethods *wrappers.BoolValue `protobuf:"bytes,3,opt,name=stats_for_all_methods,json=statsForAllMethods,proto3,oneof"`
}

func (*FilterConfig_IndividualMethodStatsAllowlist) isFilterConfig_PerMethodStatSpecifier() {}

func (*FilterConfig_StatsForAllMethods) isFilterConfig_PerMethodStatSpecifier() {}

func (m *FilterConfig) GetPerMethodStatSpecifier() isFilterConfig_PerMethodStatSpecifier {
	if m != nil {
		return m.PerMethodStatSpecifier
	}
	return nil
}

func (m *FilterConfig) GetIndividualMethodStatsAllowlist() *v3.GrpcMethodList {
	if x, ok := m.GetPerMethodStatSpecifier().(*FilterConfig_IndividualMethodStatsAllowlist); ok {
		return x.IndividualMethodStatsAllowlist
	}
	return nil
}

func (m *FilterConfig) GetStatsForAllMethods() *wrappers.BoolValue {
	if x, ok := m.GetPerMethodStatSpecifier().(*FilterConfig_StatsForAllMethods); ok {
		return x.StatsForAllMethods
	}
	return nil
}

func (m *FilterConfig) GetEnableUpstreamStats() bool {
	if m != nil {
		return m.EnableUpstreamStats
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FilterConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FilterConfig_IndividualMethodStatsAllowlist)(nil),
		(*FilterConfig_StatsForAllMethods)(nil),
	}
}

type FilterObject struct {
	RequestMessageCount  uint64   `protobuf:"varint,1,opt,name=request_message_count,json=requestMessageCount,proto3" json:"request_message_count,omitempty"`
	ResponseMessageCount uint64   `protobuf:"varint,2,opt,name=response_message_count,json=responseMessageCount,proto3" json:"response_message_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterObject) Reset()         { *m = FilterObject{} }
func (m *FilterObject) String() string { return proto.CompactTextString(m) }
func (*FilterObject) ProtoMessage()    {}
func (*FilterObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_a99092111d6420a4, []int{1}
}

func (m *FilterObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterObject.Unmarshal(m, b)
}
func (m *FilterObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterObject.Marshal(b, m, deterministic)
}
func (m *FilterObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterObject.Merge(m, src)
}
func (m *FilterObject) XXX_Size() int {
	return xxx_messageInfo_FilterObject.Size(m)
}
func (m *FilterObject) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterObject.DiscardUnknown(m)
}

var xxx_messageInfo_FilterObject proto.InternalMessageInfo

func (m *FilterObject) GetRequestMessageCount() uint64 {
	if m != nil {
		return m.RequestMessageCount
	}
	return 0
}

func (m *FilterObject) GetResponseMessageCount() uint64 {
	if m != nil {
		return m.ResponseMessageCount
	}
	return 0
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.extensions.filters.http.grpc_stats.v3.FilterConfig")
	proto.RegisterType((*FilterObject)(nil), "envoy.extensions.filters.http.grpc_stats.v3.FilterObject")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/grpc_stats/v3/config.proto", fileDescriptor_a99092111d6420a4)
}

var fileDescriptor_a99092111d6420a4 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x5d, 0xba, 0x69, 0x42, 0x19, 0x12, 0x90, 0x31, 0x28, 0x45, 0x54, 0xdb, 0xc4, 0x61, 0x02,
	0xc9, 0x96, 0x5a, 0x0e, 0xb0, 0x0b, 0x5a, 0x27, 0x8d, 0x1d, 0x98, 0x36, 0x05, 0xc1, 0x35, 0x72,
	0x93, 0x2f, 0xa9, 0x91, 0xeb, 0xcf, 0xb3, 0x9d, 0x6c, 0xbb, 0x71, 0xe4, 0xc8, 0x99, 0x7f, 0x02,
	0x77, 0x24, 0xae, 0xfc, 0x23, 0x64, 0x3b, 0x51, 0x2b, 0x7a, 0x42, 0xbb, 0xb5, 0x7e, 0x7e, 0xef,
	0x7b, 0xef, 0x7d, 0x4e, 0xfc, 0x1a, 0x64, 0x83, 0x37, 0x14, 0xae, 0x2d, 0x48, 0xc3, 0x51, 0x1a,
	0x5a, 0x72, 0x61, 0x41, 0x1b, 0x3a, 0xb3, 0x56, 0xd1, 0x4a, 0xab, 0x3c, 0x33, 0x96, 0x59, 0x43,
	0x9b, 0x31, 0xcd, 0x51, 0x96, 0xbc, 0x22, 0x4a, 0xa3, 0xc5, 0xe4, 0xa5, 0x67, 0x92, 0x05, 0x93,
	0xb4, 0x4c, 0xe2, 0x98, 0x64, 0xc1, 0x24, 0xcd, 0x78, 0x10, 0x2e, 0xb7, 0x02, 0x34, 0x47, 0x0d,
	0x4e, 0xcf, 0xdf, 0x99, 0x83, 0x9d, 0x61, 0x91, 0x09, 0x6e, 0x6c, 0x50, 0x1e, 0x0c, 0x2b, 0xc4,
	0x4a, 0x00, 0xf5, 0xff, 0xa6, 0x75, 0x49, 0xaf, 0x34, 0x53, 0xca, 0x29, 0x07, 0xfc, 0x59, 0x5d,
	0x28, 0x46, 0x99, 0x94, 0x68, 0x99, 0xf5, 0x9e, 0xdd, 0x9c, 0xba, 0x83, 0xf7, 0x56, 0xe0, 0x06,
	0xb4, 0x73, 0xc8, 0x65, 0xeb, 0x7d, 0xf0, 0xb8, 0x61, 0x82, 0x17, 0xcc, 0x02, 0xed, 0x7e, 0x04,
	0x60, 0xff, 0xdb, 0x7a, 0x7c, 0xf7, 0xc4, 0xc7, 0x38, 0xf6, 0x56, 0x93, 0x17, 0xf1, 0x03, 0x98,
	0x73, 0x9b, 0x85, 0x6c, 0x3e, 0x10, 0xf4, 0xa3, 0xdd, 0xe8, 0xe0, 0x4e, 0x7a, 0xcf, 0x01, 0xe1,
	0xf2, 0x07, 0x77, 0x9c, 0x5c, 0xc6, 0x7b, 0x5c, 0x16, 0xbc, 0xe1, 0x45, 0xcd, 0x44, 0x97, 0xcb,
	0x57, 0x90, 0x31, 0x21, 0xf0, 0xca, 0x45, 0xec, 0xf7, 0x76, 0xa3, 0x83, 0xad, 0xd1, 0x73, 0x12,
	0xda, 0x6b, 0x1b, 0x75, 0x85, 0x90, 0x66, 0x4c, 0xde, 0x69, 0x95, 0x9f, 0x79, 0xde, 0x7b, 0x6e,
	0xec, 0xe9, 0x5a, 0x3a, 0x5c, 0x08, 0x86, 0x73, 0x37, 0xc9, 0x1c, 0x75, 0x6a, 0xc9, 0x79, 0xbc,
	0x13, 0x06, 0x94, 0xa8, 0xdd, 0x90, 0x76, 0xaa, 0xe9, 0xaf, 0xfb, 0x31, 0x03, 0x12, 0xaa, 0x24,
	0x5d, 0x95, 0x64, 0x82, 0x28, 0x3e, 0x31, 0x51, 0xc3, 0xe9, 0x5a, 0x9a, 0x78, 0xea, 0x09, 0xea,
	0x23, 0xd1, 0xaa, 0x9b, 0x64, 0x14, 0xef, 0x80, 0x64, 0x53, 0x01, 0x59, 0xad, 0x8c, 0xd5, 0xc0,
	0xe6, 0x21, 0x41, 0x7f, 0xc3, 0x67, 0xde, 0x0e, 0xe0, 0xc7, 0x16, 0xf3, 0x6e, 0x0e, 0xdf, 0x7e,
	0xff, 0xf5, 0x75, 0x78, 0xd8, 0x3e, 0xa5, 0x2e, 0x52, 0x28, 0x6c, 0xf5, 0x2d, 0x8c, 0x98, 0x50,
	0x33, 0x46, 0x96, 0x4b, 0x9e, 0x3c, 0x8d, 0x9f, 0x28, 0xd0, 0xcb, 0x8d, 0x65, 0x46, 0x41, 0xce,
	0x4b, 0x0e, 0x7a, 0xff, 0x47, 0xd4, 0xad, 0xe4, 0x7c, 0xfa, 0x19, 0x72, 0xeb, 0x2c, 0x6a, 0xb8,
	0xac, 0xc1, 0xd8, 0x6c, 0x0e, 0xc6, 0xb0, 0x0a, 0xb2, 0x1c, 0x6b, 0x69, 0xfd, 0x5a, 0x36, 0xd2,
	0xed, 0x16, 0x3c, 0x0b, 0xd8, 0xb1, 0x83, 0x92, 0x57, 0xf1, 0x23, 0x0d, 0x46, 0xa1, 0x34, 0xf0,
	0x0f, 0xa9, 0xe7, 0x49, 0x0f, 0x3b, 0x74, 0x99, 0x75, 0x8b, 0x60, 0xc1, 0xea, 0x24, 0xfd, 0xf9,
	0xe5, 0xf7, 0x9f, 0xcd, 0xde, 0xfd, 0x5e, 0xfc, 0x86, 0x63, 0x58, 0xb9, 0xd2, 0x78, 0x7d, 0x43,
	0xfe, 0xe3, 0xdb, 0x99, 0x6c, 0x85, 0x96, 0x2e, 0xdc, 0x0a, 0x2f, 0xa2, 0xe9, 0xa6, 0xdf, 0xe5,
	0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xff, 0xc7, 0x40, 0xba, 0x03, 0x00, 0x00,
}
