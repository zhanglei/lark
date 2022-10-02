package server

import (
	"lark/apps/wallet/internal/server/wallet"
	"lark/pkg/commands"
)

type server struct {
	walletServer wallet.WalletServer
}

func NewServer(walletServer wallet.WalletServer) commands.MainInstance {
	return &server{walletServer: walletServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.walletServer.Run()
}

func (s *server) Destroy() {

}
