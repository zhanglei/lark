package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/apps/chat_member/internal/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

func setGetChatMemberInfoResp(resp *pb_chat_member.GetChatMemberInfoResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetChatMemberInfo(ctx context.Context, req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp, _ error) {
	resp = &pb_chat_member.GetChatMemberInfoResp{Info: new(pb_chat_member.ChatMemberInfo)}
	var (
		w      = entity.NewMysqlWhere()
		member *po.ChatMember
		err    error
	)
	w.Query += " AND chat_id = ?"
	w.Args = append(w.Args, req.ChatId)
	w.Query += " AND uid = ?"
	w.Args = append(w.Args, req.Uid)
	member, err = s.chatMemberUserRepo.ChatMemberInfo(w)
	if err != nil {
		setGetChatMemberInfoResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.Info, member)
	return
}
