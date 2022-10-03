// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: pb_gw/gw.proto

package pb_gw

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

// MessageGatewayClient is the client API for MessageGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageGatewayClient interface {
	OnlinePushMessage(ctx context.Context, in *OnlinePushMessageReq, opts ...grpc.CallOption) (*OnlinePushMessageResp, error)
}

type messageGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageGatewayClient(cc grpc.ClientConnInterface) MessageGatewayClient {
	return &messageGatewayClient{cc}
}

func (c *messageGatewayClient) OnlinePushMessage(ctx context.Context, in *OnlinePushMessageReq, opts ...grpc.CallOption) (*OnlinePushMessageResp, error) {
	out := new(OnlinePushMessageResp)
	err := c.cc.Invoke(ctx, "/pb_gw.MessageGateway/OnlinePushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageGatewayServer is the server API for MessageGateway service.
// All implementations must embed UnimplementedMessageGatewayServer
// for forward compatibility
type MessageGatewayServer interface {
	OnlinePushMessage(context.Context, *OnlinePushMessageReq) (*OnlinePushMessageResp, error)
	mustEmbedUnimplementedMessageGatewayServer()
}

// UnimplementedMessageGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedMessageGatewayServer struct {
}

func (UnimplementedMessageGatewayServer) OnlinePushMessage(context.Context, *OnlinePushMessageReq) (*OnlinePushMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlinePushMessage not implemented")
}
func (UnimplementedMessageGatewayServer) mustEmbedUnimplementedMessageGatewayServer() {}

// UnsafeMessageGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageGatewayServer will
// result in compilation errors.
type UnsafeMessageGatewayServer interface {
	mustEmbedUnimplementedMessageGatewayServer()
}

func RegisterMessageGatewayServer(s grpc.ServiceRegistrar, srv MessageGatewayServer) {
	s.RegisterService(&MessageGateway_ServiceDesc, srv)
}

func _MessageGateway_OnlinePushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlinePushMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGatewayServer).OnlinePushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_gw.MessageGateway/OnlinePushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGatewayServer).OnlinePushMessage(ctx, req.(*OnlinePushMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageGateway_ServiceDesc is the grpc.ServiceDesc for MessageGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_gw.MessageGateway",
	HandlerType: (*MessageGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnlinePushMessage",
			Handler:    _MessageGateway_OnlinePushMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_gw/gw.proto",
}
