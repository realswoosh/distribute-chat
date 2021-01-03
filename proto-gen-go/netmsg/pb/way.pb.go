// https://developers.google.com/protocol-buffers/docs/tutorials

// [START declaration]

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: way.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type ReqWay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ReqWay) Reset() {
	*x = ReqWay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_way_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqWay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqWay) ProtoMessage() {}

func (x *ReqWay) ProtoReflect() protoreflect.Message {
	mi := &file_way_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqWay.ProtoReflect.Descriptor instead.
func (*ReqWay) Descriptor() ([]byte, []int) {
	return file_way_proto_rawDescGZIP(), []int{0}
}

func (x *ReqWay) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqWay) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AckWay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AckWay) Reset() {
	*x = AckWay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_way_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckWay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckWay) ProtoMessage() {}

func (x *AckWay) ProtoReflect() protoreflect.Message {
	mi := &file_way_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckWay.ProtoReflect.Descriptor instead.
func (*AckWay) Descriptor() ([]byte, []int) {
	return file_way_proto_rawDescGZIP(), []int{1}
}

func (x *AckWay) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AckWay) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AckWay) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_way_proto protoreflect.FileDescriptor

var file_way_proto_rawDesc = []byte{
	0x0a, 0x09, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6e, 0x65, 0x74,
	0x6d, 0x73, 0x67, 0x22, 0x32, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x57, 0x61, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x46, 0x0a, 0x06, 0x41, 0x63, 0x6b, 0x57, 0x61,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42,
	0x15, 0x5a, 0x09, 0x6e, 0x65, 0x74, 0x6d, 0x73, 0x67, 0x2f, 0x70, 0x62, 0xaa, 0x02, 0x07, 0x6e,
	0x65, 0x74, 0x2e, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_way_proto_rawDescOnce sync.Once
	file_way_proto_rawDescData = file_way_proto_rawDesc
)

func file_way_proto_rawDescGZIP() []byte {
	file_way_proto_rawDescOnce.Do(func() {
		file_way_proto_rawDescData = protoimpl.X.CompressGZIP(file_way_proto_rawDescData)
	})
	return file_way_proto_rawDescData
}

var file_way_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_way_proto_goTypes = []interface{}{
	(*ReqWay)(nil), // 0: netmsg.ReqWay
	(*AckWay)(nil), // 1: netmsg.AckWay
}
var file_way_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_way_proto_init() }
func file_way_proto_init() {
	if File_way_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_way_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqWay); i {
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
		file_way_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckWay); i {
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
			RawDescriptor: file_way_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_way_proto_goTypes,
		DependencyIndexes: file_way_proto_depIdxs,
		MessageInfos:      file_way_proto_msgTypes,
	}.Build()
	File_way_proto = out.File
	file_way_proto_rawDesc = nil
	file_way_proto_goTypes = nil
	file_way_proto_depIdxs = nil
}