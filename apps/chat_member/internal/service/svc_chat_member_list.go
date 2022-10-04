package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

func setGetChatMemberListResp(resp *pb_chat_member.GetChatMemberListResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetChatMemberList(ctx context.Context, req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp, _ error) {
	resp = &pb_chat_member.GetChatMemberListResp{List: make([]*pb_chat_member.ChatMemberBasicInfo, 0)}
	var (
		w    = entity.NewMysqlWhere()
		list []*pb_chat_member.ChatMemberBasicInfo
		err  error
	)
	w.AndQuery("chat_id=?")
	w.AppendArg(req.ChatId)
	w.AndQuery("uid>?")
	w.AppendArg(req.LastUid)
	w.SetLimit(req.Limit)
	list, err = s.chatMemberUserRepo.ChatMemberBasicInfoList(w)
	if err != nil {
		setGetChatMemberListResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	resp.List = list
	return
}
