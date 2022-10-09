package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

func setGetPushMemberResp(resp *pb_chat_member.GetPushMemberResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetPushMember(ctx context.Context, req *pb_chat_member.GetPushMemberReq) (resp *pb_chat_member.GetPushMemberResp, _ error) {
	resp = &pb_chat_member.GetPushMemberResp{Member: &pb_chat_member.PushMember{}}
	var (
		w      = entity.NewMysqlWhere()
		member *pb_chat_member.PushMember
		err    error
	)
	w.SetFilter("chat_id = ?", req.ChatId)
	w.SetFilter("uid = ?", req.Uid)

	member, err = s.chatMemberRepo.PushMember(w)
	if err != nil {
		setGetPushMemberResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	resp.Member = member
	return
}
