package gw_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_gw"
)

type MessageGatewayClient interface {
	OnlinePushMessage(req *pb_gw.OnlinePushMessageReq) (resp *pb_gw.OnlinePushMessageResp)
}

type messageGatewayClient struct {
	opt *xgrpc.ClientDialOption
}

func NewMsgGwClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) MessageGatewayClient {
	return &messageGatewayClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *messageGatewayClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *messageGatewayClient) OnlinePushMessage(req *pb_gw.OnlinePushMessageReq) (resp *pb_gw.OnlinePushMessageResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_gw.NewMessageGatewayClient(conn)
	resp, _ = client.OnlinePushMessage(context.Background(), req)
	return
}
