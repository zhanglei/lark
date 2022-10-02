package server

import (
	"lark/apps/request/internal/server/request"
	"lark/pkg/commands"
)

type server struct {
	requestServer request.RequestServer
}

func NewServer(requestServer request.RequestServer) commands.MainInstance {
	return &server{requestServer: requestServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.requestServer.Run()
}

func (s *server) Destroy() {

}
