package main

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/impl/marketplaceERC721"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/impl/mockERC721"

	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error trying to get config for marketplace-server: ", err.Error())
	}

	nftImplementations := buildNftContractsList(serviceConfig)
	auctionImplementations := buildAuctionContractsList(serviceConfig)

	srv := internal.New(nftImplementations, auctionImplementations)
	hndlr := handler.New(serviceConfig, srv)
	err = server.NewFiberServer(serviceConfig, hndlr).AddRoutes().Start()
	if err != nil {
		log.Fatal("error creating marketplace-server: ", err.Error())
	}
}

// todo make server resilient when a contract instance throws error
func buildNftContractsList(serviceConfig config.Config) map[string]internal.NftStandard {
	mockERC721Impl, err := mockERC721.New(serviceConfig)
	if err != nil {
		log.Fatal("error getting instance of mockERC721 nft contract: ", err.Error())
	}

	return map[string]internal.NftStandard{
		serviceConfig.NftContracts.MockERC721: mockERC721Impl,
	}
}

// todo make server resilient when a contract instance throws error
func buildAuctionContractsList(serviceConfig config.Config) map[string]internal.AuctionStandard {
	marketplaceERC721Impl, err := marketplaceERC721.New(serviceConfig)
	if err != nil {
		log.Fatal("error getting instance of marketplaceERC721 auction contract: ", err.Error())
	}

	return map[string]internal.AuctionStandard{
		serviceConfig.AuctionContracts.Marketplace: marketplaceERC721Impl,
	}
}
