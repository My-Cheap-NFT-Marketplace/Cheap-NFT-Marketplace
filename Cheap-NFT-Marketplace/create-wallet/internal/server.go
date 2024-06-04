package internal

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/internal/service/dal/model"
)

type SepoliaSrv interface {
	GenerateNewWallet(ctx context.Context) (model.NewWallet, error)
}

type Service struct {
	sepoliaSrv SepoliaSrv
}

func New(sepoliaSrv SepoliaSrv) Service {
	return Service{
		sepoliaSrv: sepoliaSrv,
	}
}

func (s Service) GetANewWallet(ctx context.Context) (model.NewWallet, error) {
	return s.sepoliaSrv.GenerateNewWallet(ctx)
}
