package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/offline_push/internal/config"
	"lark/pkg/common/xkafka"
	"lark/pkg/global"
)

type OfflinePushService interface {
}

type offlinePushService struct {
	cfg           *config.Config
	consumerGroup *xkafka.MConsumerGroup
	msgHandle     map[string]global.KafkaMessageHandler
	msgCount      int64
}

func NewOfflinePushService(cfg *config.Config) OfflinePushService {
	svc := &offlinePushService{cfg: cfg}
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

func (s *offlinePushService) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (s *offlinePushService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *offlinePushService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		s.msgHandle[msg.Topic](msg.Value, string(msg.Key))
		session.MarkMessage(msg, "")
	}
	return nil
}
