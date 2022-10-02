package wallet

import (
	"context"
	"lark/pkg/proto/pb_wallet"
)

func (s *walletServer) CheckBalance(ctx context.Context, req *pb_wallet.CheckBalanceReq) (resp *pb_wallet.CheckBalanceResp, err error) {
	return s.walletService.CheckBalance(ctx, req)
}

func (s *walletServer) Exchange(ctx context.Context, req *pb_wallet.ExchangeReq) (resp *pb_wallet.ExchangeResp, err error) {
	return s.walletService.Exchange(ctx, req)
}

func (s *walletServer) Recharge(ctx context.Context, req *pb_wallet.RechargeReq) (resp *pb_wallet.RechargeResp, err error) {
	return s.walletService.Recharge(ctx, req)
}

func (s *walletServer) Transfer(ctx context.Context, req *pb_wallet.TransferReq) (resp *pb_wallet.TransferResp, err error) {
	return s.walletService.Transfer(ctx, req)
}
