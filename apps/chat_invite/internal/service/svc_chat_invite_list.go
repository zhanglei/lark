package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_invite"
)

func setChatInviteListResp(resp *pb_invite.ChatInviteListResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatInviteService) ChatInviteList(_ context.Context, req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp, _ error) {
	resp = &pb_invite.ChatInviteListResp{List: make([]*pb_invite.ChatInviteInfo, 0)}
	var (
		w    = entity.NewMysqlWhere()
		list []*po.ChatInvite
		err  error
	)
	w.Limit = int(req.Limit)
	w.And("invite_id>?", req.MaxInviteId)

	if req.HandleResult > 0 {
		w.And("handle_result=?", req.HandleResult)
	}
	switch req.Role {
	case pb_enum.INVITE_ROLE_INITIATOR: // 发起者
		w.And("initiator_uid=?", req.Uid)
	case pb_enum.INVITE_ROLE_APPROVER: // 审批人
		w.And("invitee_uid=?", req.Uid)
	}
	list, err = s.chatInviteRepo.ChatInviteList(w)
	if err != nil {
		setChatInviteListResp(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
		xlog.Warn(resp, ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
		return
	}
	copier.Copy(&resp.List, list)
	return
}
