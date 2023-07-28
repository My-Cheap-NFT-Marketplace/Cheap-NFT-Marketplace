package internal

import "github.com/Cheap-NFT-Marketplace/wallet-information/internal/dal/model"

type WalletDetail struct {
	NonceDetails   model.NonceDetails   `json:"nonceDetails"`
	BalanceDetails model.BalanceDetails `json:"balanceDetails"`
}
