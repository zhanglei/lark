package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_link"
	"lark/pkg/proto/pb_user"
)

func setUserOnlineResp(resp *pb_link.UserOnlineResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *linkService) UserOnline(ctx context.Context, req *pb_link.UserOnlineReq) (resp *pb_link.UserOnlineResp, err error) {
	var (
		userOnlineReq = &pb_user.UserOnlineReq{
			Uid:      req.Uid,
			ServerId: req.ServerId,
			Platform: req.Platform,
		}
		userOnlineResp *pb_user.UserOnlineResp

		chatMemberOnlineReq = &pb_chat_member.ChatMemberOnlineReq{
			Uid:      req.Uid,
			ServerId: req.ServerId,
			Platform: req.Platform,
		}
		chatMemberOnlineResp *pb_chat_member.ChatMemberOnlineResp
	)
	//0、修改用户状态---socket负责
	//1、更新用户ws服务器 && 用户登录平台
	userOnlineResp = s.userClient.UserOnline(userOnlineReq)
	if userOnlineResp == nil {
		setUserOnlineResp(resp, ERROR_CODE_LINK_GRPC_SERVICE_FAILURE, ERROR_LINK_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_LINK_GRPC_SERVICE_FAILURE, ERROR_LINK_GRPC_SERVICE_FAILURE)
		return
	}
	if userOnlineResp.Code > 0 {
		setUserOnlineResp(resp, userOnlineResp.Code, userOnlineResp.Msg)
		xlog.Warn(userOnlineResp.Code, userOnlineResp.Msg)
	}
	//2、更新chat缓存
	chatMemberOnlineResp = s.chatMemberClient.ChatMemberOnline(chatMemberOnlineReq)
	if chatMemberOnlineResp == nil {
		setUserOnlineResp(resp, ERROR_CODE_LINK_GRPC_SERVICE_FAILURE, ERROR_LINK_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_LINK_GRPC_SERVICE_FAILURE, ERROR_LINK_GRPC_SERVICE_FAILURE)
		return
	}
	if chatMemberOnlineResp.Code > 0 {
		setUserOnlineResp(resp, chatMemberOnlineResp.Code, chatMemberOnlineResp.Msg)
		xlog.Warn(chatMemberOnlineResp.Code, chatMemberOnlineResp.Msg)
	}
	return
}
