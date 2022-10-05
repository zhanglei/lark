package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
)

func setAvatarResp(resp *pb_user.SetAvatarResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *userService) SetAvatar(ctx context.Context, req *pb_user.SetAvatarReq) (resp *pb_user.SetAvatarResp, _ error) {
	resp = &pb_user.SetAvatarResp{Avatar: &pb_user.UserAvatar{}}
	var (
		avatar = new(po.Avatar)
		u      = entity.NewMysqlUpdate()
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
	err = s.avatarRepo.TxSaveAvatar(tx, avatar)
	if err != nil {
		setAvatarResp(resp, ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
		xlog.Warn(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED, err.Error())
		return
	}

	switch req.OwnerType {
	case pb_enum.AVATAR_OWNER_USER_AVATAR:
		u.AndQuery("sync=1")
		u.SetFilter("uid=?", req.OwnerId)
		u.Set("member_avatar_key", avatar.AvatarSmall)
	case pb_enum.AVATAR_OWNER_CHAT_AVATAR:
		u.SetFilter("chat_id=?", req.OwnerId)
		u.Set("chat_avatar_key", avatar.AvatarSmall)
	}

	err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
	if err != nil {
		setAvatarResp(resp, ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
		xlog.Warn(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED, err.Error())
		return
	}
	copier.Copy(resp.Avatar, avatar)
	return
}
