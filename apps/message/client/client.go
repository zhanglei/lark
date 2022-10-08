package msg_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_msg"
)

type MsgClient interface {
	SendChatMessage(req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp)
	MessageOperation(req *pb_msg.MessageOperationReq) (resp *pb_msg.MessageOperationResp)
}

type msgClient struct {
	opt *xgrpc.ClientDialOption
}

func NewMsgClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) MsgClient {
	return &msgClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *msgClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *msgClient) SendChatMessage(req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_msg.NewMessageClient(conn)
	resp, _ = client.SendChatMessage(context.Background(), req)
	return
}

func (c *msgClient) MessageOperation(req *pb_msg.MessageOperationReq) (resp *pb_msg.MessageOperationResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_msg.NewMessageClient(conn)
	resp, _ = client.MessageOperation(context.Background(), req)
	return
}
