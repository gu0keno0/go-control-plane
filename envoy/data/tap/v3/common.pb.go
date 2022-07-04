// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.21.2
// source: envoy/data/tap/v3/common.proto

package tapv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// Wrapper for tapped body data. This includes HTTP request/response body, transport socket received
// and transmitted data, etc.
type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to BodyType:
	//	*Body_AsBytes
	//	*Body_AsString
	BodyType isBody_BodyType `protobuf_oneof:"body_type"`
	// Specifies whether body data has been truncated to fit within the specified
	// :ref:`max_buffered_rx_bytes
	// <envoy_v3_api_field_config.tap.v3.OutputConfig.max_buffered_rx_bytes>` and
	// :ref:`max_buffered_tx_bytes
	// <envoy_v3_api_field_config.tap.v3.OutputConfig.max_buffered_tx_bytes>` settings.
	Truncated bool `protobuf:"varint,3,opt,name=truncated,proto3" json:"truncated,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v3_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v3_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v3_common_proto_rawDescGZIP(), []int{0}
}

func (m *Body) GetBodyType() isBody_BodyType {
	if m != nil {
		return m.BodyType
	}
	return nil
}

func (x *Body) GetAsBytes() []byte {
	if x, ok := x.GetBodyType().(*Body_AsBytes); ok {
		return x.AsBytes
	}
	return nil
}

func (x *Body) GetAsString() string {
	if x, ok := x.GetBodyType().(*Body_AsString); ok {
		return x.AsString
	}
	return ""
}

func (x *Body) GetTruncated() bool {
	if x != nil {
		return x.Truncated
	}
	return false
}

type isBody_BodyType interface {
	isBody_BodyType()
}

type Body_AsBytes struct {
	// Body data as bytes. By default, tap body data will be present in this field, as the proto
	// `bytes` type can contain any valid byte.
	AsBytes []byte `protobuf:"bytes,1,opt,name=as_bytes,json=asBytes,proto3,oneof"`
}

type Body_AsString struct {
	// Body data as string. This field is only used when the :ref:`JSON_BODY_AS_STRING
	// <envoy_v3_api_enum_value_config.tap.v3.OutputSink.Format.JSON_BODY_AS_STRING>` sink
	// format type is selected. See the documentation for that option for why this is useful.
	AsString string `protobuf:"bytes,2,opt,name=as_string,json=asString,proto3,oneof"`
}

func (*Body_AsBytes) isBody_BodyType() {}

func (*Body_AsString) isBody_BodyType() {}

var File_envoy_data_tap_v3_common_proto protoreflect.FileDescriptor

var file_envoy_data_tap_v3_common_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x70,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1b,
	0x0a, 0x08, 0x61, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x07, 0x61, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x61,
	0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x61, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72,
	0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74,
	0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x22, 0x9a, 0xc5, 0x88, 0x1e, 0x1d, 0x0a,
	0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x42, 0x0b, 0x0a, 0x09,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x78, 0x0a, 0x1f, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74,
	0x61, 0x70, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x61, 0x70, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_tap_v3_common_proto_rawDescOnce sync.Once
	file_envoy_data_tap_v3_common_proto_rawDescData = file_envoy_data_tap_v3_common_proto_rawDesc
)

func file_envoy_data_tap_v3_common_proto_rawDescGZIP() []byte {
	file_envoy_data_tap_v3_common_proto_rawDescOnce.Do(func() {
		file_envoy_data_tap_v3_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_tap_v3_common_proto_rawDescData)
	})
	return file_envoy_data_tap_v3_common_proto_rawDescData
}

var file_envoy_data_tap_v3_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_data_tap_v3_common_proto_goTypes = []interface{}{
	(*Body)(nil), // 0: envoy.data.tap.v3.Body
}
var file_envoy_data_tap_v3_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_data_tap_v3_common_proto_init() }
func file_envoy_data_tap_v3_common_proto_init() {
	if File_envoy_data_tap_v3_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_tap_v3_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Body); i {
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
	file_envoy_data_tap_v3_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Body_AsBytes)(nil),
		(*Body_AsString)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_tap_v3_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_tap_v3_common_proto_goTypes,
		DependencyIndexes: file_envoy_data_tap_v3_common_proto_depIdxs,
		MessageInfos:      file_envoy_data_tap_v3_common_proto_msgTypes,
	}.Build()
	File_envoy_data_tap_v3_common_proto = out.File
	file_envoy_data_tap_v3_common_proto_rawDesc = nil
	file_envoy_data_tap_v3_common_proto_goTypes = nil
	file_envoy_data_tap_v3_common_proto_depIdxs = nil
}
