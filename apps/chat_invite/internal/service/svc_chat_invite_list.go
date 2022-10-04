package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_req"
)

func setChatRequestListResp(resp *pb_req.ChatRequestListResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatInviteService) ChatRequestList(_ context.Context, req *pb_req.ChatRequestListReq) (resp *pb_req.ChatRequestListResp, _ error) {
	resp = &pb_req.ChatRequestListResp{List: make([]*pb_req.ChatRequestInfo, 0)}
	var (
		w    = entity.NewMysqlWhere()
		list []po.ChatRequest
		err  error
	)
	w.Limit = int(req.PageSize)
	w.Query = " AND request_id>?"
	w.Args = append(w.Args, req.MaxRequestId)
	if req.HandleResult > 0 {
		w.Query = " AND handle_result=?"
		w.Args = append(w.Args, req.HandleResult)
	}
	switch req.Role {
	case pb_req.REQUEST_ROLE_INITIATOR: // 发起者
		w.Query = " AND initiator_uid=?"
		w.Args = append(w.Args, req.Uid)
	case pb_req.REQUEST_ROLE_APPROVER: // 审批人
		w.Query = " AND target_id=?"
		w.Args = append(w.Args, req.Uid)
	}
	list, err = s.chatInviteRepo.RequestList(w)
	if err != nil {
		setChatRequestListResp(resp, ERROR_CODE_REQUEST_QUERY_DB_FAILED, ERROR_REQUEST_QUERY_DB_FAILED)
		xlog.Warn(resp, ERROR_CODE_REQUEST_QUERY_DB_FAILED, ERROR_REQUEST_QUERY_DB_FAILED, err)
		return
	}
	copier.Copy(resp.List, list)
	return
}
