package dal

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/internal/dal/model"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/internal/platform"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"sync"
)

// GetBalanceAndPendingBalance
// todo improve code to avoid repeat declarations
func (sc SepoliaConn) GetBalanceAndPendingBalance(ctx context.Context, address common.Address, blockNumber *big.Int) model.BalanceDetails {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	balanceDetails := &model.BalanceDetails{}
	go func(wg *sync.WaitGroup, balanceDetails *model.BalanceDetails) {
		defer wg.Done()
		var balanceAt model.BalanceDetail
		resp, err := sc.conn.BalanceAt(ctx, address, blockNumber)
		if err != nil {
			balanceAt.Error = err.Error()
			balanceDetails.BalanceAt = balanceAt
			return
		} else {
			exp18, err := platform.ConvertToExponentDesired(resp, params.Wei)
			exp9, err := platform.ConvertToExponentDesired(resp, params.GWei)
			exp0, err := platform.ConvertToExponentDesired(resp, params.Ether)

			if err != nil {
				balanceAt.Error = err.Error()
				balanceDetails.BalanceAt = balanceAt
				return
			}
			exp18InFloat, _ := exp18.Float64()
			expo9InFloat, _ := exp9.Float64()
			expo0InFloat, _ := exp0.Float64()
			balanceAt.Denomination = model.Denomination{
				Exp18: platform.SetDecimalFormat(exp18InFloat, "0"),
				Exp9:  platform.SetDecimalFormat(expo9InFloat, "9"),
				Exp0:  platform.SetDecimalFormat(expo0InFloat, "18"),
			}
		}
		balanceDetails.BalanceAt = balanceAt
	}(wg, balanceDetails)

	go func(wg *sync.WaitGroup, balanceDetails *model.BalanceDetails) {
		defer wg.Done()
		var pendingBalanceAt model.BalanceDetail
		resp, err := sc.conn.PendingBalanceAt(ctx, address)
		if err != nil {
			pendingBalanceAt.Error = err.Error()
			balanceDetails.PendingBalanceAt = pendingBalanceAt
			return
		} else {
			exp18, err := platform.ConvertToExponentDesired(resp, params.Wei)
			exp9, err := platform.ConvertToExponentDesired(resp, params.GWei)
			exp0, err := platform.ConvertToExponentDesired(resp, params.Ether)

			if err != nil {
				pendingBalanceAt.Error = err.Error()
				balanceDetails.PendingBalanceAt = pendingBalanceAt
				return
			}
			exp18InFloat, _ := exp18.Float64()
			expo9InFloat, _ := exp9.Float64()
			expo0InFloat, _ := exp0.Float64()
			pendingBalanceAt.Denomination = model.Denomination{
				Exp18: platform.SetDecimalFormat(exp18InFloat, "0"),
				Exp9:  platform.SetDecimalFormat(expo9InFloat, "9"),
				Exp0:  platform.SetDecimalFormat(expo0InFloat, "18"),
			}
		}
		balanceDetails.PendingBalanceAt = pendingBalanceAt
	}(wg, balanceDetails)

	wg.Wait()
	return *balanceDetails
}
