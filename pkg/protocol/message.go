package protocol

type ChatMessage struct {
	SrvMsgId        int64  `json:"srv_msg_id"`                                // 服务端消息号
	CliMsgId        int64  `json:"cli_msg_id" validate:"required,gte=0"`      // 客户端消息号
	SenderId        int64  `json:"sender_id" validate:"required,gte=0"`       // 发送者uid
	ReceiverId      int64  `json:"receiver_id" validate:"required,gte=0"`     // 接收者uid
	SenderPlatform  int    `json:"sender_platform" validate:"required,gte=0"` // 发送者平台
	SenderNickname  string `json:"sender_nickname"`                           // 发送者昵称
	SenderAvatarUrl string `json:"sender_avatar_url"`                         // 发送者头像
	ChatId          int64  `json:"chat_id" validate:"required,gte=0"`         // 会话ID
	ChatType        int    `json:"chat_type" validate:"required,gte=0"`       // 会话类型
	SeqId           int    `json:"seq_id"`                                    // 消息唯一ID
	MsgFrom         int    `json:"msg_from"`                                  // 消息来源
	MsgType         int    `json:"msg_type" validate:"required,gte=0"`        // 消息类型
	Body            string `json:"body"`                                      // 消息本体
	Status          int    `json:"status"`                                    // 消息状态
	SentTs          int64  `json:"sent_ts" validate:"required,gte=0"`         // 客户端本地发送时间
	SrvTs           int64  `json:"srv_ts"`                                    // 服务端接收消息的时间
}
