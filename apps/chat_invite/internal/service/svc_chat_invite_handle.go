package service

import (
	"context"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_invite"
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
		chatId int64
		err    error
	)
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	u.SetFilter("invite_id=?", req.InviteId)

	u.Set("handler_uid", req.HandlerUid)
	u.Set("handle_result", req.HandleResult)
	u.Set("handle_msg", req.HandleMsg)

	tx = xmysql.GetTX()
	err = s.chatInviteRepo.TxUpdateChatInvite(tx, u)
	if err != nil {
		setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_UPDATE_VALUE_FAILED, ERROR_CHAT_INVITE_UPDATE_VALUE_FAILED)
		xlog.Warn(resp, ERROR_CODE_CHAT_INVITE_UPDATE_VALUE_FAILED, ERROR_CHAT_INVITE_UPDATE_VALUE_FAILED, err)
		return
	}
	if req.HandleResult == pb_enum.INVITE_HANDLE_RESULT_ACCEPT {
		w.SetFilter("invite_id=?", req.InviteId)
		invite, err = s.chatInviteRepo.TxChatInvite(tx, w)
		if err != nil {
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
			xlog.Warn(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
			return
		}
		switch pb_enum.CHAT_TYPE(invite.ChatType) {
		case pb_enum.CHAT_TYPE_PRIVATE:
			chatId = xsnowflake.NewSnowflakeID()
			member1 := &po.ChatMember{
				ChatId: chatId,
				Uid:    invite.InitiatorUid,
			}
			member2 := &po.ChatMember{
				ChatId: chatId,
				Uid:    invite.InviteeUid,
			}
			err = s.chatInviteRepo.TxChatUsersCreate(tx, []*po.ChatMember{member1, member2})
		case pb_enum.CHAT_TYPE_GROUP:
			member := &po.ChatMember{
				ChatId: invite.InviteeUid,
				Uid:    invite.InitiatorUid,
			}
			err = s.chatInviteRepo.TxChatUsersCreate(tx, []*po.ChatMember{member})
		}
		if err != nil {
			setChatInviteHandleResp(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
			xlog.Warn(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED, err)
			return
		}
		// TODO: 申请成功推送
	}
	return
}
