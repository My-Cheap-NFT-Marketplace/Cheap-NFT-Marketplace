package internal

import (
	"context"
	"errors"
	"fmt"
	input "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/impl"
	"math/big"
)

type ERC721Standard interface {
	BalanceOf(ctx context.Context, privateKey string) (*big.Int, error)
	TokenOfOwnerByIndex(ctx context.Context, privateKey string, index *big.Int) (*big.Int, error)
	BuildNtfObject(ctx context.Context, tokenID *big.Int) (impl.TokenObj, error)
	PutNftOnSale(ctx context.Context, privateKey string, tokenId *big.Int) (impl.TransactionOutputObj, error)
}

type Service struct {
	contracts map[string]ERC721Standard
}

func New(contracts map[string]ERC721Standard) Service {
	return Service{
		contracts: contracts,
	}
}

func (s Service) NFTListForAddress(ctx context.Context, input input.InputToGetMyNftList) ([]interface{}, error) {
	contract, err := s.getContractInstance(input.NftContract)
	if err != nil {
		return nil, err
	}

	balance, err := contract.BalanceOf(ctx, input.PrivateKey)
	if err != nil {
		return nil, err
	}

	var tokenIdList []interface{}
	for i := big.NewInt(0); i.Cmp(balance) < 0; i.Add(i, big.NewInt(1)) {
		tokenID, err := contract.TokenOfOwnerByIndex(ctx, input.PrivateKey, i)
		if err != nil {
			return nil, err
		}
		tokenObj, err := contract.BuildNtfObject(ctx, tokenID)
		if err != nil {
			return nil, err
		}
		tokenIdList = append(tokenIdList, tokenObj)
	}

	return tokenIdList, nil
}

func (s Service) PutMyNftOnSale(ctx context.Context, input input.InputToPutNftOnSale) (interface{}, error) {
	contract, err := s.getContractInstance(input.NftContract)
	if err != nil {
		return nil, err
	}
	var trx impl.TransactionOutputObj
	tokenId := new(big.Int)
	tokenId.SetString(input.TokenId, 10)
	trx, err = contract.PutNftOnSale(ctx, input.PrivateKey, tokenId)
	if err != nil {
		return trx, err
	}
	return trx, nil
}

func (s Service) getContractInstance(contract string) (ERC721Standard, error) {
	contractInstance, ok := s.contracts[contract]
	if !ok {
		err := errors.New(fmt.Sprintf("contract %s is not registered in the list of available contracts", contract))
		return nil, err
	}
	return contractInstance, nil
}
