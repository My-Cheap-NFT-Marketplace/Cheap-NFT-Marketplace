package internal

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	input "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/impl"
	"math/big"
)

type NftStandard interface {
	BalanceOf(ctx context.Context, privateKey *ecdsa.PrivateKey) (*big.Int, error)
	TokenOfOwnerByIndex(ctx context.Context, privateKey *ecdsa.PrivateKey, index *big.Int) (*big.Int, error)
	BuildNtfObject(ctx context.Context, tokenID *big.Int) (impl.TokenObj, error)
	PutNftOnSale(ctx context.Context, privateKey *ecdsa.PrivateKey, tokenId *big.Int) (impl.TransactionOutputObj, error)
}

type AuctionStandard interface {
	PutOrderToBuyNft(ctx context.Context, privateKey *ecdsa.PrivateKey, tokenId *big.Int, bid *big.Int) (impl.TransactionOutputObj, error)
}

type Service struct {
	nftContracts     map[string]NftStandard
	auctionContracts map[string]AuctionStandard
}

func New(nftContracts map[string]NftStandard, auctionContracts map[string]AuctionStandard) Service {
	return Service{
		nftContracts:     nftContracts,
		auctionContracts: auctionContracts,
	}
}

func (s Service) NFTListForAddress(ctx context.Context, input input.InputToGetMyNftListConverted) ([]interface{}, error) {
	nftContract, err := s.getNfContractInstance(input.NftContract)
	if err != nil {
		return nil, err
	}

	balance, err := nftContract.BalanceOf(ctx, input.PrivateKey)
	if err != nil {
		return nil, err
	}

	var tokenIdList []interface{}
	for i := big.NewInt(0); i.Cmp(balance) < 0; i.Add(i, big.NewInt(1)) {
		tokenID, err := nftContract.TokenOfOwnerByIndex(ctx, input.PrivateKey, i)
		if err != nil {
			return nil, err
		}
		tokenObj, err := nftContract.BuildNtfObject(ctx, tokenID)
		if err != nil {
			return nil, err
		}
		tokenIdList = append(tokenIdList, tokenObj)
	}

	return tokenIdList, nil
}

func (s Service) PutMyNftOnSale(ctx context.Context, input input.InputToPutNftOnSaleConverted) (interface{}, error) {
	nftContract, err := s.getNfContractInstance(input.NftContract)
	if err != nil {
		return nil, err
	}
	var trx impl.TransactionOutputObj
	trx, err = nftContract.PutNftOnSale(ctx, input.PrivateKey, input.TokenId)
	if err != nil {
		return trx, err
	}
	return trx, nil
}

func (s Service) BuyNftOnSale(ctx context.Context, input input.InputToBuyNftConverted) (interface{}, error) {
	auctionContract, err := s.getAuctionContractInstance(input.AuctionContract)
	if err != nil {
		return nil, err
	}
	resp, err := auctionContract.PutOrderToBuyNft(ctx, input.PrivateKey, input.TokenId, input.Bid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s Service) getNfContractInstance(contract string) (NftStandard, error) {
	contractInstance, ok := s.nftContracts[contract]
	if !ok {
		err := errors.New(fmt.Sprintf("nft contract %s is not registered in the list of available contracts", contract))
		return nil, err
	}
	return contractInstance, nil
}

func (s Service) getAuctionContractInstance(contract string) (AuctionStandard, error) {
	contractInstance, ok := s.auctionContracts[contract]
	if !ok {
		err := errors.New(fmt.Sprintf("auction contract %s is not registered in the list of available contracts", contract))
		return nil, err
	}
	return contractInstance, nil
}
