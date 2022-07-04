// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.21.2
// source: envoy/extensions/load_balancing_policies/wrr_locality/v3/wrr_locality.proto

package wrr_localityv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for the wrr_locality LB policy. See the :ref:`load balancing architecture overview
// <arch_overview_load_balancing_types>` for more information.
// [#extension: envoy.clusters.lb_policy]
type WrrLocality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The child LB policy to create for endpoint-picking within the chosen locality.
	EndpointPickingPolicy *v3.LoadBalancingPolicy `protobuf:"bytes,1,opt,name=endpoint_picking_policy,json=endpointPickingPolicy,proto3" json:"endpoint_picking_policy,omitempty"`
}

func (x *WrrLocality) Reset() {
	*x = WrrLocality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrrLocality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrrLocality) ProtoMessage() {}

func (x *WrrLocality) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrrLocality.ProtoReflect.Descriptor instead.
func (*WrrLocality) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescGZIP(), []int{0}
}

func (x *WrrLocality) GetEndpointPickingPolicy() *v3.LoadBalancingPolicy {
	if x != nil {
		return x.EndpointPickingPolicy
	}
	return nil
}

var File_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x77, 0x72, 0x72, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x77, 0x72, 0x72, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x77, 0x72, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0b, 0x57, 0x72, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x6e, 0x0a, 0x17, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x15,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0xd4, 0x01, 0x0a, 0x46, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2e, 0x77, 0x72, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x33,
	0x42, 0x10, 0x57, 0x72, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x77, 0x72, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x77, 0x72, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescData = file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_goTypes = []interface{}{
	(*WrrLocality)(nil),            // 0: envoy.extensions.load_balancing_policies.wrr_locality.v3.WrrLocality
	(*v3.LoadBalancingPolicy)(nil), // 1: envoy.config.cluster.v3.LoadBalancingPolicy
}
var file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.load_balancing_policies.wrr_locality.v3.WrrLocality.endpoint_picking_policy:type_name -> envoy.config.cluster.v3.LoadBalancingPolicy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_init() }
func file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_init() {
	if File_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrrLocality); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto = out.File
	file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_wrr_locality_v3_wrr_locality_proto_depIdxs = nil
}
