package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/chat_member/client"
	gw_client "lark/apps/msg_gateway/client"
	"lark/apps/push/internal/config"
	"lark/pkg/common/xkafka"
	"lark/pkg/global"
	"lark/pkg/proto/pb_enum"
)

type PushService interface {
}

type pushService struct {
	cfg                  *config.Config
	messageGatewayClient gw_client.MessageGatewayClient
	platforms            map[pb_enum.PLATFORM_TYPE]bool
	platformList         []pb_enum.PLATFORM_TYPE
	consumerGroup        *xkafka.MConsumerGroup
	msgHandle            map[string]global.KafkaMessageHandler
	chatMemberClient     chat_member_client.ChatMemberClient
}

func NewPushService(cfg *config.Config) PushService {
	chatMemberClient := chat_member_client.NewChatMemberClient(cfg.Etcd, cfg.ChatMemberServer, cfg.GrpcServer.Jaeger, cfg.Name)
	svc := &pushService{cfg: cfg, platforms: make(map[pb_enum.PLATFORM_TYPE]bool), chatMemberClient: chatMemberClient}
	for _, v := range cfg.Platforms {
		svc.platformList = append(svc.platformList, pb_enum.PLATFORM_TYPE(v.Type))
		svc.platforms[pb_enum.PLATFORM_TYPE(v.Type)] = v.OfflinePush
	}
	svc.messageGatewayClient = gw_client.NewMsgGwClient(cfg.Etcd, cfg.PushOnlineServer, cfg.GrpcServer.Jaeger, cfg.Name)

	svc.msgHandle = make(map[string]global.KafkaMessageHandler)
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

func (s *pushService) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (s *pushService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *pushService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		s.msgHandle[msg.Topic](msg.Value, string(msg.Key))
		session.MarkMessage(msg, "")
	}
	return nil
}
