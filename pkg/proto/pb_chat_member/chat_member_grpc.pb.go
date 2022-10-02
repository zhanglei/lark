// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb_chat_member/chat_member.proto

package pb_chat_member

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatMemberClient is the client API for ChatMember service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatMemberClient interface {
	GetChatMemberUidList(ctx context.Context, in *GetChatMemberUidListReq, opts ...grpc.CallOption) (*GetChatMemberUidListResp, error)
	GetChatMemberPushConfig(ctx context.Context, in *GetChatMemberPushConfigReq, opts ...grpc.CallOption) (*GetChatMemberPushConfigResp, error)
	GetChatMemberPushConfigList(ctx context.Context, in *GetChatMemberPushConfigListReq, opts ...grpc.CallOption) (*GetChatMemberPushConfigListResp, error)
	GetChatMemberSetting(ctx context.Context, in *GetChatMemberSettingReq, opts ...grpc.CallOption) (*GetChatMemberSettingResp, error)
	GetChatMemberInfo(ctx context.Context, in *GetChatMemberInfoReq, opts ...grpc.CallOption) (*GetChatMemberInfoResp, error)
	ChatMemberVerify(ctx context.Context, in *ChatMemberVerifyReq, opts ...grpc.CallOption) (*ChatMemberVerifyResp, error)
	ChatMemberOnline(ctx context.Context, in *ChatMemberOnlineReq, opts ...grpc.CallOption) (*ChatMemberOnlineResp, error)
}

type chatMemberClient struct {
	cc grpc.ClientConnInterface
}

func NewChatMemberClient(cc grpc.ClientConnInterface) ChatMemberClient {
	return &chatMemberClient{cc}
}

func (c *chatMemberClient) GetChatMemberUidList(ctx context.Context, in *GetChatMemberUidListReq, opts ...grpc.CallOption) (*GetChatMemberUidListResp, error) {
	out := new(GetChatMemberUidListResp)
	err := c.cc.Invoke(ctx, "/pb_chat_member.ChatMember/GetChatMemberUidList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetChatMemberPushConfig(ctx context.Context, in *GetChatMemberPushConfigReq, opts ...grpc.CallOption) (*GetChatMemberPushConfigResp, error) {
	out := new(GetChatMemberPushConfigResp)
	err := c.cc.Invoke(ctx, "/pb_chat_member.ChatMember/GetChatMemberPushConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetChatMemberPushConfigList(ctx context.Context, in *GetChatMemberPushConfigListReq, opts ...grpc.CallOption) (*GetChatMemberPushConfigListResp, error) {
	out := new(GetChatMemberPushConfigListResp)
	err := c.cc.Invoke(ctx, "/pb_chat_member.ChatMember/GetChatMemberPushConfigList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetChatMemberSetting(ctx context.Context, in *GetChatMemberSettingReq, opts ...grpc.CallOption) (*GetChatMemberSettingResp, error) {
	out := new(GetChatMemberSettingResp)
	err := c.cc.Invoke(ctx, "/pb_chat_member.ChatMember/GetChatMemberSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetChatMemberInfo(ctx context.Context, in *GetChatMemberInfoReq, opts ...grpc.CallOption) (*GetChatMemberInfoResp, error) {
	out := new(GetChatMemberInfoResp)
	err := c.cc.Invoke(ctx, "/pb_chat_member.ChatMember/GetChatMemberInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) ChatMemberVerify(ctx context.Context, in *ChatMemberVerifyReq, opts ...grpc.CallOption) (*ChatMemberVerifyResp, error) {
	out := new(ChatMemberVerifyResp)
	err := c.cc.Invoke(ctx, "/pb_chat_member.ChatMember/ChatMemberVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) ChatMemberOnline(ctx context.Context, in *ChatMemberOnlineReq, opts ...grpc.CallOption) (*ChatMemberOnlineResp, error) {
	out := new(ChatMemberOnlineResp)
	err := c.cc.Invoke(ctx, "/pb_chat_member.ChatMember/ChatMemberOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatMemberServer is the server API for ChatMember service.
// All implementations must embed UnimplementedChatMemberServer
// for forward compatibility
type ChatMemberServer interface {
	GetChatMemberUidList(context.Context, *GetChatMemberUidListReq) (*GetChatMemberUidListResp, error)
	GetChatMemberPushConfig(context.Context, *GetChatMemberPushConfigReq) (*GetChatMemberPushConfigResp, error)
	GetChatMemberPushConfigList(context.Context, *GetChatMemberPushConfigListReq) (*GetChatMemberPushConfigListResp, error)
	GetChatMemberSetting(context.Context, *GetChatMemberSettingReq) (*GetChatMemberSettingResp, error)
	GetChatMemberInfo(context.Context, *GetChatMemberInfoReq) (*GetChatMemberInfoResp, error)
	ChatMemberVerify(context.Context, *ChatMemberVerifyReq) (*ChatMemberVerifyResp, error)
	ChatMemberOnline(context.Context, *ChatMemberOnlineReq) (*ChatMemberOnlineResp, error)
	mustEmbedUnimplementedChatMemberServer()
}

// UnimplementedChatMemberServer must be embedded to have forward compatible implementations.
type UnimplementedChatMemberServer struct {
}

func (UnimplementedChatMemberServer) GetChatMemberUidList(context.Context, *GetChatMemberUidListReq) (*GetChatMemberUidListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMemberUidList not implemented")
}
func (UnimplementedChatMemberServer) GetChatMemberPushConfig(context.Context, *GetChatMemberPushConfigReq) (*GetChatMemberPushConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMemberPushConfig not implemented")
}
func (UnimplementedChatMemberServer) GetChatMemberPushConfigList(context.Context, *GetChatMemberPushConfigListReq) (*GetChatMemberPushConfigListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMemberPushConfigList not implemented")
}
func (UnimplementedChatMemberServer) GetChatMemberSetting(context.Context, *GetChatMemberSettingReq) (*GetChatMemberSettingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMemberSetting not implemented")
}
func (UnimplementedChatMemberServer) GetChatMemberInfo(context.Context, *GetChatMemberInfoReq) (*GetChatMemberInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMemberInfo not implemented")
}
func (UnimplementedChatMemberServer) ChatMemberVerify(context.Context, *ChatMemberVerifyReq) (*ChatMemberVerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatMemberVerify not implemented")
}
func (UnimplementedChatMemberServer) ChatMemberOnline(context.Context, *ChatMemberOnlineReq) (*ChatMemberOnlineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatMemberOnline not implemented")
}
func (UnimplementedChatMemberServer) mustEmbedUnimplementedChatMemberServer() {}

// UnsafeChatMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatMemberServer will
// result in compilation errors.
type UnsafeChatMemberServer interface {
	mustEmbedUnimplementedChatMemberServer()
}

func RegisterChatMemberServer(s grpc.ServiceRegistrar, srv ChatMemberServer) {
	s.RegisterService(&ChatMember_ServiceDesc, srv)
}

func _ChatMember_GetChatMemberUidList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMemberUidListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetChatMemberUidList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_member.ChatMember/GetChatMemberUidList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetChatMemberUidList(ctx, req.(*GetChatMemberUidListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetChatMemberPushConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMemberPushConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetChatMemberPushConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_member.ChatMember/GetChatMemberPushConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetChatMemberPushConfig(ctx, req.(*GetChatMemberPushConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetChatMemberPushConfigList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMemberPushConfigListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetChatMemberPushConfigList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_member.ChatMember/GetChatMemberPushConfigList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetChatMemberPushConfigList(ctx, req.(*GetChatMemberPushConfigListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetChatMemberSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMemberSettingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetChatMemberSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_member.ChatMember/GetChatMemberSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetChatMemberSetting(ctx, req.(*GetChatMemberSettingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetChatMemberInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMemberInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetChatMemberInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_member.ChatMember/GetChatMemberInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetChatMemberInfo(ctx, req.(*GetChatMemberInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_ChatMemberVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).ChatMemberVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_member.ChatMember/ChatMemberVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).ChatMemberVerify(ctx, req.(*ChatMemberVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_ChatMemberOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).ChatMemberOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_member.ChatMember/ChatMemberOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).ChatMemberOnline(ctx, req.(*ChatMemberOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatMember_ServiceDesc is the grpc.ServiceDesc for ChatMember service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatMember_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_chat_member.ChatMember",
	HandlerType: (*ChatMemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChatMemberUidList",
			Handler:    _ChatMember_GetChatMemberUidList_Handler,
		},
		{
			MethodName: "GetChatMemberPushConfig",
			Handler:    _ChatMember_GetChatMemberPushConfig_Handler,
		},
		{
			MethodName: "GetChatMemberPushConfigList",
			Handler:    _ChatMember_GetChatMemberPushConfigList_Handler,
		},
		{
			MethodName: "GetChatMemberSetting",
			Handler:    _ChatMember_GetChatMemberSetting_Handler,
		},
		{
			MethodName: "GetChatMemberInfo",
			Handler:    _ChatMember_GetChatMemberInfo_Handler,
		},
		{
			MethodName: "ChatMemberVerify",
			Handler:    _ChatMember_ChatMemberVerify_Handler,
		},
		{
			MethodName: "ChatMemberOnline",
			Handler:    _ChatMember_ChatMemberOnline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_chat_member/chat_member.proto",
}