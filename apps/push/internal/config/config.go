package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name             string              `yaml:"name"`
	Log              string              `yaml:"log"`
	Monitor          *conf.Monitor       `yaml:"monitor"`
	GrpcServer       *conf.Grpc          `yaml:"grpc_server"`
	PushOnlineServer *conf.GrpcServer    `yaml:"push_online_server"`
	ChatMemberServer *conf.GrpcServer    `yaml:"chat_member_server"`
	Etcd             *conf.Etcd          `yaml:"etcd"`
	Platforms        []*conf.Platform    `yaml:"platforms"`
	MsgConsumer      *conf.KafkaConsumer `yaml:"msg_consumer"`
	Redis            *conf.Redis         `yaml:"redis"`
}

var (
	config = new(Config)
)

var confFile = flag.String("cfg", "./configs/push.yaml", "config file")

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
