package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name             string              `yaml:"name"`
	Port             int                 `yaml:"port"`
	Log              string              `yaml:"log"`
	Etcd             *conf.Etcd          `yaml:"etcd"`
	Redis            *conf.Redis         `yaml:"redis"`
	AuthServer       *conf.GrpcServer    `yaml:"auth_server"`
	UserServer       *conf.GrpcServer    `yaml:"user_server"`
	ChatMsgServer    *conf.GrpcServer    `yaml:"chat_msg_server"`
	LinkServer       *conf.GrpcServer    `yaml:"link_server"`
	ChatMemberServer *conf.GrpcServer    `yaml:"chat_member_server"`
	ChatInviteServer *conf.GrpcServer    `yaml:"chat_invite_server"`
	ChatServer       *conf.GrpcServer    `yaml:"chat_server"`
	AvatarServer     *conf.GrpcServer    `yaml:"avatar_server"`
	MsgProducer      *conf.KafkaProducer `yaml:"msg_producer"`
	Minio            *conf.Minio         `yaml:"minio"`
	Jaeger           *conf.Jaeger        `yaml:"jaeger"`
}

var (
	config = new(Config)
)

var confFile = flag.String("cfg", "./configs/api_gateway.yaml", "api_gateway config")

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)
	xlog.Shared(config.Log, config.Name)
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
