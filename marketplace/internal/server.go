package internal

import (
	"context"
	inputAddNFTToSell "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/service/dal/repository/model"
)

type nftToSellIntrf interface {
	CreateNftRecordToSell(ctx context.Context, input model.NftToSell) (model.NftToSell, error)
	GetNftsToSell(ctx context.Context, input map[string]interface{}) ([]model.NftToSell, error)
	UpdateNftToSell(ctx context.Context, input model.NftToSell) (model.NftToSell, error)
	DeleteNftToSell(ctx context.Context, input model.NftToSell) (model.ExecResult, error)
}

type Service struct {
	dalNftToSell nftToSellIntrf
}

func New(sepoliaSrv nftToSellIntrf) Service {
	return Service{
		dalNftToSell: sepoliaSrv,
	}
}

func (s Service) NFTListForAWallet(ctx context.Context) ([]model.NftToSell, error) {
	return s.dalNftToSell.GetNftsToSell(ctx, nil)
}

func (s Service) UserAddNFTToSell(ctx context.Context, input inputAddNFTToSell.AddNFTToSell) (model.NftToSell, error) {
	nftToInsert := model.NftToSell{
		TokenId:         input.TokenId,
		Owner:           input.Owner,
		ContractAddress: input.ContractAddress,
		Creator:         input.Creator,
		TokenStandard:   input.TokenStandard,
		Status:          "on_sale",
	}

	return s.dalNftToSell.CreateNftRecordToSell(ctx, nftToInsert)
}
