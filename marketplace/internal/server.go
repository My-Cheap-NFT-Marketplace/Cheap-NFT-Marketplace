package internal

import (
	"context"
)

type SepoliaSrv interface {
	GetNFTListForAWallet(ctx context.Context) error
}

type Service struct {
	sepoliaSrv SepoliaSrv
}

func New(sepoliaSrv SepoliaSrv) Service {
	return Service{
		sepoliaSrv: sepoliaSrv,
	}
}

func (s Service) NFTListForAWallet(ctx context.Context) error {
	return s.sepoliaSrv.GetNFTListForAWallet(ctx)
}
