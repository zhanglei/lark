package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func setUserAvatarResp(resp *pb_user.SetUserAvatarResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *userService) SetUserAvatar(ctx context.Context, req *pb_user.SetUserAvatarReq) (resp *pb_user.SetUserAvatarResp, _ error) {
	resp = &pb_user.SetUserAvatarResp{Avatar: &pb_user.UserAvatar{}}
	var (
		avatar = new(po.UserAvatar)
		err    error
	)
	copier.Copy(avatar, req)
	tx := xmysql.GetTX()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = s.userAvatarRepo.TxSaveAvatar(tx, avatar)
	if err != nil {
		setUserAvatarResp(resp, ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
		xlog.Warn(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED, err.Error())
		return
	}
	u := entity.NewMysqlUpdate()
	u.AndQuery("uid=?")
	u.AppendArg(avatar.Uid)
	u.AndQuery("sync=1")
	u.Set("avatar_key", avatar.AvatarSmall)
	err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
	if err != nil {
		setUserAvatarResp(resp, ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
		xlog.Warn(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED, err.Error())
		return
	}
	copier.Copy(resp.Avatar, avatar)
	return
}
