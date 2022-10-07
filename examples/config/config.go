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
	Etcd             *conf.Etcd          `yaml:"etcd"`
	Mysql            *conf.Mysql         `yaml:"mysql"`
	Redis            *conf.Redis         `yaml:"redis"`
	PushOnlineServer *conf.GrpcServer    `yaml:"push_online_server"`
	ChatMemberServer *conf.GrpcServer    `yaml:"chat_member_server"`
	Platforms        []*conf.Platform    `yaml:"platforms"`
	MsgConsumer      *conf.KafkaConsumer `yaml:"msg_consumer"`
	Elasticsearch    *conf.Elasticsearch `yaml:"elasticsearch"`
}

var (
	config = new(Config)
)

var confFile = flag.String("cfg", "./configs/test.yaml", "config file")

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
