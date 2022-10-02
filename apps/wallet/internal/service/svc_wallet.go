package service

import (
	"context"
	"lark/pkg/proto/pb_wallet"
)

type WalletService interface {
	CheckBalance(ctx context.Context, req *pb_wallet.CheckBalanceReq) (resp *pb_wallet.CheckBalanceResp, err error)
	Exchange(ctx context.Context, req *pb_wallet.ExchangeReq) (resp *pb_wallet.ExchangeResp, err error)
	Recharge(ctx context.Context, req *pb_wallet.RechargeReq) (resp *pb_wallet.RechargeResp, err error)
	Transfer(ctx context.Context, req *pb_wallet.TransferReq) (resp *pb_wallet.TransferResp, err error)
}

type walletService struct {
}

func NewWalletService() WalletService {
	return &walletService{}
}
