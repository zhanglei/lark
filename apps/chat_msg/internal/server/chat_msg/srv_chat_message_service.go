package chat_msg

import (
	"context"
	"lark/pkg/proto/pb_chat_msg"
)

func (s *chatMessageServer) GetChatMessages(ctx context.Context, req *pb_chat_msg.GetChatMessagesReq) (resp *pb_chat_msg.GetChatMessagesResp, err error) {
	return s.chatMessageService.GetChatMessages(ctx, req)
}
