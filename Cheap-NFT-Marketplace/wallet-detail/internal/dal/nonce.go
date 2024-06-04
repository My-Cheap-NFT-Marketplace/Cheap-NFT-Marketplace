package dal

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/internal/dal/model"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
)

func (sc SepoliaConn) GetNonceAndPendingNonce(ctx context.Context, address common.Address, blockNumber *big.Int) model.NonceDetails {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	nonceDetails := &model.NonceDetails{}
	go func(wg *sync.WaitGroup, nonceDetails *model.NonceDetails) {
		defer wg.Done()
		var pendingNonceAt model.NonceDetail
		count, err := sc.conn.PendingNonceAt(ctx, address)
		if err != nil {
			pendingNonceAt.Error = err.Error()
			nonceDetails.PendingNonceAt = pendingNonceAt
			return
		}
		pendingNonceAt.Count = count
		nonceDetails.PendingNonceAt = pendingNonceAt
	}(wg, nonceDetails)

	go func(wg *sync.WaitGroup, nonceDetails *model.NonceDetails) {
		defer wg.Done()
		var nonceAt model.NonceDetail
		count, err := sc.conn.NonceAt(ctx, address, blockNumber)
		if err != nil {
			nonceAt.Error = err.Error()
			nonceDetails.NonceAt = nonceAt
			return
		}
		nonceAt.Count = count
		nonceDetails.NonceAt = nonceAt
	}(wg, nonceDetails)

	wg.Wait()
	return *nonceDetails
}
