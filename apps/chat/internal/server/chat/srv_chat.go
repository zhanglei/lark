package chat

import (
	"lark/apps/chat/internal/config"
	"lark/apps/chat/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_chat"
)

type ChatServer interface {
	Run()
}

type chatServer struct {
	pb_chat.UnimplementedChatServer
	cfg         *config.Config
	grpcServer  *xgrpc.GrpcServer
	ChatService service.ChatService
}

func NewChatServer(cfg *config.Config, ChatService service.ChatService) ChatServer {
	return &chatServer{cfg: cfg, ChatService: ChatService}
}

func (s *chatServer) Run() {

}
