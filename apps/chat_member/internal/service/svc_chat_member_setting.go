package service

import (
	"context"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/utils"
)

func setGetChatMemberSettingResp(resp *pb_chat_member.GetChatMemberSettingResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetChatMemberSetting(ctx context.Context, req *pb_chat_member.GetChatMemberSettingReq) (resp *pb_chat_member.GetChatMemberSettingResp, _ error) {
	resp = new(pb_chat_member.GetChatMemberSettingResp)
	var (
		w    = entity.NewMysqlWhere()
		user *po.ChatMember
		err  error
	)
	w.Query += " AND chat_id = ?"
	w.Args = append(w.Args, req.ChatId)
	w.Query += " AND uid = ?"
	w.Args = append(w.Args, req.Uid)
	user, err = s.chatMemberUserRepo.ChatMemberSetting(w)
	if err != nil {
		setGetChatMemberSettingResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	resp.Setting = user.Settings
	go func(chatId int64, uid int64) {
		key := constant.RK_SYNC_CHAT_MEMBERS_SETTINGS_HASH + utils.Int64ToStr(chatId)
		xredis.HSetNX(key, utils.Int64ToStr(user.Uid), user.Settings)
	}(req.ChatId, req.Uid)
	return
}
