package repository

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/cmd/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PgConnection struct {
	conn *sqlx.DB
}

func NewConnection(config config.Config) (PgConnection, error) {
	var pgConnection PgConnection
	conn, err := sqlx.Connect(config.Database.Marketplace.DriverName, config.Database.Marketplace.DataSourceName)
	if err != nil {
		return pgConnection, err
	}

	pgConnection.conn = conn
	return pgConnection, nil
}
