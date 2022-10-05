package chat_invite_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_invite"
)

type ChatInviteClient interface {
	NewChatInvite(req *pb_invite.NewChatInviteReq) (resp *pb_invite.NewChatInviteResp)
	ChatInviteHandle(req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp)
	ChatInviteList(req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp)
}

type chatInviteClient struct {
	opt *xgrpc.ClientDialOption
}

func NewChatInviteClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ChatInviteClient {
	return &chatInviteClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *chatInviteClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *chatInviteClient) NewChatInvite(req *pb_invite.NewChatInviteReq) (resp *pb_invite.NewChatInviteResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_invite.NewInviteClient(conn)
	resp, _ = client.NewChatInvite(context.Background(), req)
	return
}

func (c *chatInviteClient) ChatInviteHandle(req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_invite.NewInviteClient(conn)
	resp, _ = client.ChatInviteHandle(context.Background(), req)
	return
}

func (c *chatInviteClient) ChatInviteList(req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_invite.NewInviteClient(conn)
	resp, _ = client.ChatInviteList(context.Background(), req)
	return
}
