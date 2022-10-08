package service

import (
	"context"
	"lark/domain/do"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
)

func setChatMemberVerifyResp(resp *pb_chat_member.ChatMemberVerifyResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) ChatMemberVerify(ctx context.Context, req *pb_chat_member.ChatMemberVerifyReq) (resp *pb_chat_member.ChatMemberVerifyResp, _ error) {
	resp = new(pb_chat_member.ChatMemberVerifyResp)
	if len(req.UidList) == 0 {
		return
	}
	var (
		w    = entity.NewMysqlWhere()
		list []*do.ChatMemberInfo
		err  error
	)
	w.Query += " AND chat_id = ?"
	w.Args = append(w.Args, req.ChatId)

	w.Query += " AND uid IN(?)"
	w.Args = append(w.Args, req.UidList)
	list, err = s.chatMemberRepo.ChatMemberList(w)
	if err != nil {
		setChatMemberVerifyResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	switch req.ChatType {
	case pb_enum.CHAT_TYPE_PRIVATE:
		if len(list) == 2 {
			resp.Ok = true
		}
	case pb_enum.CHAT_TYPE_GROUP:
		if len(list) == 1 {
			resp.Ok = true
		}
	}
	// 在push service 缓存
	//err = s.cachePushMembers(list)
	//if err != nil {
	//	setChatMemberVerifyResp(resp, ERROR_CODE_CHAT_MEMBER_CHCHE_MEMBER_FAILED, ERROR_CHAT_MEMBER_CHCHE_MEMBER_FAILED)
	//	xlog.Warn(ERROR_CODE_CHAT_MEMBER_CHCHE_MEMBER_FAILED, ERROR_CHAT_MEMBER_CHCHE_MEMBER_FAILED, err.Error())
	//	return
	//}
	return
}
