package service

import (
	"context"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_enum"
)

func setNewGroupChatResp(resp *pb_chat.NewGroupChatResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *chatService) NewGroupChat(ctx context.Context, req *pb_chat.NewGroupChatReq) (resp *pb_chat.NewGroupChatResp, _ error) {
	resp = &pb_chat.NewGroupChatResp{Chat: &pb_chat.ChatInfo{}}
	var (
		uidCount = len(req.UidList)
		creator  *po.User
		tx       *gorm.DB
		w        = entity.NewMysqlWhere()
		chat     *po.Chat
		err      error
	)

	var (
		invitationMsg string
		index         int
		uid           int64
		invite        *po.ChatInvite
		inviteList    = make([]*po.ChatInvite, uidCount)
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
		AvatarKey:  creator.AvatarKey,
		Title:      req.Title,
		About:      req.About,
	}

	// 3 构建邀请信息
	invitationMsg = creator.Nickname + "邀请你加入" + chat.Title
	for index, uid = range req.UidList {
		invite = &po.ChatInvite{
			InviteId:      xsnowflake.NewSnowflakeID(),
			ChatId:        chat.ChatId,
			ChatType:      chat.ChatType,
			InitiatorUid:  req.CreatorUid,
			InviteeUid:    uid,
			InvitationMsg: invitationMsg,
		}
		inviteList[index] = invite
	}

	tx = xmysql.GetTX()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	// 4 chat入库
	err = s.chatRepo.TxCreate(tx, chat)
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED, err.Error())
		return
	}
	if uidCount == 0 {
		return
	}
	// 5 邀请信息入库
	err = s.chatInvite.TxNewChatInviteList(tx, inviteList)
	if err != nil {
		setNewGroupChatResp(resp, ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED, err.Error())
		return
	}
	// TODO: 邀请推送
	return
}
