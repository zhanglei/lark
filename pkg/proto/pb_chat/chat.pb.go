// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pb_chat/chat.proto

package pb_chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "lark/pkg/proto/pb_enum"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *NewChatReq) Reset() {
	*x = NewChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewChatReq) ProtoMessage() {}

func (x *NewChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewChatReq.ProtoReflect.Descriptor instead.
func (*NewChatReq) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *NewChatReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type NewChatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NewChatResp) Reset() {
	*x = NewChatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewChatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewChatResp) ProtoMessage() {}

func (x *NewChatResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewChatResp.ProtoReflect.Descriptor instead.
func (*NewChatResp) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *NewChatResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *NewChatResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pb_chat_chat_proto protoreflect.FileDescriptor

var file_pb_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x12, 0x70,
	0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x33, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x3c, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x12, 0x13, 0x2e,
	0x70, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4e, 0x65, 0x77,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x62,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x3b, 0x70, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_chat_chat_proto_rawDescOnce sync.Once
	file_pb_chat_chat_proto_rawDescData = file_pb_chat_chat_proto_rawDesc
)

func file_pb_chat_chat_proto_rawDescGZIP() []byte {
	file_pb_chat_chat_proto_rawDescOnce.Do(func() {
		file_pb_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_chat_chat_proto_rawDescData)
	})
	return file_pb_chat_chat_proto_rawDescData
}

var file_pb_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_chat_chat_proto_goTypes = []interface{}{
	(*NewChatReq)(nil),  // 0: pb_chat.NewChatReq
	(*NewChatResp)(nil), // 1: pb_chat.NewChatResp
}
var file_pb_chat_chat_proto_depIdxs = []int32{
	0, // 0: pb_chat.Chat.NewChat:input_type -> pb_chat.NewChatReq
	1, // 1: pb_chat.Chat.NewChat:output_type -> pb_chat.NewChatResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_chat_chat_proto_init() }
func file_pb_chat_chat_proto_init() {
	if File_pb_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewChatReq); i {
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
		file_pb_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewChatResp); i {
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
			RawDescriptor: file_pb_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_chat_chat_proto_goTypes,
		DependencyIndexes: file_pb_chat_chat_proto_depIdxs,
		MessageInfos:      file_pb_chat_chat_proto_msgTypes,
	}.Build()
	File_pb_chat_chat_proto = out.File
	file_pb_chat_chat_proto_rawDesc = nil
	file_pb_chat_chat_proto_goTypes = nil
	file_pb_chat_chat_proto_depIdxs = nil
}
