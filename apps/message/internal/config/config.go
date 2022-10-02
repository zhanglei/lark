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
	GrpcServer       *conf.Grpc          `yaml:"grpc_server"`
	ChatMemberServer *conf.GrpcServer    `yaml:"chat_member_server"`
	Etcd             *conf.Etcd          `yaml:"etcd"`
	Redis            *conf.Redis         `yaml:"redis"`
	MsgProducer      *conf.KafkaProducer `yaml:"msg_producer"`
	Jaeger           *conf.Jaeger        `yaml:"jaeger"`
}

var (
	config = new(Config)
)

var confFile = flag.String("cfg", "./configs/message.yaml", "config file")

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
