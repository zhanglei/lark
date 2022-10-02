package main

import (
	"lark/examples/chat/client"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	manager := client.NewManager()
	manager.Run()

	wg.Wait()
}
