package po

import "lark/pkg/entity"

type ChatInvite struct {
	entity.GormEntityTs
	InviteId      int64  `gorm:"column:invite_id;primary_key" json:"invite_id"`                // invite ID
	InvitedTs     int64  `gorm:"column:invited_ts;autoCreateTime:milli" json:"invited_ts"`     // 邀请时间
	ChatId        int64  `gorm:"column:chat_id;default:0" json:"chat_id"`                      // Chat ID
	ChatType      int    `gorm:"column:chat_type;default:0" json:"chat_type"`                  // 1:私聊/2:群聊
	InitiatorUid  int64  `gorm:"column:initiator_uid;default:0;NOT NULL" json:"initiator_uid"` // 发起人 UID
	InviteeUid    int64  `gorm:"column:invitee_uid;default:0;NOT NULL" json:"invitee_uid"`     // 被邀请人UID/群ID
	InvitationMsg string `gorm:"column:invitation_msg;NOT NULL" json:"invitation_msg"`         // 邀请消息
	HandlerUid    int64  `gorm:"column:handler_uid;default:0" json:"handler_uid"`              // 处理人 UID
	HandleResult  int    `gorm:"column:handle_result;default:0" json:"handle_result"`          // 结果
	HandleMsg     string `gorm:"column:handle_msg" json:"handle_msg"`                          // 处理消息
	HandledTs     int64  `gorm:"column:handled_ts;default:0" json:"handled_ts"`                // 处理时间
}
