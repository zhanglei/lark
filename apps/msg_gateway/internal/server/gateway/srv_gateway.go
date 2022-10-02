package gateway

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/msg_gateway/internal/config"
	"lark/apps/msg_gateway/internal/server/websocket/ws"
	"lark/apps/msg_gateway/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xkafka"
	"lark/pkg/common/xmonitor"
	"lark/pkg/proto/pb_gw"
)

type GatewayServer interface {
	Run()
}

type gatewayServer struct {
	pb_gw.UnimplementedMessageGatewayServer
	conf       *config.Config
	wsServer   *ws.WServer
	wsService  service.WsService
	grpcServer *xgrpc.GrpcServer
	producer   *xkafka.Producer
}

func NewGatewayServer(conf *config.Config, wsService service.WsService) GatewayServer {
	srv := &gatewayServer{conf: conf, wsService: wsService}
	srv.wsServer = ws.NewWServer(conf.WsServer.Port, conf.GrpcServer.ServerID, wsService.MessageCallback)
	srv.producer = xkafka.NewKafkaProducer(conf.MsgProducer.Address, conf.MsgProducer.Topic)
	return srv
}

func (s *gatewayServer) Run() {
	xmonitor.RunMonitor(s.conf.Monitor.Port)
	go s.wsServer.Run()
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	srv, closer = xgrpc.NewServer(s.conf.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_gw.RegisterMessageGatewayServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.conf.GrpcServer, s.conf.Etcd)
	s.grpcServer.RunServer(srv)
}
