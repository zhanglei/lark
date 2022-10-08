package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	chat_member_client "lark/apps/chat_member/client"
	"lark/apps/message/internal/config"
	kafka "lark/pkg/common/xkafka"
	"lark/pkg/proto/pb_msg"
)

type MessageService interface {
	SendChatMessage(ctx context.Context, req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp, _ error)
	MessageOperation(ctx context.Context, req *pb_msg.MessageOperationReq) (resp *pb_msg.MessageOperationResp, err error)
}

type messageService struct {
	cfg              *config.Config
	validate         *validator.Validate
	producer         *kafka.Producer
	chatMemberClient chat_member_client.ChatMemberClient
}

func NewMessageService(cfg *config.Config) MessageService {
	chatMemberClient := chat_member_client.NewChatMemberClient(cfg.Etcd, cfg.ChatMemberServer, cfg.GrpcServer.Jaeger, cfg.Name)
	svc := &messageService{cfg: cfg, validate: validator.New(), chatMemberClient: chatMemberClient}
	svc.producer = kafka.NewKafkaProducer(cfg.MsgProducer.Address, cfg.MsgProducer.Topic)
	return svc
}
