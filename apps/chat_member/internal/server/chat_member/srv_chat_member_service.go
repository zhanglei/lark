package chat_member

import (
	"context"
	"lark/pkg/proto/pb_chat_member"
)

func (s *chatMemberServer) GetChatMemberUidList(ctx context.Context, req *pb_chat_member.GetChatMemberUidListReq) (resp *pb_chat_member.GetChatMemberUidListResp, err error) {
	return s.chatMemberService.GetChatMemberUidList(ctx, req)
}

func (s *chatMemberServer) GetChatMemberSetting(ctx context.Context, req *pb_chat_member.GetChatMemberSettingReq) (resp *pb_chat_member.GetChatMemberSettingResp, err error) {
	return s.chatMemberService.GetChatMemberSetting(ctx, req)
}

func (s *chatMemberServer) GetChatMemberInfo(ctx context.Context, req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp, err error) {
	return s.chatMemberService.GetChatMemberInfo(ctx, req)
}

func (s *chatMemberServer) ChatMemberVerify(ctx context.Context, req *pb_chat_member.ChatMemberVerifyReq) (resp *pb_chat_member.ChatMemberVerifyResp, err error) {
	return s.chatMemberService.ChatMemberVerify(ctx, req)
}

func (s *chatMemberServer) ChatMemberOnline(ctx context.Context, req *pb_chat_member.ChatMemberOnlineReq) (resp *pb_chat_member.ChatMemberOnlineResp, err error) {
	return s.chatMemberService.ChatMemberOnline(ctx, req)
}

func (s *chatMemberServer) GetChatMemberPushConfigList(ctx context.Context, req *pb_chat_member.GetChatMemberPushConfigListReq) (resp *pb_chat_member.GetChatMemberPushConfigListResp, err error) {
	return s.chatMemberService.GetChatMemberPushConfigList(ctx, req)
}

func (s *chatMemberServer) GetChatMemberPushConfig(ctx context.Context, req *pb_chat_member.GetChatMemberPushConfigReq) (resp *pb_chat_member.GetChatMemberPushConfigResp, err error) {
	return s.chatMemberService.GetChatMemberPushConfig(ctx, req)
}
