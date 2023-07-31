package internal

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/cmd/server/handler/model"
	dalModel "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/internal/service/dal/model"
)

type SepoliaSrv interface {
	ExecSupplyMechanismsToAddMockERC20ToAccount(ctx context.Context, input model.AddTokenMockERC20ToAddress) (dalModel.TransactionOutput, error)
	ExecSupplyMechanismsToAddMockERC721ToAccount(ctx context.Context, input model.AddTokenMockERCM721ToAddress) (dalModel.TransactionOutput, error)
	ExecGetBalanceForAccount(ctx context.Context, input model.GetBalanceForAddress) (interface{}, error)
}

type Service struct {
	sepoliaSrv SepoliaSrv
}

func New(sepoliaSrv SepoliaSrv) Service {
	return Service{
		sepoliaSrv: sepoliaSrv,
	}
}

func (s Service) GetSupplyMockERC20ToAccount(ctx context.Context, input model.AddTokenMockERC20ToAddress) (dalModel.TransactionOutput, error) {
	return s.sepoliaSrv.ExecSupplyMechanismsToAddMockERC20ToAccount(ctx, input)
}

func (s Service) GetSupplyNftMockERC721ToAccount(ctx context.Context, input model.AddTokenMockERCM721ToAddress) (dalModel.TransactionOutput, error) {
	return s.sepoliaSrv.ExecSupplyMechanismsToAddMockERC721ToAccount(ctx, input)
}

func (s Service) GetBalanceForAccount(ctx context.Context, input model.GetBalanceForAddress) (interface{}, error) {
	return s.sepoliaSrv.ExecGetBalanceForAccount(ctx, input)
}
