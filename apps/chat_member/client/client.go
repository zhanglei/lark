package chat_member_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_chat_member"
)

type ChatMemberClient interface {
	GetChatMemberUidList(req *pb_chat_member.GetChatMemberUidListReq) (resp *pb_chat_member.GetChatMemberUidListResp)
	GetChatMemberSetting(req *pb_chat_member.GetChatMemberSettingReq) (resp *pb_chat_member.GetChatMemberSettingResp)
	GetChatMemberInfo(req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp)
	ChatMemberVerify(req *pb_chat_member.ChatMemberVerifyReq) (resp *pb_chat_member.ChatMemberVerifyResp)
	ChatMemberOnline(req *pb_chat_member.ChatMemberOnlineReq) (resp *pb_chat_member.ChatMemberOnlineResp)
	GetPushMemberList(req *pb_chat_member.GetPushMemberListReq) (resp *pb_chat_member.GetPushMemberListResp)
	GetPushMember(req *pb_chat_member.GetPushMemberReq) (resp *pb_chat_member.GetPushMemberResp)
	GetChatMemberList(req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp)
}

type chatMemberClient struct {
	opt *xgrpc.ClientDialOption
}

func NewChatMemberClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ChatMemberClient {
	return &chatMemberClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *chatMemberClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *chatMemberClient) GetChatMemberUidList(req *pb_chat_member.GetChatMemberUidListReq) (resp *pb_chat_member.GetChatMemberUidListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.GetChatMemberUidList(context.Background(), req)
	return
}

func (c *chatMemberClient) GetChatMemberSetting(req *pb_chat_member.GetChatMemberSettingReq) (resp *pb_chat_member.GetChatMemberSettingResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.GetChatMemberSetting(context.Background(), req)
	return
}

func (c *chatMemberClient) GetChatMemberInfo(req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.GetChatMemberInfo(context.Background(), req)
	return
}

func (c *chatMemberClient) ChatMemberVerify(req *pb_chat_member.ChatMemberVerifyReq) (resp *pb_chat_member.ChatMemberVerifyResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.ChatMemberVerify(context.Background(), req)
	return
}

func (c *chatMemberClient) ChatMemberOnline(req *pb_chat_member.ChatMemberOnlineReq) (resp *pb_chat_member.ChatMemberOnlineResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.ChatMemberOnline(context.Background(), req)
	return
}

func (c *chatMemberClient) GetPushMemberList(req *pb_chat_member.GetPushMemberListReq) (resp *pb_chat_member.GetPushMemberListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.GetPushMemberList(context.Background(), req)
	return
}

func (c *chatMemberClient) GetPushMember(req *pb_chat_member.GetPushMemberReq) (resp *pb_chat_member.GetPushMemberResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.GetPushMember(context.Background(), req)
	return
}

func (c *chatMemberClient) GetChatMemberList(req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	resp, _ = client.GetChatMemberList(context.Background(), req)
	return
}
