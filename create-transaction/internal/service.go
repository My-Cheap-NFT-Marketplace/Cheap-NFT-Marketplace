package internal

import (
	"context"
)

type SepoliaSrv interface {
}

type Service struct {
	sepoliaSrv SepoliaSrv
}

func New(sepoliaSrv SepoliaSrv) Service {
	return Service{
		sepoliaSrv: sepoliaSrv,
	}
}

func (s Service) GetTransferNftToken(ctx context.Context) interface{} {
	return nil
}
