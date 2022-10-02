package auth_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_auth"
)

type AuthClient interface {
	Register(req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp)
	Login(req *pb_auth.LoginReq) (resp *pb_auth.LoginResp)
}

type authClient struct {
	opt *xgrpc.ClientDialOption
}

func NewAuthClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) AuthClient {
	return &authClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *authClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *authClient) Register(req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	resp, _ = client.Register(context.Background(), req)
	return
}

func (c *authClient) Login(req *pb_auth.LoginReq) (resp *pb_auth.LoginResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	resp, _ = client.Login(context.Background(), req)
	return
}
