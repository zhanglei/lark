package chat

import (
	"context"
	"lark/pkg/proto/pb_chat"
)

func (s *chatServer) NewGroupChat(ctx context.Context, req *pb_chat.NewGroupChatReq) (resp *pb_chat.NewGroupChatResp, err error) {
	return s.ChatService.NewGroupChat(ctx, req)
}

func (s *chatServer) SetGroupChat(ctx context.Context, req *pb_chat.SetGroupChatReq) (resp *pb_chat.SetGroupChatResp, err error) {
	return s.ChatService.SetGroupChat(ctx, req)
}
