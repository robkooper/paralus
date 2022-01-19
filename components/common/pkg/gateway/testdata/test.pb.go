// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: pkg/gateway/testdata/test.proto

package testdata

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestYAML struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Time   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Labels map[string]string      `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestYAML) Reset() {
	*x = TestYAML{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_gateway_testdata_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestYAML) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestYAML) ProtoMessage() {}

func (x *TestYAML) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_gateway_testdata_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestYAML.ProtoReflect.Descriptor instead.
func (*TestYAML) Descriptor() ([]byte, []int) {
	return file_pkg_gateway_testdata_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestYAML) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestYAML) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *TestYAML) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type TestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestMetadata) Reset() {
	*x = TestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_gateway_testdata_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMetadata) ProtoMessage() {}

func (x *TestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_gateway_testdata_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMetadata.ProtoReflect.Descriptor instead.
func (*TestMetadata) Descriptor() ([]byte, []int) {
	return file_pkg_gateway_testdata_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TestObjectSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestObjectSpec) Reset() {
	*x = TestObjectSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_gateway_testdata_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestObjectSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestObjectSpec) ProtoMessage() {}

func (x *TestObjectSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_gateway_testdata_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestObjectSpec.ProtoReflect.Descriptor instead.
func (*TestObjectSpec) Descriptor() ([]byte, []int) {
	return file_pkg_gateway_testdata_test_proto_rawDescGZIP(), []int{2}
}

type TestObjectStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestObjectStatus) Reset() {
	*x = TestObjectStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_gateway_testdata_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestObjectStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestObjectStatus) ProtoMessage() {}

func (x *TestObjectStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_gateway_testdata_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestObjectStatus.ProtoReflect.Descriptor instead.
func (*TestObjectStatus) Descriptor() ([]byte, []int) {
	return file_pkg_gateway_testdata_test_proto_rawDescGZIP(), []int{3}
}

type TestObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *TestMetadata     `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *TestObjectSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *TestObjectStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TestObject) Reset() {
	*x = TestObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_gateway_testdata_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestObject) ProtoMessage() {}

func (x *TestObject) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_gateway_testdata_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestObject.ProtoReflect.Descriptor instead.
func (*TestObject) Descriptor() ([]byte, []int) {
	return file_pkg_gateway_testdata_test_proto_rawDescGZIP(), []int{4}
}

func (x *TestObject) GetMetadata() *TestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TestObject) GetSpec() *TestObjectSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TestObject) GetStatus() *TestObjectStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_pkg_gateway_testdata_test_proto protoreflect.FileDescriptor

var file_pkg_gateway_testdata_test_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcc, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x59, 0x41, 0x4d, 0x4c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0x92, 0x41, 0x02,
	0x40, 0x01, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x59, 0x41, 0x4d, 0x4c,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x22, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x12, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x0a, 0x54, 0x65,
	0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x47, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x13, 0xc8, 0xde, 0x1f, 0x00, 0xd0, 0xde, 0x1f, 0x01, 0xea, 0xde, 0x1f, 0x07,
	0x2c, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x36, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0xd0,
	0xde, 0x1f, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0xd0, 0xde, 0x1f, 0x01, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x70, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x68, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x14, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x76, 0x33, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0xb9, 0x01, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0xca, 0x02, 0x08, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0xe2,
	0x02, 0x14, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0xc8, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01, 0xc0, 0xe3, 0x1e,
	0x01, 0xc8, 0xe3, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_gateway_testdata_test_proto_rawDescOnce sync.Once
	file_pkg_gateway_testdata_test_proto_rawDescData = file_pkg_gateway_testdata_test_proto_rawDesc
)

func file_pkg_gateway_testdata_test_proto_rawDescGZIP() []byte {
	file_pkg_gateway_testdata_test_proto_rawDescOnce.Do(func() {
		file_pkg_gateway_testdata_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_gateway_testdata_test_proto_rawDescData)
	})
	return file_pkg_gateway_testdata_test_proto_rawDescData
}

var file_pkg_gateway_testdata_test_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_gateway_testdata_test_proto_goTypes = []interface{}{
	(*TestYAML)(nil),              // 0: testdata.TestYAML
	(*TestMetadata)(nil),          // 1: testdata.TestMetadata
	(*TestObjectSpec)(nil),        // 2: testdata.TestObjectSpec
	(*TestObjectStatus)(nil),      // 3: testdata.TestObjectStatus
	(*TestObject)(nil),            // 4: testdata.TestObject
	nil,                           // 5: testdata.TestYAML.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_pkg_gateway_testdata_test_proto_depIdxs = []int32{
	6, // 0: testdata.TestYAML.time:type_name -> google.protobuf.Timestamp
	5, // 1: testdata.TestYAML.labels:type_name -> testdata.TestYAML.LabelsEntry
	1, // 2: testdata.TestObject.metadata:type_name -> testdata.TestMetadata
	2, // 3: testdata.TestObject.spec:type_name -> testdata.TestObjectSpec
	3, // 4: testdata.TestObject.status:type_name -> testdata.TestObjectStatus
	4, // 5: testdata.Test.Get:input_type -> testdata.TestObject
	4, // 6: testdata.Test.Get:output_type -> testdata.TestObject
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_gateway_testdata_test_proto_init() }
func file_pkg_gateway_testdata_test_proto_init() {
	if File_pkg_gateway_testdata_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_gateway_testdata_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestYAML); i {
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
		file_pkg_gateway_testdata_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMetadata); i {
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
		file_pkg_gateway_testdata_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestObjectSpec); i {
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
		file_pkg_gateway_testdata_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestObjectStatus); i {
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
		file_pkg_gateway_testdata_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestObject); i {
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
			RawDescriptor: file_pkg_gateway_testdata_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_gateway_testdata_test_proto_goTypes,
		DependencyIndexes: file_pkg_gateway_testdata_test_proto_depIdxs,
		MessageInfos:      file_pkg_gateway_testdata_test_proto_msgTypes,
	}.Build()
	File_pkg_gateway_testdata_test_proto = out.File
	file_pkg_gateway_testdata_test_proto_rawDesc = nil
	file_pkg_gateway_testdata_test_proto_goTypes = nil
	file_pkg_gateway_testdata_test_proto_depIdxs = nil
}
