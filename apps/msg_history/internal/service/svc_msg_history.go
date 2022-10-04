package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/msg_history/internal/config"
	"lark/domain/repo"
	"lark/pkg/common/xkafka"
	"lark/pkg/global"
)

type MessageHistoryService interface {
}

type messageHistoryService struct {
	conf               *config.Config
	messageHistoryRepo repo.MessageHistoryRepository
	consumerGroup      *xkafka.MConsumerGroup
	msgHandle          map[string]global.KafkaMessageHandler
}

func NewMessageHistoryService(conf *config.Config, messageHistoryRepo repo.MessageHistoryRepository) MessageHistoryService {
	svc := &messageHistoryService{conf: conf, messageHistoryRepo: messageHistoryRepo}

	svc.msgHandle = make(map[string]global.KafkaMessageHandler)
	svc.msgHandle[conf.MsgConsumer.Topic[0]] = svc.MessageHandler
	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_1_0_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		conf.MsgConsumer.Topic,
		conf.MsgConsumer.Address,
		conf.MsgConsumer.GroupID)
	go func(consumerGroup *xkafka.MConsumerGroup) {
		consumerGroup.RegisterHandleAndConsumer(svc)
	}(svc.consumerGroup)
	return svc
}

func (s *messageHistoryService) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (s *messageHistoryService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *messageHistoryService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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
