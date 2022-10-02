package main

import (
	_ "lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/server"
	"lark/pkg/commands"
)

func main() {
	commands.Run(server.NewServer())
}
