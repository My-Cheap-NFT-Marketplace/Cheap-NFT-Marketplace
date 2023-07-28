package main

import (
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/config"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/server"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/server/handler"
	"github.com/Cheap-NFT-Marketplace/wallet-information/internal"
	"github.com/Cheap-NFT-Marketplace/wallet-information/internal/dal"
	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error trying to start wallet-information service: ", err.Error())
	}

	dalSrv, err := dal.New(serviceConfig)
	if err != nil {
		log.Fatal("error getting connection with sepolia network: ", err.Error())
	}

	srv := internal.New(dalSrv)
	hndlr := handler.New(serviceConfig, srv)
	err = server.NewFiberServer(serviceConfig, hndlr).AddRoutes().Start()
	if err != nil {
		log.Fatal("error creating wallet-information service: ", err.Error())
	}
}
