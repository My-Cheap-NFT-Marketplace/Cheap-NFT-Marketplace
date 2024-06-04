package main

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/cmd/server"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/cmd/server/handler"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/internal"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/internal/service/dal"
	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error trying to start create-wallet service: ", err.Error())
	}

	dalSrv, err := dal.New(serviceConfig)
	if err != nil {
		log.Fatal("error getting connection with sepolia network: ", err.Error())
	}

	srv := internal.New(dalSrv)
	hndlr := handler.New(serviceConfig, srv)
	err = server.NewFiberServer(serviceConfig, hndlr).AddRoutes().Start()
	if err != nil {
		log.Fatal("error creating create-wallet server: ", err.Error())
	}
}
