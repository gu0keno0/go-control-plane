// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.21.2
// source: envoy/extensions/filters/common/dependency/v3/dependency.proto

package dependencyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

type Dependency_DependencyType int32

const (
	Dependency_HEADER           Dependency_DependencyType = 0
	Dependency_FILTER_STATE_KEY Dependency_DependencyType = 1
	Dependency_DYNAMIC_METADATA Dependency_DependencyType = 2
)

// Enum value maps for Dependency_DependencyType.
var (
	Dependency_DependencyType_name = map[int32]string{
		0: "HEADER",
		1: "FILTER_STATE_KEY",
		2: "DYNAMIC_METADATA",
	}
	Dependency_DependencyType_value = map[string]int32{
		"HEADER":           0,
		"FILTER_STATE_KEY": 1,
		"DYNAMIC_METADATA": 2,
	}
)

func (x Dependency_DependencyType) Enum() *Dependency_DependencyType {
	p := new(Dependency_DependencyType)
	*p = x
	return p
}

func (x Dependency_DependencyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Dependency_DependencyType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_common_dependency_v3_dependency_proto_enumTypes[0].Descriptor()
}

func (Dependency_DependencyType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_enumTypes[0]
}

func (x Dependency_DependencyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Dependency_DependencyType.Descriptor instead.
func (Dependency_DependencyType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescGZIP(), []int{0, 0}
}

// Dependency specification and string identifier.
type Dependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of dependency.
	Type Dependency_DependencyType `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.extensions.filters.common.dependency.v3.Dependency_DependencyType" json:"type,omitempty"`
	// The string identifier for the dependency.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Dependency) Reset() {
	*x = Dependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dependency) ProtoMessage() {}

func (x *Dependency) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dependency.ProtoReflect.Descriptor instead.
func (*Dependency) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescGZIP(), []int{0}
}

func (x *Dependency) GetType() Dependency_DependencyType {
	if x != nil {
		return x.Type
	}
	return Dependency_HEADER
}

func (x *Dependency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Dependency specification for a filter. For a filter chain to be valid, any
// dependency that is required must be provided by an earlier filter.
type FilterDependencies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of dependencies required on the decode path.
	DecodeRequired []*Dependency `protobuf:"bytes,1,rep,name=decode_required,json=decodeRequired,proto3" json:"decode_required,omitempty"`
	// A list of dependencies provided on the encode path.
	DecodeProvided []*Dependency `protobuf:"bytes,2,rep,name=decode_provided,json=decodeProvided,proto3" json:"decode_provided,omitempty"`
	// A list of dependencies required on the decode path.
	EncodeRequired []*Dependency `protobuf:"bytes,3,rep,name=encode_required,json=encodeRequired,proto3" json:"encode_required,omitempty"`
	// A list of dependencies provided on the encode path.
	EncodeProvided []*Dependency `protobuf:"bytes,4,rep,name=encode_provided,json=encodeProvided,proto3" json:"encode_provided,omitempty"`
}

func (x *FilterDependencies) Reset() {
	*x = FilterDependencies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterDependencies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterDependencies) ProtoMessage() {}

func (x *FilterDependencies) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterDependencies.ProtoReflect.Descriptor instead.
func (*FilterDependencies) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescGZIP(), []int{1}
}

func (x *FilterDependencies) GetDecodeRequired() []*Dependency {
	if x != nil {
		return x.DecodeRequired
	}
	return nil
}

func (x *FilterDependencies) GetDecodeProvided() []*Dependency {
	if x != nil {
		return x.DecodeProvided
	}
	return nil
}

func (x *FilterDependencies) GetEncodeRequired() []*Dependency {
	if x != nil {
		return x.EncodeRequired
	}
	return nil
}

func (x *FilterDependencies) GetEncodeProvided() []*Dependency {
	if x != nil {
		return x.EncodeProvided
	}
	return nil
}

// Matching requirements for a filter. For a match tree to be used with a filter, the match
// requirements must be satisfied.
//
// This protobuf is provided by the filter implementation as a way to communicate the matching
// requirements to the filter factories, allowing for config rejection if the requirements are
// not satisfied.
type MatchingRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataInputAllowList *MatchingRequirements_DataInputAllowList `protobuf:"bytes,1,opt,name=data_input_allow_list,json=dataInputAllowList,proto3" json:"data_input_allow_list,omitempty"`
}

func (x *MatchingRequirements) Reset() {
	*x = MatchingRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchingRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingRequirements) ProtoMessage() {}

func (x *MatchingRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingRequirements.ProtoReflect.Descriptor instead.
func (*MatchingRequirements) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescGZIP(), []int{2}
}

func (x *MatchingRequirements) GetDataInputAllowList() *MatchingRequirements_DataInputAllowList {
	if x != nil {
		return x.DataInputAllowList
	}
	return nil
}

type MatchingRequirements_DataInputAllowList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An explicit list of data inputs that are allowed to be used with this filter.
	TypeUrl []string `protobuf:"bytes,1,rep,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
}

func (x *MatchingRequirements_DataInputAllowList) Reset() {
	*x = MatchingRequirements_DataInputAllowList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchingRequirements_DataInputAllowList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingRequirements_DataInputAllowList) ProtoMessage() {}

func (x *MatchingRequirements_DataInputAllowList) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingRequirements_DataInputAllowList.ProtoReflect.Descriptor instead.
func (*MatchingRequirements_DataInputAllowList) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MatchingRequirements_DataInputAllowList) GetTypeUrl() []string {
	if x != nil {
		return x.TypeUrl
	}
	return nil
}

var File_envoy_extensions_filters_common_dependency_v3_dependency_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x33, 0x2f,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x5c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x48, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x4b, 0x45, 0x59, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43,
	0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x22, 0xa4, 0x03, 0x0a, 0x12,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x62, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x62, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x12, 0x62, 0x0a, 0x0f, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0e,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x62,
	0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x15,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x2f, 0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x42, 0xbb, 0x01, 0x0a, 0x3b, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0f, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76,
	0x33, 0x3b, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescData = file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDesc
)

func file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescData)
	})
	return file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDescData
}

var file_envoy_extensions_filters_common_dependency_v3_dependency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_common_dependency_v3_dependency_proto_goTypes = []interface{}{
	(Dependency_DependencyType)(0),                  // 0: envoy.extensions.filters.common.dependency.v3.Dependency.DependencyType
	(*Dependency)(nil),                              // 1: envoy.extensions.filters.common.dependency.v3.Dependency
	(*FilterDependencies)(nil),                      // 2: envoy.extensions.filters.common.dependency.v3.FilterDependencies
	(*MatchingRequirements)(nil),                    // 3: envoy.extensions.filters.common.dependency.v3.MatchingRequirements
	(*MatchingRequirements_DataInputAllowList)(nil), // 4: envoy.extensions.filters.common.dependency.v3.MatchingRequirements.DataInputAllowList
}
var file_envoy_extensions_filters_common_dependency_v3_dependency_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.filters.common.dependency.v3.Dependency.type:type_name -> envoy.extensions.filters.common.dependency.v3.Dependency.DependencyType
	1, // 1: envoy.extensions.filters.common.dependency.v3.FilterDependencies.decode_required:type_name -> envoy.extensions.filters.common.dependency.v3.Dependency
	1, // 2: envoy.extensions.filters.common.dependency.v3.FilterDependencies.decode_provided:type_name -> envoy.extensions.filters.common.dependency.v3.Dependency
	1, // 3: envoy.extensions.filters.common.dependency.v3.FilterDependencies.encode_required:type_name -> envoy.extensions.filters.common.dependency.v3.Dependency
	1, // 4: envoy.extensions.filters.common.dependency.v3.FilterDependencies.encode_provided:type_name -> envoy.extensions.filters.common.dependency.v3.Dependency
	4, // 5: envoy.extensions.filters.common.dependency.v3.MatchingRequirements.data_input_allow_list:type_name -> envoy.extensions.filters.common.dependency.v3.MatchingRequirements.DataInputAllowList
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_common_dependency_v3_dependency_proto_init() }
func file_envoy_extensions_filters_common_dependency_v3_dependency_proto_init() {
	if File_envoy_extensions_filters_common_dependency_v3_dependency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dependency); i {
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
		file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterDependencies); i {
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
		file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchingRequirements); i {
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
		file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchingRequirements_DataInputAllowList); i {
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
			RawDescriptor: file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_common_dependency_v3_dependency_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_common_dependency_v3_dependency_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_common_dependency_v3_dependency_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_common_dependency_v3_dependency_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_common_dependency_v3_dependency_proto = out.File
	file_envoy_extensions_filters_common_dependency_v3_dependency_proto_rawDesc = nil
	file_envoy_extensions_filters_common_dependency_v3_dependency_proto_goTypes = nil
	file_envoy_extensions_filters_common_dependency_v3_dependency_proto_depIdxs = nil
}
