package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name        string              `yaml:"name"`
	Log         string              `yaml:"log"`
	Etcd        *conf.Etcd          `yaml:"etcd"`
	Mysql       *conf.Mysql         `yaml:"mysql"`
	Redis       *conf.Redis         `yaml:"redis"`
	MsgConsumer *conf.KafkaConsumer `yaml:"msg_consumer"`
}

var (
	config = new(Config)
)

var confFile = flag.String("cfg", "./configs/msg_history.yaml", "config file")

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
