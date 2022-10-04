package dos

import (
	"lark/pkg/proto/pb_enum"
)

type ChatMemberDo struct {
	Uid       int64                   `json:"uid"` // Uid
	Platforms []pb_enum.PLATFORM_TYPE `json:"platforms"`
	ServerId  int32                   `json:"server_id"`
}
