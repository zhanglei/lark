package server

import (
	"lark/apps/offline_push/internal/server/offline_push"
	"lark/pkg/commands"
)

type server struct {
	offlinePushServer offline_push.OfflinePushServer
}

func NewServer(offlinePushServer offline_push.OfflinePushServer) commands.MainInstance {
	return &server{offlinePushServer: offlinePushServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.offlinePushServer.Run()
}

func (s *server) Destroy() {

}
