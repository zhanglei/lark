package svc_auth

import (
	auth_client "lark/apps/auth/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_auth"
	link_client "lark/apps/link/client"
	kafka "lark/pkg/common/xkafka"
	"lark/pkg/xhttp"
)

type AuthService interface {
	Register(params *dto_auth.RegisterReq) (resp *xhttp.Resp)
	Login(params *dto_auth.LoginReq) (resp *xhttp.Resp)
}

type authService struct {
	authClient auth_client.AuthClient
	linkClient link_client.LinkClient
	producer   *kafka.Producer
}

func NewAuthService() AuthService {
	conf := config.GetConfig()
	authClient := auth_client.NewAuthClient(conf.Etcd, conf.AuthServer, conf.Jaeger, conf.Name)
	linkClient := link_client.NewLinkClient(conf.Etcd, conf.LinkServer, conf.Jaeger, conf.Name)
	producer := kafka.NewKafkaProducer(conf.MsgProducer.Address, conf.MsgProducer.Topic)
	return &authService{authClient: authClient, linkClient: linkClient, producer: producer}
}
