package user_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_user"
)

type UserClient interface {
	GetUserList(req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp)
	GetChatUserInfo(req *pb_user.GetChatUserInfoReq) (resp *pb_user.GetChatUserInfoResp)
	GetUserInfo(req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp)
	UserOnline(req *pb_user.UserOnlineReq) (resp *pb_user.UserOnlineResp)
}

type userClient struct {
	opt *xgrpc.ClientDialOption
}

func NewUserClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) UserClient {
	return &userClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *userClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *userClient) GetUserList(req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	resp, _ = client.GetUserList(context.Background(), req)
	return
}

func (c *userClient) GetChatUserInfo(req *pb_user.GetChatUserInfoReq) (resp *pb_user.GetChatUserInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	resp, _ = client.GetChatUserInfo(context.Background(), req)
	return
}

func (c *userClient) GetUserInfo(req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	resp, _ = client.GetUserInfo(context.Background(), req)
	return
}

func (c *userClient) UserOnline(req *pb_user.UserOnlineReq) (resp *pb_user.UserOnlineResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	resp, _ = client.UserOnline(context.Background(), req)
	return
}
