package service

import (
	"context"
	"gorm.io/gorm"
	"lark/apps/request/internal/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_req"
)

func setChatRequestHandlerResp(resp *pb_req.ChatRequestHandlerResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *requestService) ChatRequestHandler(_ context.Context, req *pb_req.ChatRequestHandlerReq) (resp *pb_req.ChatRequestHandlerResp, _ error) {
	var (
		tx      *gorm.DB
		u       = entity.NewMysqlUpdate()
		w       = entity.NewMysqlWhere()
		request *po.ChatRequest
		chatId  int64
		err     error
	)
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	u.Query += " AND request_id=?"
	u.Args = append(u.Args, req.RequestId)
	u.Set("handler_uid", req.HandlerUid)
	u.Set("handle_result", req.HandleResult)
	u.Set("handle_msg", req.HandleMsg)

	tx = xmysql.GetTX()
	err = s.requestRepo.TxRequestUpdate(tx, u)
	if err != nil {
		setChatRequestHandlerResp(resp, ERROR_CODE_REQUEST_UPDATE_VALUE_FAILED, ERROR_REQUEST_UPDATE_VALUE_FAILED)
		xlog.Warn(resp, ERROR_CODE_REQUEST_UPDATE_VALUE_FAILED, ERROR_REQUEST_UPDATE_VALUE_FAILED, err)
		return
	}
	if req.HandleResult == pb_req.REQUEST_HANDLE_RESULT_ACCEPT {
		w.Query += " AND request_id=?"
		w.Args = append(w.Args, req.RequestId)
		request, err = s.requestRepo.TxRequest(tx, w)
		if err != nil {
			setChatRequestHandlerResp(resp, ERROR_CODE_REQUEST_QUERY_DB_FAILED, ERROR_REQUEST_QUERY_DB_FAILED)
			xlog.Warn(resp, ERROR_CODE_REQUEST_QUERY_DB_FAILED, ERROR_REQUEST_QUERY_DB_FAILED, err)
			return
		}
		switch pb_enum.CHAT_TYPE(request.ChatType) {
		case pb_enum.CHAT_TYPE_PRIVATE:
			chatId = xsnowflake.NewSnowflakeID()
			user1 := &po.ChatUser{
				ChatId: chatId,
				Uid:    request.InitiatorUid,
			}
			user2 := &po.ChatUser{
				ChatId: chatId,
				Uid:    request.TargetId,
			}
			err = s.requestRepo.TxChatUsersCreate(tx, []*po.ChatUser{user1, user2})
		case pb_enum.CHAT_TYPE_GROUP:
			user := &po.ChatUser{
				ChatId: request.TargetId,
				Uid:    request.InitiatorUid,
			}
			err = s.requestRepo.TxChatUsersCreate(tx, []*po.ChatUser{user})
		}
		if err != nil {
			setChatRequestHandlerResp(resp, ERROR_CODE_REQUEST_INSERT_VALUE_FAILED, ERROR_REQUEST_INSERT_VALUE_FAILED)
			xlog.Warn(resp, ERROR_CODE_REQUEST_INSERT_VALUE_FAILED, ERROR_REQUEST_INSERT_VALUE_FAILED, err)
			return
		}
		// TODO: 申请成功推送
	}
	return
}
