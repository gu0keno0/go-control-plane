// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.21.2
// source: envoy/extensions/network/dns_resolver/cares/v3/cares_dns_resolver.proto

package caresv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// Configuration for c-ares DNS resolver.
type CaresDnsResolverConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of dns resolver addresses.
	// :ref:`use_resolvers_as_fallback<envoy_v3_api_field_extensions.network.dns_resolver.cares.v3.CaresDnsResolverConfig.use_resolvers_as_fallback>`
	// below dictates if the DNS client should override system defaults or only use the provided
	// resolvers if the system defaults are not available, i.e., as a fallback.
	Resolvers []*v3.Address `protobuf:"bytes,1,rep,name=resolvers,proto3" json:"resolvers,omitempty"`
	// If true use the resolvers listed in the
	// :ref:`resolvers<envoy_v3_api_field_extensions.network.dns_resolver.cares.v3.CaresDnsResolverConfig.resolvers>`
	// field only if c-ares is unable to obtain a
	// nameserver from the system (e.g., /etc/resolv.conf).
	// Otherwise, the resolvers listed in the resolvers list will override the default system
	// resolvers. Defaults to false.
	UseResolversAsFallback bool `protobuf:"varint,3,opt,name=use_resolvers_as_fallback,json=useResolversAsFallback,proto3" json:"use_resolvers_as_fallback,omitempty"`
	// The resolver will query available network interfaces and determine if there are no available
	// interfaces for a given IP family. It will then filter these addresses from the results it
	// presents. e.g., if there are no available IPv4 network interfaces, the resolver will not
	// provide IPv4 addresses.
	FilterUnroutableFamilies bool `protobuf:"varint,4,opt,name=filter_unroutable_families,json=filterUnroutableFamilies,proto3" json:"filter_unroutable_families,omitempty"`
	// Configuration of DNS resolver option flags which control the behavior of the DNS resolver.
	DnsResolverOptions *v3.DnsResolverOptions `protobuf:"bytes,2,opt,name=dns_resolver_options,json=dnsResolverOptions,proto3" json:"dns_resolver_options,omitempty"`
}

func (x *CaresDnsResolverConfig) Reset() {
	*x = CaresDnsResolverConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaresDnsResolverConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaresDnsResolverConfig) ProtoMessage() {}

func (x *CaresDnsResolverConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaresDnsResolverConfig.ProtoReflect.Descriptor instead.
func (*CaresDnsResolverConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescGZIP(), []int{0}
}

func (x *CaresDnsResolverConfig) GetResolvers() []*v3.Address {
	if x != nil {
		return x.Resolvers
	}
	return nil
}

func (x *CaresDnsResolverConfig) GetUseResolversAsFallback() bool {
	if x != nil {
		return x.UseResolversAsFallback
	}
	return false
}

func (x *CaresDnsResolverConfig) GetFilterUnroutableFamilies() bool {
	if x != nil {
		return x.FilterUnroutableFamilies
	}
	return false
}

func (x *CaresDnsResolverConfig) GetDnsResolverOptions() *v3.DnsResolverOptions {
	if x != nil {
		return x.DnsResolverOptions
	}
	return nil
}

var File_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto protoreflect.FileDescriptor

var file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDesc = []byte{
	0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x61, 0x72, 0x65, 0x73, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x2e, 0x63, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x16, 0x43,
	0x61, 0x72, 0x65, 0x73, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x19,
	0x75, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x73,
	0x5f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x41, 0x73, 0x46,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x1a, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x5f, 0x75, 0x6e, 0x72, 0x6f, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x55, 0x6e, 0x72, 0x6f, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x14, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x64,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0xbe, 0x01, 0x0a, 0x3c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e, 0x73,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x73, 0x2e,
	0x76, 0x33, 0x42, 0x15, 0x43, 0x61, 0x72, 0x65, 0x73, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x73, 0x2f,
	0x76, 0x33, 0x3b, 0x63, 0x61, 0x72, 0x65, 0x73, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescOnce sync.Once
	file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescData = file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDesc
)

func file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescGZIP() []byte {
	file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescData)
	})
	return file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDescData
}

var file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_goTypes = []interface{}{
	(*CaresDnsResolverConfig)(nil), // 0: envoy.extensions.network.dns_resolver.cares.v3.CaresDnsResolverConfig
	(*v3.Address)(nil),             // 1: envoy.config.core.v3.Address
	(*v3.DnsResolverOptions)(nil),  // 2: envoy.config.core.v3.DnsResolverOptions
}
var file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.network.dns_resolver.cares.v3.CaresDnsResolverConfig.resolvers:type_name -> envoy.config.core.v3.Address
	2, // 1: envoy.extensions.network.dns_resolver.cares.v3.CaresDnsResolverConfig.dns_resolver_options:type_name -> envoy.config.core.v3.DnsResolverOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_init() }
func file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_init() {
	if File_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaresDnsResolverConfig); i {
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
			RawDescriptor: file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_msgTypes,
	}.Build()
	File_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto = out.File
	file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_rawDesc = nil
	file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_goTypes = nil
	file_envoy_extensions_network_dns_resolver_cares_v3_cares_dns_resolver_proto_depIdxs = nil
}
