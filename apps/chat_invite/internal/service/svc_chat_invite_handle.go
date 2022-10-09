package service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_invite"
	"lark/pkg/utils"
)

func setChatInviteHandleResp(resp *pb_invite.ChatInviteHandleResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatInviteService) ChatInviteHandle(ctx context.Context, req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp, _ error) {
	resp = new(pb_invite.ChatInviteHandleResp)
	var (
		tx     *gorm.DB
		u      = entity.NewMysqlUpdate()
		w      = entity.NewMysqlWhere()
		invite *po.ChatInvite
		err    error
	)
	// 1 校验邀请
	w.SetFilter("invite_id=?", req.InviteId)
	invite, err = s.chatInviteRepo.ChatInvite(w)
	if err != nil {
		setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
		return
	}
	if invite.InviteId == 0 {
		setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
		return
	}
	if invite.HandleResult != 0 {
		setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_HAS_HANDLED, ERROR_CHAT_INVITE_HAS_HANDLED)
		xlog.Warn(ERROR_CODE_CHAT_INVITE_HAS_HANDLED, ERROR_CHAT_INVITE_HAS_HANDLED)
		return
	}
	if req.HandlerUid != invite.InviteeUid {
		setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_BAD_HANDLER, ERROR_CHAT_INVITE_BAD_HANDLER)
		xlog.Warn(ERROR_CODE_CHAT_INVITE_BAD_HANDLER, ERROR_CHAT_INVITE_BAD_HANDLER)
		return
	}
	// 2 更新邀请
	u.SetFilter("invite_id=?", req.InviteId)
	u.Set("handler_uid", req.HandlerUid)
	u.Set("handle_result", req.HandleResult)
	u.Set("handle_msg", req.HandleMsg)
	u.Set("handled_ts", utils.NowMilli())

	tx = xmysql.GetTX()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = s.chatInviteRepo.TxUpdateChatInvite(tx, u)
	if err != nil {
		setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_UPDATE_VALUE_FAILED, ERROR_CHAT_INVITE_UPDATE_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INVITE_UPDATE_VALUE_FAILED, ERROR_CHAT_INVITE_UPDATE_VALUE_FAILED, err)
		return
	}

	// 3 同意邀请
	if req.HandleResult == pb_enum.INVITE_HANDLE_RESULT_ACCEPT {
		var (
			chat        *po.Chat
			members     []*po.ChatMember
			member      *po.ChatMember
			mumberCount int
			list        []*po.User
			user        *po.User
			index       int
			uidList     []int64
			pushKey     string
			pushMaps    map[string]interface{}
		)

		switch pb_enum.CHAT_TYPE(invite.ChatType) {
		case pb_enum.CHAT_TYPE_PRIVATE:
			// 4 创建Chat
			chat = &po.Chat{
				ChatId:   invite.ChatId,
				ChatHash: utils.MemberHash(invite.InitiatorUid, invite.InviteeUid),
				ChatType: int(pb_enum.CHAT_TYPE_PRIVATE),
			}
			err = s.chatRepo.TxCreate(tx, chat)
			if err != nil {
				setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
				xlog.Warn(ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED, err.Error())
				return
			}

			uidList = []int64{invite.InitiatorUid, invite.InviteeUid}
			mumberCount = len(uidList)
			w.Reset()
			w.SetFilter("uid IN(?)", uidList)
		case pb_enum.CHAT_TYPE_GROUP:
			w.Reset()
			w.SetFilter("chat_id=?", invite.ChatId)
			chat, err = s.chatRepo.TxChat(tx, w)
			if err != nil {
				setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
				xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
				return
			}

			uidList = []int64{invite.InviteeUid}
			mumberCount = len(uidList)
			w.Reset()
			w.SetFilter("uid IN(?)", uidList)
		}
		list, err = s.userRepo.UserList(w)
		if err != nil {
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
			return
		}
		if len(list) != mumberCount {
			err = ERR_CHAT_INVITE_QUERY_DB_FAILED
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			return
		}
		members = make([]*po.ChatMember, mumberCount)
		for index, user = range list {
			member = &po.ChatMember{
				ChatId:          invite.ChatId,
				ChatType:        invite.ChatType,
				Uid:             user.Uid,
				DisplayName:     user.Nickname,
				MemberAvatarKey: user.AvatarKey,
				Sync:            1,
				Platform:        user.Platform,
				ServerId:        user.ServerId,
			}
			members[index] = member
			if pb_enum.CHAT_TYPE(invite.ChatType) == pb_enum.CHAT_TYPE_GROUP {
				member.ChatAvatarKey = chat.AvatarKey
			}
		}
		// 5 成为 chat member
		err = s.chatMemberRepo.TxCreateMultiple(tx, members)
		if err != nil {
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED, err)
			return
		}
		pushMaps = make(map[string]interface{})
		for _, member = range members {
			pushMaps[utils.Int64ToStr(member.Uid)] = fmt.Sprintf("%d,%d,%d,%d", member.Uid, member.Platform, member.ServerId, member.Mute)
		}
		// 6 缓存 chat member
		pushKey = constant.RK_SYNC_CHAT_MEMBERS_PUSH_MEMBER_HASH + utils.Int64ToStr(member.ChatId)
		err = xredis.HMSet(pushKey, pushMaps)
		if err != nil {
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_CACHE_CHAT_MEMBER_FAILED, ERROR_CHAT_INVITE_CACHE_CHAT_MEMBER_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_CACHE_CHAT_MEMBER_FAILED, ERROR_CHAT_INVITE_CACHE_CHAT_MEMBER_FAILED, err)
			return
		}
		// TODO: 邀请成功推送

	}
	return
}
