package main

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/impl/mockERC721"
	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error trying to start marketplace-server: ", err.Error())
	}

	mockERC721Impl, err := mockERC721.New(serviceConfig)
	if err != nil {
		log.Fatal("error getting connection with sepolia network: ", err.Error())
	}

	nftImplementations := map[string]internal.ERC721Standard{
		serviceConfig.NftContracts.MockERC721: mockERC721Impl,
	}

	srv := internal.New(nftImplementations)
	hndlr := handler.New(serviceConfig, srv)
	err = server.NewFiberServer(serviceConfig, hndlr).AddRoutes().Start()
	if err != nil {
		log.Fatal("error creating marketplace-server: ", err.Error())
	}
}
