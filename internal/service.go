package internal

import "context"

type SepoliaSrv interface {
	GetWalletDetail(ctx context.Context) (interface{}, error)
}

type Service struct {
	sepoliaSrv SepoliaSrv
}

func New(sepoliaSrv SepoliaSrv) Service {
	return Service{
		sepoliaSrv: sepoliaSrv,
	}
}

func (ss Service) GetWalletDetail(ctx context.Context) (interface{}, error) {
	resp, err := ss.sepoliaSrv.GetWalletDetail(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
