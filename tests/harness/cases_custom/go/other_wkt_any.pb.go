// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: other_wkt_any.proto

package cases_custom

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AnyNone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *any.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyNone) Reset() {
	*x = AnyNone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_wkt_any_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyNone) ProtoMessage() {}

func (x *AnyNone) ProtoReflect() protoreflect.Message {
	mi := &file_other_wkt_any_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyNone.ProtoReflect.Descriptor instead.
func (*AnyNone) Descriptor() ([]byte, []int) {
	return file_other_wkt_any_proto_rawDescGZIP(), []int{0}
}

func (x *AnyNone) GetVal() *any.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

type AnyRequired struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *any.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyRequired) Reset() {
	*x = AnyRequired{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_wkt_any_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyRequired) ProtoMessage() {}

func (x *AnyRequired) ProtoReflect() protoreflect.Message {
	mi := &file_other_wkt_any_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyRequired.ProtoReflect.Descriptor instead.
func (*AnyRequired) Descriptor() ([]byte, []int) {
	return file_other_wkt_any_proto_rawDescGZIP(), []int{1}
}

func (x *AnyRequired) GetVal() *any.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

type AnyIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *any.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyIn) Reset() {
	*x = AnyIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_wkt_any_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyIn) ProtoMessage() {}

func (x *AnyIn) ProtoReflect() protoreflect.Message {
	mi := &file_other_wkt_any_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyIn.ProtoReflect.Descriptor instead.
func (*AnyIn) Descriptor() ([]byte, []int) {
	return file_other_wkt_any_proto_rawDescGZIP(), []int{2}
}

func (x *AnyIn) GetVal() *any.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

type AnyNotIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *any.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyNotIn) Reset() {
	*x = AnyNotIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_wkt_any_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyNotIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyNotIn) ProtoMessage() {}

func (x *AnyNotIn) ProtoReflect() protoreflect.Message {
	mi := &file_other_wkt_any_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyNotIn.ProtoReflect.Descriptor instead.
func (*AnyNotIn) Descriptor() ([]byte, []int) {
	return file_other_wkt_any_proto_rawDescGZIP(), []int{3}
}

func (x *AnyNotIn) GetVal() *any.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_other_wkt_any_proto protoreflect.FileDescriptor

var file_other_wkt_any_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x77, 0x6b, 0x74, 0x5f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x07, 0x41, 0x6e, 0x79, 0x4e, 0x6f, 0x6e, 0x65,
	0x12, 0x26, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x0b, 0x41, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x65, 0x0a, 0x05, 0x41, 0x6e, 0x79,
	0x49, 0x6e, 0x12, 0x5c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x34, 0xfa, 0x42, 0x31, 0xa2, 0x01, 0x2e, 0x12, 0x2c, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x69, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x12, 0x5d, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42,
	0x35, 0xfa, 0x42, 0x32, 0xa2, 0x01, 0x2f, 0x1a, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0x0e, 0x5a, 0x0c, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_other_wkt_any_proto_rawDescOnce sync.Once
	file_other_wkt_any_proto_rawDescData = file_other_wkt_any_proto_rawDesc
)

func file_other_wkt_any_proto_rawDescGZIP() []byte {
	file_other_wkt_any_proto_rawDescOnce.Do(func() {
		file_other_wkt_any_proto_rawDescData = protoimpl.X.CompressGZIP(file_other_wkt_any_proto_rawDescData)
	})
	return file_other_wkt_any_proto_rawDescData
}

var file_other_wkt_any_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_other_wkt_any_proto_goTypes = []interface{}{
	(*AnyNone)(nil),     // 0: tests.harness.cases_custom.AnyNone
	(*AnyRequired)(nil), // 1: tests.harness.cases_custom.AnyRequired
	(*AnyIn)(nil),       // 2: tests.harness.cases_custom.AnyIn
	(*AnyNotIn)(nil),    // 3: tests.harness.cases_custom.AnyNotIn
	(*any.Any)(nil),     // 4: google.protobuf.Any
}
var file_other_wkt_any_proto_depIdxs = []int32{
	4, // 0: tests.harness.cases_custom.AnyNone.val:type_name -> google.protobuf.Any
	4, // 1: tests.harness.cases_custom.AnyRequired.val:type_name -> google.protobuf.Any
	4, // 2: tests.harness.cases_custom.AnyIn.val:type_name -> google.protobuf.Any
	4, // 3: tests.harness.cases_custom.AnyNotIn.val:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_other_wkt_any_proto_init() }
func file_other_wkt_any_proto_init() {
	if File_other_wkt_any_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_other_wkt_any_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyNone); i {
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
		file_other_wkt_any_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyRequired); i {
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
		file_other_wkt_any_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyIn); i {
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
		file_other_wkt_any_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyNotIn); i {
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
			RawDescriptor: file_other_wkt_any_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_other_wkt_any_proto_goTypes,
		DependencyIndexes: file_other_wkt_any_proto_depIdxs,
		MessageInfos:      file_other_wkt_any_proto_msgTypes,
	}.Build()
	File_other_wkt_any_proto = out.File
	file_other_wkt_any_proto_rawDesc = nil
	file_other_wkt_any_proto_goTypes = nil
	file_other_wkt_any_proto_depIdxs = nil
}
