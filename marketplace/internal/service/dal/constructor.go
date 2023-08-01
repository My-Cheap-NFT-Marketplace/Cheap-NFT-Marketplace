package dal

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/service/dal/repository"
)

type DataAccessLayer struct {
	config       config.Config
	dbConnection repository.PgConnection
}

func New(config config.Config, dbConnection repository.PgConnection) (DataAccessLayer, error) {
	return DataAccessLayer{
		config:       config,
		dbConnection: dbConnection,
	}, nil
}
