package request

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/request/internal/config"
	"lark/apps/request/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_req"
)

type RequestServer interface {
	Run()
}

type requestServer struct {
	pb_req.UnimplementedRequestServer
	cfg            *config.Config
	grpcServer     *xgrpc.GrpcServer
	requestService service.RequestService
}

func NewRequestServer(cfg *config.Config, requestService service.RequestService) RequestServer {
	return &requestServer{cfg: cfg, requestService: requestService}
}

func (s *requestServer) Run() {
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	srv, closer = xgrpc.NewServer(s.cfg.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_req.RegisterRequestServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
