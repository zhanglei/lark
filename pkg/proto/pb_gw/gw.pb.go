// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pb_gw/gw.proto

package pb_gw

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "lark/pkg/proto/pb_chat_member"
	pb_enum "lark/pkg/proto/pb_enum"
	pb_msg "lark/pkg/proto/pb_msg"
	pb_obj "lark/pkg/proto/pb_obj"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OnlinePushMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic    pb_enum.TOPIC          `protobuf:"varint,1,opt,name=topic,proto3,enum=pb_enum.TOPIC" json:"topic,omitempty"`
	SubTopic pb_enum.SUB_TOPIC      `protobuf:"varint,2,opt,name=sub_topic,json=subTopic,proto3,enum=pb_enum.SUB_TOPIC" json:"sub_topic,omitempty"`
	Members  []*pb_obj.Int64Array   `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Msg      *pb_msg.SrvChatMessage `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *OnlinePushMessageReq) Reset() {
	*x = OnlinePushMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gw_gw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePushMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMessageReq) ProtoMessage() {}

func (x *OnlinePushMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gw_gw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMessageReq.ProtoReflect.Descriptor instead.
func (*OnlinePushMessageReq) Descriptor() ([]byte, []int) {
	return file_pb_gw_gw_proto_rawDescGZIP(), []int{0}
}

func (x *OnlinePushMessageReq) GetTopic() pb_enum.TOPIC {
	if x != nil {
		return x.Topic
	}
	return pb_enum.TOPIC(0)
}

func (x *OnlinePushMessageReq) GetSubTopic() pb_enum.SUB_TOPIC {
	if x != nil {
		return x.SubTopic
	}
	return pb_enum.SUB_TOPIC(0)
}

func (x *OnlinePushMessageReq) GetMembers() []*pb_obj.Int64Array {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *OnlinePushMessageReq) GetMsg() *pb_msg.SrvChatMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

type OnlinePushMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32               `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	PushResps []*PlatformPushResp `protobuf:"bytes,3,rep,name=push_resps,json=pushResps,proto3" json:"push_resps,omitempty"`
}

func (x *OnlinePushMessageResp) Reset() {
	*x = OnlinePushMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gw_gw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePushMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMessageResp) ProtoMessage() {}

func (x *OnlinePushMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gw_gw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMessageResp.ProtoReflect.Descriptor instead.
func (*OnlinePushMessageResp) Descriptor() ([]byte, []int) {
	return file_pb_gw_gw_proto_rawDescGZIP(), []int{1}
}

func (x *OnlinePushMessageResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OnlinePushMessageResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OnlinePushMessageResp) GetPushResps() []*PlatformPushResp {
	if x != nil {
		return x.PushResps
	}
	return nil
}

type PlatformPushResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string                `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Platform pb_enum.PLATFORM_TYPE `protobuf:"varint,3,opt,name=platform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"platform,omitempty"`
}

func (x *PlatformPushResp) Reset() {
	*x = PlatformPushResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gw_gw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformPushResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformPushResp) ProtoMessage() {}

func (x *PlatformPushResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gw_gw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformPushResp.ProtoReflect.Descriptor instead.
func (*PlatformPushResp) Descriptor() ([]byte, []int) {
	return file_pb_gw_gw_proto_rawDescGZIP(), []int{2}
}

func (x *PlatformPushResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PlatformPushResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PlatformPushResp) GetPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.Platform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

type OnlinePushMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64                 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ServerId int32                 `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Platform pb_enum.PLATFORM_TYPE `protobuf:"varint,3,opt,name=platform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"platform,omitempty"`
	Mute     pb_enum.MUTE_TYPE     `protobuf:"varint,4,opt,name=mute,proto3,enum=pb_enum.MUTE_TYPE" json:"mute,omitempty"`
}

func (x *OnlinePushMember) Reset() {
	*x = OnlinePushMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gw_gw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePushMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMember) ProtoMessage() {}

func (x *OnlinePushMember) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gw_gw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMember.ProtoReflect.Descriptor instead.
func (*OnlinePushMember) Descriptor() ([]byte, []int) {
	return file_pb_gw_gw_proto_rawDescGZIP(), []int{3}
}

func (x *OnlinePushMember) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *OnlinePushMember) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *OnlinePushMember) GetPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.Platform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *OnlinePushMember) GetMute() pb_enum.MUTE_TYPE {
	if x != nil {
		return x.Mute
	}
	return pb_enum.MUTE_TYPE(0)
}

var File_pb_gw_gw_proto protoreflect.FileDescriptor

var file_pb_gw_gw_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x5f, 0x67, 0x77, 0x2f, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x62, 0x5f, 0x67, 0x77, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x62, 0x5f,
	0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70,
	0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x70, 0x62, 0x5f, 0x6f, 0x62, 0x6a, 0x2f, 0x6f, 0x62, 0x6a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x14, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x2f, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x55,
	0x42, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x6f, 0x62, 0x6a, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x28, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x72, 0x76, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x75, 0x0a, 0x15, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x62, 0x5f, 0x67, 0x77, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x52, 0x09, 0x70, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x73,
	0x22, 0x6c, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x9d,
	0x01, 0x0a, 0x10, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x26, 0x0a, 0x04, 0x6d, 0x75, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x4d,
	0x55, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x6d, 0x75, 0x74, 0x65, 0x32, 0x60,
	0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x12, 0x4e, 0x0a, 0x11, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x77, 0x2e, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x77, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x67, 0x77, 0x3b, 0x70, 0x62, 0x5f, 0x67,
	0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_gw_gw_proto_rawDescOnce sync.Once
	file_pb_gw_gw_proto_rawDescData = file_pb_gw_gw_proto_rawDesc
)

func file_pb_gw_gw_proto_rawDescGZIP() []byte {
	file_pb_gw_gw_proto_rawDescOnce.Do(func() {
		file_pb_gw_gw_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_gw_gw_proto_rawDescData)
	})
	return file_pb_gw_gw_proto_rawDescData
}

var file_pb_gw_gw_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_gw_gw_proto_goTypes = []interface{}{
	(*OnlinePushMessageReq)(nil),  // 0: pb_gw.OnlinePushMessageReq
	(*OnlinePushMessageResp)(nil), // 1: pb_gw.OnlinePushMessageResp
	(*PlatformPushResp)(nil),      // 2: pb_gw.PlatformPushResp
	(*OnlinePushMember)(nil),      // 3: pb_gw.OnlinePushMember
	(pb_enum.TOPIC)(0),            // 4: pb_enum.TOPIC
	(pb_enum.SUB_TOPIC)(0),        // 5: pb_enum.SUB_TOPIC
	(*pb_obj.Int64Array)(nil),     // 6: pb_obj.Int64Array
	(*pb_msg.SrvChatMessage)(nil), // 7: pb_msg.SrvChatMessage
	(pb_enum.PLATFORM_TYPE)(0),    // 8: pb_enum.PLATFORM_TYPE
	(pb_enum.MUTE_TYPE)(0),        // 9: pb_enum.MUTE_TYPE
}
var file_pb_gw_gw_proto_depIdxs = []int32{
	4, // 0: pb_gw.OnlinePushMessageReq.topic:type_name -> pb_enum.TOPIC
	5, // 1: pb_gw.OnlinePushMessageReq.sub_topic:type_name -> pb_enum.SUB_TOPIC
	6, // 2: pb_gw.OnlinePushMessageReq.members:type_name -> pb_obj.Int64Array
	7, // 3: pb_gw.OnlinePushMessageReq.msg:type_name -> pb_msg.SrvChatMessage
	2, // 4: pb_gw.OnlinePushMessageResp.push_resps:type_name -> pb_gw.PlatformPushResp
	8, // 5: pb_gw.PlatformPushResp.platform:type_name -> pb_enum.PLATFORM_TYPE
	8, // 6: pb_gw.OnlinePushMember.platform:type_name -> pb_enum.PLATFORM_TYPE
	9, // 7: pb_gw.OnlinePushMember.mute:type_name -> pb_enum.MUTE_TYPE
	0, // 8: pb_gw.MessageGateway.OnlinePushMessage:input_type -> pb_gw.OnlinePushMessageReq
	1, // 9: pb_gw.MessageGateway.OnlinePushMessage:output_type -> pb_gw.OnlinePushMessageResp
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pb_gw_gw_proto_init() }
func file_pb_gw_gw_proto_init() {
	if File_pb_gw_gw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_gw_gw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePushMessageReq); i {
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
		file_pb_gw_gw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePushMessageResp); i {
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
		file_pb_gw_gw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformPushResp); i {
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
		file_pb_gw_gw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePushMember); i {
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
			RawDescriptor: file_pb_gw_gw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_gw_gw_proto_goTypes,
		DependencyIndexes: file_pb_gw_gw_proto_depIdxs,
		MessageInfos:      file_pb_gw_gw_proto_msgTypes,
	}.Build()
	File_pb_gw_gw_proto = out.File
	file_pb_gw_gw_proto_rawDesc = nil
	file_pb_gw_gw_proto_goTypes = nil
	file_pb_gw_gw_proto_depIdxs = nil
}
