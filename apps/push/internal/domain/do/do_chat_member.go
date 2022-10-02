package do

import "lark/pkg/proto/pb_msg"

type ChatMemberDo struct {
	Uid       int64                  `json:"uid"` // Uid
	Platforms []pb_msg.PLATFORM_TYPE `json:"platforms"`
	ServerId  int32                  `json:"server_id"`
}
