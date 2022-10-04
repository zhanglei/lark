package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/msg_hot/internal/config"
	"lark/domain/repos"
	"lark/pkg/common/xkafka"
	"lark/pkg/global"
)

type MessageHotService interface {
}

type messageHotService struct {
	conf           *config.Config
	messageHotRepo repos.MessageHotRepository
	consumerGroup  *xkafka.MConsumerGroup
	msgHandle      map[string]global.KafkaMessageHandler
}

func NewMessageHotService(conf *config.Config, messageHotRepo repos.MessageHotRepository) MessageHotService {
	svc := &messageHotService{conf: conf, messageHotRepo: messageHotRepo}
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

func (s *messageHotService) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (s *messageHotService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *messageHotService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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
