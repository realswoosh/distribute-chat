// https://developers.google.com/protocol-buffers/docs/tutorials

// [START declaration]

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: room.proto

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomIdx          int32  `protobuf:"varint,1,opt,name=roomIdx,proto3" json:"roomIdx,omitempty"`
	Title            string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ParticipateCount int32  `protobuf:"varint,3,opt,name=participateCount,proto3" json:"participateCount,omitempty"`
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *RoomInfo) GetRoomIdx() int32 {
	if x != nil {
		return x.RoomIdx
	}
	return 0
}

func (x *RoomInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RoomInfo) GetParticipateCount() int32 {
	if x != nil {
		return x.ParticipateCount
	}
	return 0
}

type AckRoomList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomList []*RoomInfo `protobuf:"bytes,1,rep,name=roomList,proto3" json:"roomList,omitempty"`
}

func (x *AckRoomList) Reset() {
	*x = AckRoomList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckRoomList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckRoomList) ProtoMessage() {}

func (x *AckRoomList) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckRoomList.ProtoReflect.Descriptor instead.
func (*AckRoomList) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *AckRoomList) GetRoomList() []*RoomInfo {
	if x != nil {
		return x.RoomList
	}
	return nil
}

type ReqRoomJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomIdx int32  `protobuf:"varint,1,opt,name=roomIdx,proto3" json:"roomIdx,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ReqRoomJoin) Reset() {
	*x = ReqRoomJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqRoomJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRoomJoin) ProtoMessage() {}

func (x *ReqRoomJoin) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRoomJoin.ProtoReflect.Descriptor instead.
func (*ReqRoomJoin) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{3}
}

func (x *ReqRoomJoin) GetRoomIdx() int32 {
	if x != nil {
		return x.RoomIdx
	}
	return 0
}

func (x *ReqRoomJoin) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type AckRoomJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomIdx int32       `protobuf:"varint,1,opt,name=roomIdx,proto3" json:"roomIdx,omitempty"`
	Title   string      `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Users   []*UserInfo `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *AckRoomJoin) Reset() {
	*x = AckRoomJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckRoomJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckRoomJoin) ProtoMessage() {}

func (x *AckRoomJoin) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckRoomJoin.ProtoReflect.Descriptor instead.
func (*AckRoomJoin) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{4}
}

func (x *AckRoomJoin) GetRoomIdx() int32 {
	if x != nil {
		return x.RoomIdx
	}
	return 0
}

func (x *AckRoomJoin) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AckRoomJoin) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6e, 0x65,
	0x74, 0x6d, 0x73, 0x67, 0x22, 0x32, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x3b, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3d, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x65, 0x0a, 0x0b,
	0x41, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74,
	0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x42, 0x15, 0x5a, 0x09, 0x6e, 0x65, 0x74, 0x6d, 0x73, 0x67, 0x2f, 0x70, 0x62,
	0xaa, 0x02, 0x07, 0x6e, 0x65, 0x74, 0x2e, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_room_proto_goTypes = []interface{}{
	(*UserInfo)(nil),    // 0: netmsg.UserInfo
	(*RoomInfo)(nil),    // 1: netmsg.RoomInfo
	(*AckRoomList)(nil), // 2: netmsg.AckRoomList
	(*ReqRoomJoin)(nil), // 3: netmsg.ReqRoomJoin
	(*AckRoomJoin)(nil), // 4: netmsg.AckRoomJoin
}
var file_room_proto_depIdxs = []int32{
	1, // 0: netmsg.AckRoomList.roomList:type_name -> netmsg.RoomInfo
	0, // 1: netmsg.AckRoomJoin.users:type_name -> netmsg.UserInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomInfo); i {
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
		file_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckRoomList); i {
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
		file_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqRoomJoin); i {
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
		file_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckRoomJoin); i {
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
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}