package service

import (
	"context"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
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
		// 不再支持操作
		return
	}

	// 2 更新邀请
	u.SetFilter("invite_id=?", req.InviteId)
	u.Set("handler_uid", req.HandlerUid)
	u.Set("handle_result", req.HandleResult)
	u.Set("handle_msg", req.HandleMsg)

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
			avatars     []*po.Avatar
			avatar      *po.Avatar
			avatarMaps  = map[int64]string{}
			index       int
			uidList     []int64
		)
		if pb_enum.CHAT_TYPE(invite.ChatType) == pb_enum.CHAT_TYPE_PRIVATE {
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
		}

		w.Reset()
		switch pb_enum.CHAT_TYPE(invite.ChatType) {
		case pb_enum.CHAT_TYPE_PRIVATE:
			mumberCount = 2
			uidList = []int64{invite.InitiatorUid, invite.InviteeUid}
			w.SetFilter("uid IN(?)", uidList)
		case pb_enum.CHAT_TYPE_GROUP:
			mumberCount = 1
			uidList = []int64{invite.InviteeUid}
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
		w.Reset()
		w.SetFilter("owner_id IN(?)", uidList)
		w.SetFilter("owner_type=?", int32(pb_enum.AVATAR_OWNER_USER_AVATAR))
		avatars, err = s.avatarRepo.AvatarList(w)
		if err != nil {
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
			return
		}
		if len(avatars) != mumberCount {
			err = ERR_CHAT_INVITE_QUERY_DB_FAILED
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			return
		}

		for _, avatar = range avatars {
			avatarMaps[avatar.OwnerId] = avatar.AvatarMedium
		}
		members = make([]*po.ChatMember, mumberCount)
		for index, user = range list {
			member = &po.ChatMember{
				ChatId:      invite.ChatId,
				ChatType:    invite.ChatType,
				Uid:         user.Uid,
				DisplayName: user.Nickname,
				Sync:        1,
				Platform:    user.Platform,
				ServerId:    user.ServerId,
			}
			member.MemberAvatarKey = avatarMaps[member.Uid]
			members[index] = member
		}
		err = s.chatInviteRepo.TxChatUsersCreate(tx, members)
		if err != nil {
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED, err)
			return
		}
		// TODO: 邀请成功推送
	}
	return
}
