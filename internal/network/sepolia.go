package network

import (
	"context"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/config"
)

type SepoliaConn struct {
	config config.Config
	conn   interface{}
}

func New(config config.Config) SepoliaConn {
	return SepoliaConn{
		config: config,
		conn:   nil,
	}
}

func (sc SepoliaConn) GetWalletDetail(ctx context.Context) (interface{}, error) {
	mock := map[string]string{
		"balance": "9999",
		"address": "0x597c9bc3f00a4df00f85e9334628f6cdf03a1184",
	}

	return mock, nil
}
