package internal

import (
	"context"
	input "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/service/dal/repository/model"
)

type nftToSellIntrf interface {
	CreateNftRecordToSell(ctx context.Context, input model.NftToSell) (model.NftToSell, error)
	GetNftsToSell(ctx context.Context, input model.QueryNft) ([]model.NftToSell, error)
	UpdateNftToSell(ctx context.Context, input model.NftToSell) (model.NftToSell, error)
	DeleteNftToSell(ctx context.Context, input model.NftToSell) (model.ExecResult, error)
}

type Service struct {
	dalNftToSell nftToSellIntrf
}

var onSale = "on_sale"

func New(sepoliaSrv nftToSellIntrf) Service {
	return Service{
		dalNftToSell: sepoliaSrv,
	}
}

func (s Service) NFTListForAWallet(ctx context.Context, input input.GetNFTs) ([]model.NftToSell, error) {
	nftToGet := model.QueryNft{
		Limit:  input.Limit,
		Offset: input.Offset,
	}

	nftToGet.Owner = input.Owner
	nftToGet.TokenId = input.TokenId
	return s.dalNftToSell.GetNftsToSell(ctx, nftToGet)
}

func (s Service) UserAddNFTToSell(ctx context.Context, input input.AddNFTToSell) (model.NftToSell, error) {
	nftToInsert := model.NftToSell{
		TokenId:         input.TokenId,
		Owner:           input.Owner,
		ContractAddress: input.ContractAddress,
		Creator:         input.Creator,
		TokenStandard:   input.TokenStandard,
		Status:          &onSale,
	}

	return s.dalNftToSell.CreateNftRecordToSell(ctx, nftToInsert)
}
