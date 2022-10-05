package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func setUserInfoResp(resp *pb_user.UserInfoResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *userService) GetUserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, _ error) {
	resp = &pb_user.UserInfoResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_user.UserAvatar{}}}
	var (
		key     string
		jsonStr string
		err     error
	)
	key = constant.RK_SYNC_USERS_INFO + utils.Int64ToStr(req.Uid)
	jsonStr, err = xredis.Get(key)
	if err != nil {
		//TODO:读取Redis失败
		xlog.Warn(ERROR_CODE_USER_REDIS_GET_FAILED, ERROR_USER_REDIS_GET_FAILED, err.Error())
	}
	if jsonStr != "" {
		utils.Unmarshal(jsonStr, resp.UserInfo)
		return
	}
	err = s.queryUserInfo(req.Uid, resp)
	if err != nil {
		setUserInfoResp(resp, ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	err = s.queryUserAvatar(resp.UserInfo.Uid, resp)
	if err != nil {
		setUserInfoResp(resp, ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	go func(r *pb_user.UserInfoResp) {
		var (
			jsonStr string
			key     string
		)
		jsonStr, _ = utils.Marshal(r.UserInfo)
		if jsonStr == "" {
			return
		}
		key = constant.RK_SYNC_USERS_INFO + utils.Int64ToStr(r.UserInfo.Uid)
		xredis.Set(key, jsonStr, constant.CONST_DURATION_USER_INFO_SECOND)
	}(resp)
	return
}

func (s *userService) queryUserInfo(uid int64, resp *pb_user.UserInfoResp) (err error) {
	var (
		w    = entity.NewMysqlWhere()
		user *po.User
	)
	w.Query += " AND uid = ?"
	w.Args = append(w.Args, uid)
	user, err = s.userRepo.UserInfo(w)
	if err != nil {
		setUserInfoResp(resp, ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.UserInfo, user)
	return
}

func (s *userService) queryUserAvatar(uid int64, resp *pb_user.UserInfoResp) (err error) {
	var (
		w      = entity.NewMysqlWhere()
		avatar *po.Avatar
	)
	w.Query += " AND uid = ?"
	w.Args = append(w.Args, uid)
	avatar, err = s.avatarRepo.Avatar(w)
	if err != nil {
		setUserInfoResp(resp, ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.UserInfo.Avatar, avatar)
	return
}
