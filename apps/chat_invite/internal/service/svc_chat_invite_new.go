package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/proto/pb_invite"
)

func setNewChatInviteResp(resp *pb_invite.NewChatInviteResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatInviteService) NewChatInvite(_ context.Context, req *pb_invite.NewChatInviteReq) (resp *pb_invite.NewChatInviteResp, _ error) {
	resp = new(pb_invite.NewChatInviteResp)
	var (
		invite = new(po.ChatInvite)
		err    error
	)
	copier.Copy(invite, req)
	invite.InviteId = xsnowflake.NewSnowflakeID()
	err = s.chatInviteRepo.NewChatInvite(invite)
	if err != nil {
		setNewChatInviteResp(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED)
		xlog.Warn(resp, ERROR_CODE_CHAT_INVITE_INSERT_VALUE_FAILED, ERROR_CHAT_INVITE_INSERT_VALUE_FAILED, err)
		return
	}
	return
}
