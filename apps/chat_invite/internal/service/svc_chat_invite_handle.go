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

	// 2 同意邀请
	if req.HandleResult == pb_enum.INVITE_HANDLE_RESULT_ACCEPT {
		var (
			members     []*po.ChatMember
			member      *po.ChatMember
			mumberCount int
			list        []*po.User
			user        *po.User
			avatars     []*po.UserAvatar
			avatar      *po.UserAvatar
			avatarMaps  = map[int64]string{}
			index       int
			chatHash    string
		)
		w.Reset()
		switch pb_enum.CHAT_TYPE(invite.ChatType) {
		case pb_enum.CHAT_TYPE_PRIVATE:
			mumberCount = 2
			chatHash = utils.MemberHash(invite.InitiatorUid, invite.InviteeUid)
			w.SetFilter("uid IN(?)", []int64{invite.InitiatorUid, invite.InviteeUid})
		case pb_enum.CHAT_TYPE_GROUP:
			mumberCount = 1
			w.SetFilter("uid IN(?)", []int64{invite.InviteeUid})
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

		avatars, err = s.userAvatarRepo.UserAvatarList(w)
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
			avatarMaps[avatar.Uid] = avatar.AvatarMedium
		}
		members = make([]*po.ChatMember, mumberCount)
		for index, user = range list {
			member = &po.ChatMember{
				ChatId:      invite.ChatId,
				ChatHash:    chatHash,
				ChatType:    invite.ChatType,
				Uid:         user.Uid,
				DisplayName: user.Nickname,
				Sync:        1,
				Platform:    user.Platform,
				ServerId:    user.ServerId,
			}
			member.AvatarKey = avatarMaps[member.Uid]
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
