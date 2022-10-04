package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xgopool"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
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
		w        = entity.NewMysqlWhere()
		list     []*po.ChatMember
		members  []*pb_chat_member.ChatMemberInfo
		member   *po.ChatMember
		pbMember *pb_chat_member.ChatMemberInfo
		err      error
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
			for _, member = range list {
				pbMember = new(pb_chat_member.ChatMemberInfo)
				copier.Copy(pbMember, member)
				if pbMember.Uid == req.UidList[0] {
					resp.MemberInfo = pbMember
				}
				members = append(members, pbMember)
			}
		}
	case pb_enum.CHAT_TYPE_GROUP:
		if len(list) == 1 {
			pbMember = new(pb_chat_member.ChatMemberInfo)
			copier.Copy(pbMember, list[0])
			resp.MemberInfo = pbMember
			members = append(members, pbMember)
		}
	}
	xgopool.Go(func() {
		s.CacheChatMemberInfo(req.ChatId, members)
	})
	return
}

func (s *chatMemberService) CacheChatMemberInfo(chatId int64, list []*pb_chat_member.ChatMemberInfo) {
	if len(list) == 0 {
		return
	}
	var (
		pbMember *pb_chat_member.ChatMemberInfo
		key      string
		jsonStr  string
	)
	for _, pbMember = range list {
		jsonStr, _ = utils.Marshal(pbMember)
		if jsonStr == "" {
			continue
		}
		key = constant.RK_SYNC_CHAT_MEMBERS_INFO_HASH + utils.Int64ToStr(chatId)
		xredis.HSetNX(key, utils.Int64ToStr(pbMember.Uid), jsonStr)
	}
}
