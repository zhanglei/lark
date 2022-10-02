package push

import (
	"lark/apps/push/internal/config"
	"lark/apps/push/internal/service"
	"lark/pkg/common/xmonitor"
	"lark/pkg/proto/pb_push"
)

type PushServer interface {
	Run()
}

type pushServer struct {
	pb_push.UnimplementedPushServer
	cfg         *config.Config
	pushService service.PushService
}

func NewPushServer(cfg *config.Config, pushService service.PushService) PushServer {
	srv := &pushServer{cfg: cfg, pushService: pushService}
	return srv
}

func (s *pushServer) Run() {
	xmonitor.RunMonitor(s.cfg.Monitor.Port)
}
