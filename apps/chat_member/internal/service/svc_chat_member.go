package service

import (
	"context"
	"github.com/Shopify/sarama"
	"lark/apps/chat_member/internal/config"
	user_client "lark/apps/user/client"
	"lark/domain/repos"
	"lark/pkg/common/xkafka"
	"lark/pkg/global"
	"lark/pkg/proto/pb_chat_member"
)

type ChatMemberService interface {
	GetChatMemberUidList(ctx context.Context, req *pb_chat_member.GetChatMemberUidListReq) (resp *pb_chat_member.GetChatMemberUidListResp, err error)
	GetChatMemberSetting(ctx context.Context, req *pb_chat_member.GetChatMemberSettingReq) (resp *pb_chat_member.GetChatMemberSettingResp, err error)
	GetChatMemberInfo(ctx context.Context, req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp, err error)
	ChatMemberVerify(ctx context.Context, req *pb_chat_member.ChatMemberVerifyReq) (resp *pb_chat_member.ChatMemberVerifyResp, err error)
	ChatMemberOnline(ctx context.Context, req *pb_chat_member.ChatMemberOnlineReq) (resp *pb_chat_member.ChatMemberOnlineResp, err error)
	GetChatMemberPushConfigList(ctx context.Context, req *pb_chat_member.GetChatMemberPushConfigListReq) (resp *pb_chat_member.GetChatMemberPushConfigListResp, err error)
	GetChatMemberPushConfig(ctx context.Context, req *pb_chat_member.GetChatMemberPushConfigReq) (resp *pb_chat_member.GetChatMemberPushConfigResp, err error)
	GetChatMemberList(ctx context.Context, req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp, err error)
}

type chatMemberService struct {
	cfg                *config.Config
	chatMemberUserRepo repos.ChatMemberRepository
	userClient         user_client.UserClient
	consumerGroup      *xkafka.MConsumerGroup
	msgHandle          map[string]global.KafkaMessageHandler
}

func NewChatMemberService(cfg *config.Config, chatMemberUserRepo repos.ChatMemberRepository) ChatMemberService {
	userClient := user_client.NewUserClient(cfg.Etcd, cfg.UserServer, cfg.GrpcServer.Jaeger, cfg.Name)
	svc := &chatMemberService{cfg: cfg, chatMemberUserRepo: chatMemberUserRepo, userClient: userClient, msgHandle: make(map[string]global.KafkaMessageHandler)}
	svc.msgHandle[cfg.MsgConsumer.Topic[0]] = svc.MessageHandler
	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_1_0_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		cfg.MsgConsumer.Topic,
		cfg.MsgConsumer.Address,
		cfg.MsgConsumer.GroupID)
	go func(consumerGroup *xkafka.MConsumerGroup) {
		consumerGroup.RegisterHandleAndConsumer(svc)
	}(svc.consumerGroup)
	return svc
}

func (s *chatMemberService) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (s *chatMemberService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *chatMemberService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		err error
	)
	for msg := range claim.Messages() {
		if err = s.msgHandle[msg.Topic](msg.Value, string(msg.Key)); err != nil {
			break
		}
		session.MarkMessage(msg, "")
	}
	return nil
}
