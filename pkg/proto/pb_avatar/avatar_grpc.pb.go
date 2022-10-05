// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: pb_avatar/avatar.proto

package pb_avatar

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

// AvatarClient is the client API for Avatar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvatarClient interface {
	SetAvatar(ctx context.Context, in *SetAvatarReq, opts ...grpc.CallOption) (*SetAvatarResp, error)
}

type avatarClient struct {
	cc grpc.ClientConnInterface
}

func NewAvatarClient(cc grpc.ClientConnInterface) AvatarClient {
	return &avatarClient{cc}
}

func (c *avatarClient) SetAvatar(ctx context.Context, in *SetAvatarReq, opts ...grpc.CallOption) (*SetAvatarResp, error) {
	out := new(SetAvatarResp)
	err := c.cc.Invoke(ctx, "/Avatar/SetAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvatarServer is the server API for Avatar service.
// All implementations must embed UnimplementedAvatarServer
// for forward compatibility
type AvatarServer interface {
	SetAvatar(context.Context, *SetAvatarReq) (*SetAvatarResp, error)
	mustEmbedUnimplementedAvatarServer()
}

// UnimplementedAvatarServer must be embedded to have forward compatible implementations.
type UnimplementedAvatarServer struct {
}

func (UnimplementedAvatarServer) SetAvatar(context.Context, *SetAvatarReq) (*SetAvatarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAvatar not implemented")
}
func (UnimplementedAvatarServer) mustEmbedUnimplementedAvatarServer() {}

// UnsafeAvatarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvatarServer will
// result in compilation errors.
type UnsafeAvatarServer interface {
	mustEmbedUnimplementedAvatarServer()
}

func RegisterAvatarServer(s grpc.ServiceRegistrar, srv AvatarServer) {
	s.RegisterService(&Avatar_ServiceDesc, srv)
}

func _Avatar_SetAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAvatarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServer).SetAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Avatar/SetAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServer).SetAvatar(ctx, req.(*SetAvatarReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Avatar_ServiceDesc is the grpc.ServiceDesc for Avatar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Avatar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Avatar",
	HandlerType: (*AvatarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAvatar",
			Handler:    _Avatar_SetAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_avatar/avatar.proto",
}
