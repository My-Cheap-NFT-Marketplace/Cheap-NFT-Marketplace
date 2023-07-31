package dal

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SepoliaConn struct {
	config config.Config
	conn   *ethclient.Client
}

func New(config config.Config) (SepoliaConn, error) {
	conn, err := ethclient.Dial(config.Url)
	if err != nil {
		return SepoliaConn{}, err
	}

	return SepoliaConn{
		config: config,
		conn:   conn,
	}, nil
}

func (sc SepoliaConn) GetNFTListForAWallet(ctx context.Context) error {
	return nil
}
