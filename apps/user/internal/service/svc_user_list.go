package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/pos"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func setGetUserListResp(resp *pb_user.GetUserListResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *userService) GetUserList(ctx context.Context, req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp, _ error) {
	resp = &pb_user.GetUserListResp{List: make([]*pb_user.UserInfo, 0)}
	var (
		w    = entity.NewMysqlWhere()
		list []pos.User
		err  error
	)
	w.Query += " AND uid IN(?)"
	w.Args = append(w.Args, req.UidList)
	list, err = s.userRepo.UserList(w)
	if err != nil {
		setGetUserListResp(resp, ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.List, list)
	return
}
