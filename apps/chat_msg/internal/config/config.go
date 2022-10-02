package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name       string         `yaml:"name"`
	Log        string         `yaml:"log"`
	GrpcServer *conf.Grpc     `yaml:"grpc_server"`
	Etcd       *conf.Etcd     `yaml:"etcd"`
	Mysql      *conf.Mysql    `yaml:"mysql"`
	Mongo      *conf.Mongo    `yaml:"mongo"`
	Redis      *conf.Redis    `yaml:"redis"`
	MsgCache   *conf.MsgCache `yaml:"msg_cache"`
}

var (
	config = new(Config)
)

var confFile = flag.String("cfg", "./configs/chat_msg.yaml", "config file")

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
