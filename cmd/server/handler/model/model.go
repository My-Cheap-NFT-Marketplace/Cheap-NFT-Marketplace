package model

import "math/big"

type InputWalletDetail struct {
	Wallet Detail `json:"wallet"`
}

type Detail struct {
	Address     string   `json:"address"`
	BlockNumber *big.Int `json:"blockNumber"`
}
