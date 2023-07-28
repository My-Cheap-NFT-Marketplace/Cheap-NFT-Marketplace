package main

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/cmd/server"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/cmd/server/handler"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/internal"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/internal/dal"
	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error trying to start wallet-detail service: ", err.Error())
	}

	dalSrv, err := dal.New(serviceConfig)
	if err != nil {
		log.Fatal("error getting connection with sepolia network: ", err.Error())
	}

	srv := internal.New(dalSrv)
	hndlr := handler.New(serviceConfig, srv)
	err = server.NewFiberServer(serviceConfig, hndlr).AddRoutes().Start()
	if err != nil {
		log.Fatal("error creating wallet-detail server: ", err.Error())
	}
}
