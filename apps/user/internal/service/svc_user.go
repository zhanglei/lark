package service

import (
	"context"
	"github.com/Shopify/sarama"
	"lark/apps/user/internal/config"
	"lark/domain/repo"
	"lark/pkg/common/xkafka"
	"lark/pkg/global"
	"lark/pkg/proto/pb_user"
)

type UserService interface {
	GetUserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, err error)
	GetUserList(ctx context.Context, req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp, err error)
	GetChatUserInfo(ctx context.Context, req *pb_user.GetChatUserInfoReq) (resp *pb_user.GetChatUserInfoResp, err error)
	UserOnline(ctx context.Context, req *pb_user.UserOnlineReq) (resp *pb_user.UserOnlineResp, err error)
	SetUserAvatar(ctx context.Context, req *pb_user.SetUserAvatarReq) (resp *pb_user.SetUserAvatarResp, err error)
}

type userService struct {
	cfg            *config.Config
	userRepo       repo.UserRepository
	userAvatarRepo repo.UserAvatarRepository
	chatMemberRepo repo.ChatMemberRepository
	consumerGroup  *xkafka.MConsumerGroup
	msgHandle      map[string]global.KafkaMessageHandler
}

func NewUserService(cfg *config.Config, userRepo repo.UserRepository, userAvatarRepo repo.UserAvatarRepository, chatMemberRepo repo.ChatMemberRepository) UserService {
	svc := &userService{cfg: cfg, userRepo: userRepo, userAvatarRepo: userAvatarRepo, chatMemberRepo: chatMemberRepo, msgHandle: make(map[string]global.KafkaMessageHandler)}
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

func (s *userService) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (s *userService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *userService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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
