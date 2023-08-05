package model

import (
	"crypto/ecdsa"
	"math/big"
)

type InputToGetMyNftList struct {
	PrivateKey  string `json:"privateKey"`
	NftContract string `json:"nftContract,omitempty"`
}

type InputToPutNftOnSale struct {
	PrivateKey  string `json:"privateKey,omitempty"`
	NftContract string `json:"nftContract,omitempty"`
	TokenId     string `json:"tokenId,omitempty"`
}

type RawInputToBuyNft struct {
	PrivateKey      *string `json:"privateKey"`
	AuctionContract *string `json:"auctionContract"`
	TokenId         *string `json:"tokenId"`
	Bid             *string `json:"bid"`
}

type InputToBuyNftConverted struct {
	PrivateKey      *ecdsa.PrivateKey `json:"privateKey"`
	AuctionContract string            `json:"auctionContract"`
	TokenId         *big.Int          `json:"tokenId"`
	Bid             *big.Int          `json:"bid"`
}
