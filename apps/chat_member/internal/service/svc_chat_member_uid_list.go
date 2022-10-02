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

func setGetChatMemberUidListResp(resp *pb_chat_member.GetChatMemberUidListResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatMemberService) GetChatMemberUidList(ctx context.Context, req *pb_chat_member.GetChatMemberUidListReq) (resp *pb_chat_member.GetChatMemberUidListResp, _ error) {
	resp = new(pb_chat_member.GetChatMemberUidListResp)
	var (
		jsonStr string
		err     error
	)
	jsonStr, err = xredis.Get(constant.RK_SYNC_CHAT_MEMBERS_UID_LIST + utils.Int64ToStr(req.ChatId))
	if jsonStr != "" {
		resp.List = make([]int64, 0)
		utils.Unmarshal(jsonStr, &resp.List)
	} else {
		var (
			w       = new(entity.MysqlWhere)
			lastUid int64
			uidList []int64
			count   int
		)
		w.Sort = "uid ASC"
		w.Limit = 10000
		for {
			w.Query = ""
			w.Args = nil
			w.Query += "chat_id = ?"
			w.Args = append(w.Args, req.ChatId)
			w.Query += " AND uid > ?"
			w.Args = append(w.Args, lastUid)

			uidList, err = s.chatMemberUserRepo.ChatMemberUidList(w)
			if err != nil {
				setGetChatMemberUidListResp(resp, ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
				xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
				return
			}
			count = len(uidList)
			if count == 0 {
				break
			}
			resp.List = append(resp.List, uidList...)
			if count < w.Limit {
				break
			}
			lastUid = uidList[count-1]
		}
		jsonStr, err = utils.Marshal(resp.List)
		if err != nil {
			return
		}
		xredis.Set(constant.RK_SYNC_CHAT_MEMBERS_UID_LIST+utils.Int64ToStr(req.ChatId), jsonStr, constant.CONST_DURATION_CHAT_GROUP_UID_LIST_SECOND)
	}
	return
}
