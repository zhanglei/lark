package link_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_link"
)

type LinkClient interface {
	UserOnline(req *pb_link.UserOnlineReq) (resp *pb_link.UserOnlineResp)
}

type linkClient struct {
	opt *xgrpc.ClientDialOption
}

func NewLinkClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) LinkClient {
	return &linkClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *linkClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *linkClient) UserOnline(req *pb_link.UserOnlineReq) (resp *pb_link.UserOnlineResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_link.NewLinkClient(conn)
	resp, _ = client.UserOnline(context.Background(), req)
	return
}
