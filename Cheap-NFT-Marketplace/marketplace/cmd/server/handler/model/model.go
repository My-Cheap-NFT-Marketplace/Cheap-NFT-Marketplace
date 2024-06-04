package model

import (
	"crypto/ecdsa"
	"math/big"
)

type RawInputToGetMyNftList struct {
	PrivateKey  *string `json:"privateKey"`
	NftContract *string `json:"nftContract,omitempty"`
}

type InputToGetMyNftListConverted struct {
	PrivateKey  *ecdsa.PrivateKey `json:"privateKey"`
	NftContract string            `json:"nftContract,omitempty"`
}

type RawInputToPutNftOnSale struct {
	PrivateKey  *string `json:"privateKey"`
	NftContract *string `json:"nftContract"`
	TokenId     *string `json:"tokenId"`
}

type InputToPutNftOnSaleConverted struct {
	PrivateKey  *ecdsa.PrivateKey `json:"privateKey"`
	NftContract string            `json:"nftContract"`
	TokenId     *big.Int          `json:"tokenId"`
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
