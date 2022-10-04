package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/utils"
)

func setGetChatMemberPushConfigListResp(resp *pb_chat_member.GetChatMemberPushConfigListResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetChatMemberPushConfigList(ctx context.Context, req *pb_chat_member.GetChatMemberPushConfigListReq) (resp *pb_chat_member.GetChatMemberPushConfigListResp, _ error) {
	resp = new(pb_chat_member.GetChatMemberPushConfigListResp)
	var (
		w          = entity.NewMysqlWhere()
		count      int
		lastUid    int64
		configList []*pb_chat_member.ChatMemberPushConfig
		err        error
	)

	w.Sort = "uid ASC"
	w.Limit = 10000
	for {
		w.Args = nil
		w.Query = "chat_id = ?"
		w.Args = append(w.Args, req.ChatId)
		w.Query += " AND uid > ?"
		w.Args = append(w.Args, lastUid)
		configList, err = s.chatMemberRepo.ChatMemberPushConfigList(w)
		if err != nil {
			setGetChatMemberPushConfigListResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
			return
		}
		count = len(configList)
		if count == 0 {
			break
		}
		resp.List = append(resp.List, configList...)
		if count < w.Limit {
			break
		}
		lastUid = configList[count-1].Uid
		go s.cacheChatMemberPushConfig(configList, req.ChatId)
	}
	return
}

func (s *chatMemberService) cacheChatMemberPushConfig(list []*pb_chat_member.ChatMemberPushConfig, chatId int64) {
	if len(list) == 0 {
		return
	}
	var (
		key     string
		conf    *pb_chat_member.ChatMemberPushConfig
		jsonStr string
	)
	for _, conf = range list {
		jsonStr, _ = utils.Marshal(conf)
		key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_CONF_HASH + utils.Int64ToStr(chatId)
		xredis.HSetNX(key, utils.Int64ToStr(conf.Uid), jsonStr)
	}
}
