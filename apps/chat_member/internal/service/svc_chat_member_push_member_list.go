package service

import (
	"context"
	"fmt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/utils"
)

func setGetPushMemberListResp(resp *pb_chat_member.GetPushMemberListResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetPushMemberList(ctx context.Context, req *pb_chat_member.GetPushMemberListReq) (resp *pb_chat_member.GetPushMemberListResp, _ error) {
	resp = &pb_chat_member.GetPushMemberListResp{List: make([]string, 0)}
	var (
		w       = entity.NewMysqlWhere()
		count   int
		lastUid int64
		members []*pb_chat_member.PushMember
		member  *pb_chat_member.PushMember
		key     = constant.RK_SYNC_CHAT_MEMBERS_PUSH_MEMBER_HASH + utils.Int64ToStr(req.ChatId)
		err     error
	)

	w.Sort = "uid ASC"
	w.Limit = 5000
	for {
		var (
			values []string
			index  int
			value  string
			maps   = make(map[string]interface{})
		)
		w.Args = nil
		w.Query = "chat_id = ?"
		w.Args = append(w.Args, req.ChatId)
		w.Query += " AND uid > ?"
		w.Args = append(w.Args, lastUid)
		members, err = s.chatMemberRepo.PushMemberList(w)
		if err != nil {
			setGetPushMemberListResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
			return
		}
		count = len(members)
		if count == 0 {
			break
		}
		values = make([]string, count)
		for index, member = range members {
			value = fmt.Sprintf("%d,%d,%d,%d", member.Uid, member.Platform, member.ServerId, member.Mute)
			values[index] = value
			maps[utils.Int64ToStr(member.Uid)] = value
		}
		err = xredis.HMSet(key, maps)
		if err != nil {
			setGetPushMemberListResp(resp, ERROR_CODE_CHAT_MEMBER_CHCHE_MEMBER_FAILED, ERROR_CHAT_MEMBER_CHCHE_MEMBER_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_MEMBER_CHCHE_MEMBER_FAILED, ERROR_CHAT_MEMBER_CHCHE_MEMBER_FAILED, err.Error())
			return
		}
		resp.List = append(resp.List, values...)
		if count < w.Limit {
			break
		}
		lastUid = members[count-1].Uid
	}
	return
}

func (s *chatMemberService) cachePushMember(list []*pb_chat_member.PushMember, chatId int64) (err error) {
	if len(list) == 0 {
		return
	}
	var (
		key     string
		conf    *pb_chat_member.PushMember
		jsonStr string
		members = map[string]interface{}{}
	)
	for _, conf = range list {
		jsonStr, _ = utils.Marshal(conf)
		members[utils.Int64ToStr(conf.Uid)] = jsonStr
	}
	key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_MEMBER_HASH + utils.Int64ToStr(chatId)
	err = xredis.HMSet(key, members)
	return
}
