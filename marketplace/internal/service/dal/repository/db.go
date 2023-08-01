package repository

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"log"
)

type PgConnection struct {
	conn *sqlx.DB
}

func NewConnection(config config.Config) PgConnection {
	db, err := sqlx.Connect(config.Database.Marketplace.DriverName, config.Database.Marketplace.DataSourceName)
	if err != nil {
		log.Fatalln("error creating db connection: ", err)
	}

	return PgConnection{
		conn: db,
	}
}
