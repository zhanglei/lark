package po

import "lark/pkg/entity"

type ChatRequest struct {
	entity.GormUpdatedTs
	entity.GormDeletedTs
	RequestId    int64  `gorm:"column:request_id;primary_key" json:"request_id"`              // request ID
	RequestTs    int64  `gorm:"column:request_ts;autoCreateTime:milli" json:"request_ts"`     // 申请时间
	ChatType     int    `gorm:"column:chat_type;default:0" json:"chat_type"`                  // 1:私聊/2:群聊
	InitiatorUid int64  `gorm:"column:initiator_uid;default:0;NOT NULL" json:"initiator_uid"` // 发起人 UID
	TargetId     int64  `gorm:"column:target_id;default:0;NOT NULL" json:"target_id"`         // 被邀请人UID/群ID
	RequestMsg   string `gorm:"column:request_msg;NOT NULL" json:"request_msg"`               // request消息
	HandleResult int    `gorm:"column:handle_result;default:0" json:"handle_result"`          // 结果
	HandleMsg    string `gorm:"column:handle_msg" json:"handle_msg"`                          // 处理消息
	HandledTs    int64  `gorm:"column:handled_ts;default:0" json:"handled_ts"`                // 处理时间
}

type ChatUser struct {
	entity.GormEntityTs
	ChatId int64 `gorm:"column:chat_id;primary_key" json:"chat_id"`      // chat ID
	Uid    int64 `gorm:"column:uid;NOT NULL" json:"uid"`                 // 用户ID
	Status int   `gorm:"column:status;default:0;NOT NULL" json:"status"` // chat状态
}
