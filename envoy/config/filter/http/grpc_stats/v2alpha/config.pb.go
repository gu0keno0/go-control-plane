// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.21.2
// source: envoy/config/filter/http/grpc_stats/v2alpha/config.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// gRPC statistics filter configuration
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, the filter maintains a filter state object with the request and response message
	// counts.
	EmitFilterState bool `protobuf:"varint,1,opt,name=emit_filter_state,json=emitFilterState,proto3" json:"emit_filter_state,omitempty"`
	// Types that are assignable to PerMethodStatSpecifier:
	//	*FilterConfig_IndividualMethodStatsAllowlist
	//	*FilterConfig_StatsForAllMethods
	PerMethodStatSpecifier isFilterConfig_PerMethodStatSpecifier `protobuf_oneof:"per_method_stat_specifier"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetEmitFilterState() bool {
	if x != nil {
		return x.EmitFilterState
	}
	return false
}

func (m *FilterConfig) GetPerMethodStatSpecifier() isFilterConfig_PerMethodStatSpecifier {
	if m != nil {
		return m.PerMethodStatSpecifier
	}
	return nil
}

func (x *FilterConfig) GetIndividualMethodStatsAllowlist() *core.GrpcMethodList {
	if x, ok := x.GetPerMethodStatSpecifier().(*FilterConfig_IndividualMethodStatsAllowlist); ok {
		return x.IndividualMethodStatsAllowlist
	}
	return nil
}

func (x *FilterConfig) GetStatsForAllMethods() *wrapperspb.BoolValue {
	if x, ok := x.GetPerMethodStatSpecifier().(*FilterConfig_StatsForAllMethods); ok {
		return x.StatsForAllMethods
	}
	return nil
}

type isFilterConfig_PerMethodStatSpecifier interface {
	isFilterConfig_PerMethodStatSpecifier()
}

type FilterConfig_IndividualMethodStatsAllowlist struct {
	// If set, specifies an allowlist of service/methods that will have individual stats
	// emitted for them. Any call that does not match the allowlist will be counted
	// in a stat with no method specifier: `cluster.<name>.grpc.*`.
	IndividualMethodStatsAllowlist *core.GrpcMethodList `protobuf:"bytes,2,opt,name=individual_method_stats_allowlist,json=individualMethodStatsAllowlist,proto3,oneof"`
}

type FilterConfig_StatsForAllMethods struct {
	// If set to true, emit stats for all service/method names.
	//
	// If set to false, emit stats for all service/message types to the same stats without including
	// the service/method in the name, with prefix `cluster.<name>.grpc`. This can be useful if
	// service/method granularity is not needed, or if each cluster only receives a single method.
	//
	// .. attention::
	//   This option is only safe if all clients are trusted. If this option is enabled
	//   with untrusted clients, the clients could cause unbounded growth in the number of stats in
	//   Envoy, using unbounded memory and potentially slowing down stats pipelines.
	//
	// .. attention::
	//   If neither `individual_method_stats_allowlist` nor `stats_for_all_methods` is set, the
	//   behavior will default to `stats_for_all_methods=false`. This default value is changed due
	//   to the previous value being deprecated. This behavior can be changed with runtime override
	//   `envoy.deprecated_features.grpc_stats_filter_enable_stats_for_all_methods_by_default`.
	StatsForAllMethods *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=stats_for_all_methods,json=statsForAllMethods,proto3,oneof"`
}

func (*FilterConfig_IndividualMethodStatsAllowlist) isFilterConfig_PerMethodStatSpecifier() {}

func (*FilterConfig_StatsForAllMethods) isFilterConfig_PerMethodStatSpecifier() {}

// gRPC statistics filter state object in protobuf form.
type FilterObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count of request messages in the request stream.
	RequestMessageCount uint64 `protobuf:"varint,1,opt,name=request_message_count,json=requestMessageCount,proto3" json:"request_message_count,omitempty"`
	// Count of response messages in the response stream.
	ResponseMessageCount uint64 `protobuf:"varint,2,opt,name=response_message_count,json=responseMessageCount,proto3" json:"response_message_count,omitempty"`
}

func (x *FilterObject) Reset() {
	*x = FilterObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterObject) ProtoMessage() {}

func (x *FilterObject) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterObject.ProtoReflect.Descriptor instead.
func (*FilterObject) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescGZIP(), []int{1}
}

func (x *FilterObject) GetRequestMessageCount() uint64 {
	if x != nil {
		return x.RequestMessageCount
	}
	return 0
}

func (x *FilterObject) GetResponseMessageCount() uint64 {
	if x != nil {
		return x.ResponseMessageCount
	}
	return 0
}

var File_envoy_config_filter_http_grpc_stats_v2alpha_config_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x98, 0x02, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6d, 0x69, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x6d,
	0x69, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x6e, 0x0a,
	0x21, 0x69, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x1e, 0x69,
	0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a,
	0x15, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x12, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x42, 0x1b,
	0x0a, 0x19, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x0c, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x34, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xd9, 0x01, 0x0a, 0x39, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x2d, 0x12, 0x2b, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescData = file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDesc
)

func file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescData)
	})
	return file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDescData
}

var file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_goTypes = []interface{}{
	(*FilterConfig)(nil),         // 0: envoy.config.filter.http.grpc_stats.v2alpha.FilterConfig
	(*FilterObject)(nil),         // 1: envoy.config.filter.http.grpc_stats.v2alpha.FilterObject
	(*core.GrpcMethodList)(nil),  // 2: envoy.api.v2.core.GrpcMethodList
	(*wrapperspb.BoolValue)(nil), // 3: google.protobuf.BoolValue
}
var file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_depIdxs = []int32{
	2, // 0: envoy.config.filter.http.grpc_stats.v2alpha.FilterConfig.individual_method_stats_allowlist:type_name -> envoy.api.v2.core.GrpcMethodList
	3, // 1: envoy.config.filter.http.grpc_stats.v2alpha.FilterConfig.stats_for_all_methods:type_name -> google.protobuf.BoolValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_init() }
func file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_init() {
	if File_envoy_config_filter_http_grpc_stats_v2alpha_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
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
		file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterObject); i {
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
	file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FilterConfig_IndividualMethodStatsAllowlist)(nil),
		(*FilterConfig_StatsForAllMethods)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_grpc_stats_v2alpha_config_proto = out.File
	file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_rawDesc = nil
	file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_goTypes = nil
	file_envoy_config_filter_http_grpc_stats_v2alpha_config_proto_depIdxs = nil
}
