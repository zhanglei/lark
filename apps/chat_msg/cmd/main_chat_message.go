package main

import (
	"lark/apps/chat_msg/dig"
	"lark/apps/chat_msg/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xmongo"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xmongo.NewMongoClient(conf.Mongo)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
