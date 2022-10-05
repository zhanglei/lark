package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_invite"
	"lark/pkg/utils"
)

func setNewChatInviteResp(resp *pb_invite.NewChatInviteResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatInviteService) NewChatInvite(_ context.Context, req *pb_invite.NewChatInviteReq) (resp *pb_invite.NewChatInviteResp, _ error) {
	resp = new(pb_invite.NewChatInviteResp)
	var (
		invite = new(po.ChatInvite)
		w      = entity.NewMysqlWhere()
		chat   *po.Chat
		member *po.ChatMember
		err    error
	)

	if req.ChatType == pb_enum.CHAT_TYPE_PRIVATE {
		if req.InitiatorUid == req.InviteeUid {
			setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_INITIATOR_INVITEE_SAME, ERROR_CHAT_INVITE_INITIATOR_INVITEE_SAME)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_INITIATOR_INVITEE_SAME, ERROR_CHAT_INVITE_INITIATOR_INVITEE_SAME)
			return
		}
		w.SetFilter("chat_type=?", int32(pb_enum.CHAT_TYPE_PRIVATE))
		w.SetFilter("chat_hash=?", utils.MemberHash(req.InitiatorUid, req.InviteeUid))
		chat, err = s.chatRepo.Chat(w)
		if err != nil {
			setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err.Error())
			return
		}
		if chat.ChatId > 0 {
			setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_HAS_JOINED_CHAT, ERROR_CHAT_INVITE_HAS_JOINED_CHAT)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_HAS_JOINED_CHAT, ERROR_CHAT_INVITE_HAS_JOINED_CHAT, req.String())
			return
		}
		req.ChatId = xsnowflake.NewSnowflakeID()
	} else {
		if req.ChatId <= 0 {
			setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_PARAM_ERR, ERROR_CHAT_INVITE_PARAM_ERR)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_PARAM_ERR, ERROR_CHAT_INVITE_PARAM_ERR, req.String())
			return
		}
		w.SetFilter("chat_id=?", req.ChatId)
		w.SetFilter("uid=?", req.InviteeUid)
		member, err = s.chatMemberRepo.ChatMember(w)
		if err != nil {
			setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err.Error())
			return
		}
		if member.Uid > 0 {
			setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_HAS_JOINED_CHAT, ERROR_CHAT_INVITE_HAS_JOINED_CHAT)
			xlog.Warn(ERROR_CODE_CHAT_INVITE_HAS_JOINED_CHAT, ERROR_CHAT_INVITE_HAS_JOINED_CHAT, req.String())
			return
		}
	}

	copier.Copy(invite, req)
	invite.InviteId = xsnowflake.NewSnowflakeID()

	err = s.chatInviteRepo.NewChatInvite(invite)
	if err != nil {
		setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED, err)
		return
	}
	return
}
