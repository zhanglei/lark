package main

import (
	"lark/examples/chat/client"
	"lark/examples/config"
	"lark/pkg/common/xredis"
	"sync"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	manager := client.NewManager()
	manager.Run()

	wg.Wait()
}
