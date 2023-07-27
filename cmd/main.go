package main

import (
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/config"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/server"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/server/handler"
	"github.com/Cheap-NFT-Marketplace/wallet-information/internal"
	"github.com/Cheap-NFT-Marketplace/wallet-information/internal/network"
	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error trying to start wallet-information service: ", err.Error())
	}

	ethNetwork := network.New(serviceConfig)
	srv := internal.New(ethNetwork)
	hndlr := handler.New(serviceConfig, srv)
	err = server.NewFiberServer(serviceConfig, hndlr).AddRoutes().Start()
	if err != nil {
		log.Fatal("error creating wallet-information service: ", err.Error())
	}
}
