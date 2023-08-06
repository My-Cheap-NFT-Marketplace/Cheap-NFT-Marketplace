package dal

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/topic/dal/repository"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/topic/dal/repository/model"
)

type OrderDbIntrf interface {
	Exec(ctx context.Context, query string, args []interface{}) (model.ExecResult, error)
	ExecOrderQuery(ctx context.Context, query string, args []interface{}) (model.OrderItem, error)
	SelectOrderQuery(ctx context.Context, query string, args []interface{}) ([]model.OrderItem, error)
}
type DalOrder struct {
	config       config.Config
	dbConnection OrderDbIntrf
}

func New(config config.Config, dbConnection repository.PgConnection) (DalOrder, error) {
	return DalOrder{
		config:       config,
		dbConnection: dbConnection,
	}, nil
}
