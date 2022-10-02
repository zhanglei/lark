package offline_push

import (
	"lark/apps/offline_push/internal/config"
	"lark/apps/offline_push/internal/service"
)

type OfflinePushServer interface {
	Run()
}

type offlinePushServer struct {
	conf               *config.Config
	offlinePushService service.OfflinePushService
}

func NewOfflinePushServer(conf *config.Config, offlinePushService service.OfflinePushService) OfflinePushServer {
	return &offlinePushServer{conf: conf, offlinePushService: offlinePushService}
}

func (s *offlinePushServer) Run() {

}
