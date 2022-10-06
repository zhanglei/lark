package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_kv"
)

func setGroupChatResp(resp *pb_chat.SetGroupChatResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatService) SetGroupChat(ctx context.Context, req *pb_chat.SetGroupChatReq) (resp *pb_chat.SetGroupChatResp, _ error) {
	var (
		u   = entity.NewMysqlUpdate()
		skv *pb_kv.StrKeyValue
		ok  bool
		err error
	)
	u.SetFilter("chat_id=?", req.ChatId)
	for _, skv = range req.Kvs.StrList {
		if _, ok = chatUpdateFields[skv.Key]; ok == false {
			continue
		}
		u.Set(skv.Key, skv.Value)
	}
	err = s.chatRepo.UpdateChat(u)
	if err != nil {
		setGroupChatResp(resp, ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED, err.Error())
		return
	}
	return
}
