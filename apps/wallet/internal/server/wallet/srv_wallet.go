package wallet

import (
	"lark/apps/wallet/internal/config"
	"lark/apps/wallet/internal/service"
	"lark/pkg/common/xgrpc"
)

type WalletServer interface {
	Run()
}

type walletServer struct {
	cfg           *config.Config
	grpcServer    *xgrpc.GrpcServer
	walletService service.WalletService
}

func NewWalletServer(cfg *config.Config, walletService service.WalletService) WalletServer {
	return &walletServer{cfg: cfg, walletService: walletService}
}

func (s *walletServer) Run() {

}
