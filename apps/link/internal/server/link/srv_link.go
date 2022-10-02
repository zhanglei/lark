package link

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/link/internal/config"
	"lark/apps/link/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_link"
)

type LinkServer interface {
	Run()
}

type linkServer struct {
	pb_link.UnimplementedLinkServer
	cfg         *config.Config
	grpcServer  *xgrpc.GrpcServer
	linkService service.LinkService
}

func NewLinkServer(cfg *config.Config, linkService service.LinkService) LinkServer {
	return &linkServer{cfg: cfg, linkService: linkService}
}

func (s *linkServer) Run() {
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

	pb_link.RegisterLinkServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
