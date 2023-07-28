package internal

import (
	"context"
	"github.com/Cheap-NFT-Marketplace/wallet-information/internal/dal/model"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
)

type SepoliaSrv interface {
	GetBalanceAndPendingBalance(ctx context.Context, address common.Address, blockNumber *big.Int) model.BalanceDetails
	GetNonceAndPendingNonce(ctx context.Context, address common.Address, blockNumber *big.Int) model.NonceDetails
}

type Service struct {
	sepoliaSrv SepoliaSrv
}

func New(sepoliaSrv SepoliaSrv) Service {
	return Service{
		sepoliaSrv: sepoliaSrv,
	}
}

func (s Service) GetWalletDetail(ctx context.Context, address string, blockNumber *big.Int) WalletDetail {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	addressInByte := common.HexToAddress(address)
	walletDetail := &WalletDetail{}
	go func(wg *sync.WaitGroup, walletDetail *WalletDetail) {
		defer wg.Done()
		resp := s.sepoliaSrv.GetNonceAndPendingNonce(ctx, addressInByte, blockNumber)
		walletDetail.NonceDetails = resp
	}(wg, walletDetail)

	go func(wg *sync.WaitGroup, walletDetail *WalletDetail) {
		defer wg.Done()
		resp := s.sepoliaSrv.GetBalanceAndPendingBalance(ctx, addressInByte, blockNumber)
		walletDetail.BalanceDetails = resp
	}(wg, walletDetail)

	wg.Wait()

	return *walletDetail
}
