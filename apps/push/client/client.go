package push_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_push"
)

type PushClient interface {
	PushMessage(req *pb_push.PushMessageReq) (resp *pb_push.PushMessageResp)
}

type pushClient struct {
	opt *xgrpc.ClientDialOption
}

func NewPushClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) PushClient {
	return &pushClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *pushClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *pushClient) PushMessage(req *pb_push.PushMessageReq) (resp *pb_push.PushMessageResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_push.NewPushClient(conn)
	resp, _ = client.PushMessage(context.Background(), req)
	return
}
