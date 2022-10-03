package entity

import "strconv"

const (
	MongoCollectionMessages = "messages"
)

// 主键不能设置默认值
type Message struct {
	SrvMsgId       int64  `gorm:"column:srv_msg_id;primary_key" json:"srv_msg_id" bson:"srv_msg_id"`                       // 服务端消息号
	CliMsgId       int64  `gorm:"column:cli_msg_id;default:0;NOT NULL" json:"cli_msg_id" bson:"cli_msg_id"`                // 客户端消息号
	SenderId       int64  `gorm:"column:sender_id;default:0;NOT NULL" json:"sender_id" bson:"sender_id"`                   // 发送者uid
	ReceiverId     int64  `gorm:"column:receiver_id;default:0;NOT NULL" json:"receiver_id" bson:"receiver_id"`             // 接收者uid
	SenderPlatform int    `gorm:"column:sender_platform;default:0;NOT NULL" json:"sender_platform" bson:"sender_platform"` // 发送者平台
	ChatId         int64  `gorm:"column:chat_id;default:0;NOT NULL" json:"chat_id" bson:"chat_id"`                         // 会话ID
	ChatType       int    `gorm:"column:chat_type;default:0;NOT NULL" json:"chat_type" bson:"chat_type"`                   // 会话类型
	SeqId          int64  `gorm:"column:seq_id;default:0;NOT NULL" json:"seq_id" bson:"seq_id"`                            // 消息唯一ID
	MsgFrom        int    `gorm:"column:msg_from;default:0;NOT NULL" json:"msg_from" bson:"msg_from"`                      // 消息来源
	MsgType        int    `gorm:"column:msg_type;default:0;NOT NULL" json:"msg_type" bson:"msg_type"`                      // 消息类型
	Body           string `gorm:"column:body;NOT NULL" json:"body" bson:"body"`                                            // 消息本体
	Status         int    `gorm:"column:status;default:0;NOT NULL" json:"status" bson:"status"`                            // 消息状态
	SentTs         int64  `gorm:"column:sent_ts;default:0;NOT NULL" json:"sent_ts" bson:"sent_ts"`                         // 客户端本地发送时间
	SrvTs          int64  `gorm:"column:srv_ts;default:0;NOT NULL" json:"srv_ts" bson:"srv_ts"`                            // 服务端接收消息的时间
	UpdatedTs      int64  `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts" bson:"updated_ts"`              // 更新时间
	DeletedTs      int64  `gorm:"column:deleted_ts;default:0;NOT NULL" json:"deleted_ts" bson:"deleted_ts"`                // 删除时间
}

func (e *Message) GetSeqId() string {
	return strconv.FormatInt(e.SeqId, 10)
}

func (e *Message) GetChatId() string {
	return strconv.FormatInt(e.ChatId, 10)
}
