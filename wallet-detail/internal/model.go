package internal

import "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/internal/dal/model"

type WalletDetail struct {
	NonceDetails   model.NonceDetails   `json:"nonceDetails"`
	BalanceDetails model.BalanceDetails `json:"balanceDetails"`
}
