package chat

import (
	"context"
	"lark/pkg/proto/pb_chat"
)

func (s *chatServer) NewChat(ctx context.Context, req *pb_chat.NewChatReq) (resp *pb_chat.NewChatResp, err error) {
	return s.ChatService.NewChat(ctx, req)
}
