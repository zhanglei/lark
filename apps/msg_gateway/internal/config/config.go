package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name          string              `yaml:"name"`
	Log           string              `yaml:"log"`
	Monitor       *conf.Monitor       `yaml:"monitor"`
	GrpcServer    *conf.Grpc          `yaml:"grpc_server"`
	MessageServer *conf.GrpcServer    `yaml:"message_server"`
	Etcd          *conf.Etcd          `yaml:"etcd"`
	WsServer      *conf.WsServer      `yaml:"ws_server"`
	Redis         *conf.Redis         `yaml:"redis"`
	MsgProducer   *conf.KafkaProducer `yaml:"msg_producer"`
}

var (
	config = new(Config)
)

var confFile = flag.String("cfg", "./configs/msg_gateway.yaml", "config file")

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
