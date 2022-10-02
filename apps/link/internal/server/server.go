package server

import (
	"lark/apps/link/internal/server/link"
	"lark/pkg/commands"
)

type server struct {
	linkServer link.LinkServer
}

func NewServer(linkServer link.LinkServer) commands.MainInstance {
	return &server{linkServer: linkServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.linkServer.Run()
}

func (s *server) Destroy() {

}
