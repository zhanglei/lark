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

func (s *chatMemberServer) GetPushMemberList(ctx context.Context, req *pb_chat_member.GetPushMemberListReq) (resp *pb_chat_member.GetPushMemberListResp, err error) {
	return s.chatMemberService.GetPushMemberList(ctx, req)
}
func (s *chatMemberServer) GetPushMember(ctx context.Context, req *pb_chat_member.GetPushMemberReq) (resp *pb_chat_member.GetPushMemberResp, err error) {
	return s.chatMemberService.GetPushMember(ctx, req)
}

func (s *chatMemberServer) GetChatMemberList(ctx context.Context, req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp, err error) {
	return s.chatMemberService.GetChatMemberList(ctx, req)
}
