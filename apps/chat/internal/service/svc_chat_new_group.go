package service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
)

func setNewGroupChatResp(resp *pb_chat.NewGroupChatResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatService) NewGroupChat(ctx context.Context, req *pb_chat.NewGroupChatReq) (resp *pb_chat.NewGroupChatResp, _ error) {
	resp = &pb_chat.NewGroupChatResp{Chat: &pb_chat.ChatInfo{}}
	var (
		creator *po.User
		tx      *gorm.DB
		w       = entity.NewMysqlWhere()
		chat    *po.Chat
		err     error
	)
	var (
		avatar        *po.Avatar
		member        *po.ChatMember
		invitationMsg string
		uid           int64
		invite        *po.ChatInvite
		inviteList    = make([]*po.ChatInvite, 0)
		pushKey       string
	)

	// 1 获取创建者信息
	w.SetFilter("uid=?", req.CreatorUid)
	creator, err = s.userRepo.UserInfo(w)
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, err.Error())
		return
	}

	// 2 构建chat模型
	chat = &po.Chat{
		CreatorUid: req.CreatorUid,
		ChatType:   int(pb_enum.CHAT_TYPE_GROUP),
		AvatarKey:  constant.CONST_AVATAR_KEY_SMALL,
		Title:      req.Title,
		About:      req.About,
	}
	tx = xmysql.GetTX()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	// 3 chat入库
	err = s.chatRepo.TxCreate(tx, chat)
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED, err.Error())
		return
	}

	// 4 creator入群/入库
	member = &po.ChatMember{
		ChatId:          chat.ChatId,
		ChatType:        chat.ChatType,
		Uid:             creator.Uid,
		RoleId:          1, // 超级管理员
		DisplayName:     creator.Nickname,
		MemberAvatarKey: creator.AvatarKey,
		ChatAvatarKey:   chat.AvatarKey,
		Sync:            1,
		Platform:        creator.Platform,
		ServerId:        creator.ServerId,
	}
	err = s.chatMemberRepo.TxCreate(tx, member)
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED, err.Error())
		return
	}

	// 5 缓存成员信息
	pushKey = constant.RK_SYNC_CHAT_MEMBERS_PUSH_MEMBER_HASH + utils.Int64ToStr(member.ChatId)
	err = xredis.HSetNX(pushKey, utils.Int64ToStr(member.Uid), fmt.Sprintf("%d,%d,%d,%d", member.Uid, member.Platform, member.ServerId, member.Mute))
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_CACHE_CHAT_MEMBER_FAILED, ERROR_CHAT_CACHE_CHAT_MEMBER_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_CACHE_CHAT_MEMBER_FAILED, ERROR_CHAT_CACHE_CHAT_MEMBER_FAILED, err.Error())
		return
	}

	// 6 设置群头像
	avatar = &po.Avatar{
		OwnerId:      chat.ChatId,
		OwnerType:    int(pb_enum.AVATAR_OWNER_CHAT_AVATAR),
		AvatarSmall:  constant.CONST_AVATAR_KEY_SMALL,
		AvatarMedium: constant.CONST_AVATAR_KEY_MEDIUM,
		AvatarLarge:  constant.CONST_AVATAR_KEY_LARGE,
	}
	err = s.avatarRepo.TxCreate(tx, avatar)
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED, err.Error())
		return
	}

	// 7 构建邀请信息
	invitationMsg = creator.Nickname + CONST_CHAT_INVITE_TITLE_CONJUNCTION + chat.Title
	for _, uid = range req.UidList {
		if uid == req.CreatorUid {
			continue
		}
		invite = &po.ChatInvite{
			InviteId:      xsnowflake.NewSnowflakeID(),
			ChatId:        chat.ChatId,
			ChatType:      chat.ChatType,
			InitiatorUid:  req.CreatorUid,
			InviteeUid:    uid,
			InvitationMsg: invitationMsg,
		}
		inviteList = append(inviteList, invite)
	}
	if len(inviteList) == 0 {
		return
	}

	// 8 邀请信息入库
	err = s.chatInviteRepo.TxNewChatInviteList(tx, inviteList)
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED, err.Error())
		return
	}
	// TODO: 邀请推送
	return
}
