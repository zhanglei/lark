package server

import (
	"lark/apps/push/internal/server/push"
	"lark/pkg/commands"
)

type server struct {
	pushServer push.PushServer
}

func NewServer(pushServer push.PushServer) commands.MainInstance {
	return &server{pushServer: pushServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.pushServer.Run()
}

func (s *server) Destroy() {

}
