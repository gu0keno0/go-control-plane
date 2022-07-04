// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.21.2
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package bufferv2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Buffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
}

func (x *Buffer) Reset() {
	*x = Buffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Buffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buffer) ProtoMessage() {}

func (x *Buffer) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buffer.ProtoReflect.Descriptor instead.
func (*Buffer) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescGZIP(), []int{0}
}

func (x *Buffer) GetMaxRequestBytes() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override isBufferPerRoute_Override `protobuf_oneof:"override"`
}

func (x *BufferPerRoute) Reset() {
	*x = BufferPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BufferPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BufferPerRoute) ProtoMessage() {}

func (x *BufferPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BufferPerRoute.ProtoReflect.Descriptor instead.
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescGZIP(), []int{1}
}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (x *BufferPerRoute) GetDisabled() bool {
	if x, ok := x.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (x *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := x.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
}

type BufferPerRoute_Disabled struct {
	// Disable the buffer filter for this particular vhost or route.
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type BufferPerRoute_Buffer struct {
	// Override the global configuration of the filter with this new config.
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}

func (*BufferPerRoute_Buffer) isBufferPerRoute_Override() {}

var File_envoy_config_filter_http_buffer_v2_buffer_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a,
	0x06, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0f, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x98, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x48, 0x00, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x4e, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x42, 0x0f, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x42, 0xcc, 0x01, 0x0a, 0x30, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2f,
	0x76, 0x32, 0x3b, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x76, 0x32, 0xf2, 0x98, 0xfe, 0x8f, 0x05,
	0x29, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescData = file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDesc
)

func file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescData)
	})
	return file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDescData
}

var file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_buffer_v2_buffer_proto_goTypes = []interface{}{
	(*Buffer)(nil),                 // 0: envoy.config.filter.http.buffer.v2.Buffer
	(*BufferPerRoute)(nil),         // 1: envoy.config.filter.http.buffer.v2.BufferPerRoute
	(*wrapperspb.UInt32Value)(nil), // 2: google.protobuf.UInt32Value
}
var file_envoy_config_filter_http_buffer_v2_buffer_proto_depIdxs = []int32{
	2, // 0: envoy.config.filter.http.buffer.v2.Buffer.max_request_bytes:type_name -> google.protobuf.UInt32Value
	0, // 1: envoy.config.filter.http.buffer.v2.BufferPerRoute.buffer:type_name -> envoy.config.filter.http.buffer.v2.Buffer
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_buffer_v2_buffer_proto_init() }
func file_envoy_config_filter_http_buffer_v2_buffer_proto_init() {
	if File_envoy_config_filter_http_buffer_v2_buffer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Buffer); i {
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
		file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BufferPerRoute); i {
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
	file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_buffer_v2_buffer_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_buffer_v2_buffer_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_buffer_v2_buffer_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_buffer_v2_buffer_proto = out.File
	file_envoy_config_filter_http_buffer_v2_buffer_proto_rawDesc = nil
	file_envoy_config_filter_http_buffer_v2_buffer_proto_goTypes = nil
	file_envoy_config_filter_http_buffer_v2_buffer_proto_depIdxs = nil
}
