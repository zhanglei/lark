package chat_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_chat"
)

type ChatClient interface {
	NewGroupChat(req *pb_chat.NewGroupChatReq) (resp *pb_chat.NewGroupChatResp)
	SetGroupChat(req *pb_chat.SetGroupChatReq) (resp *pb_chat.SetGroupChatResp)
}

type chatClient struct {
	opt *xgrpc.ClientDialOption
}

func NewChatClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ChatClient {
	return &chatClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *chatClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *chatClient) NewGroupChat(req *pb_chat.NewGroupChatReq) (resp *pb_chat.NewGroupChatResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	resp, _ = client.NewGroupChat(context.Background(), req)
	return
}

func (c *chatClient) SetGroupChat(req *pb_chat.SetGroupChatReq) (resp *pb_chat.SetGroupChatResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	resp, _ = client.SetGroupChat(context.Background(), req)
	return
}
