package main

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/topic"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/topic/dal"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/topic/dal/repository"
	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error getting service configuration for transaction-eventing: ", err.Error())
	}

	dbMarketplaceInstance, err := repository.NewConnection(serviceConfig)
	if err != nil {
		log.Fatal("error creating db client connection for transaction-eventing: ", err)
	}

	events, err := messaging.NewNatsBroker(serviceConfig)
	if err != nil {
		log.Fatal("error creating messaging client for transaction-eventing:: ", err)
	}

	topicsList := buildTopicLis(serviceConfig, dbMarketplaceInstance)
	srv := internal.NewServer(serviceConfig, events, topicsList)
	srv.StartListenService(context.Background())
}

func buildTopicLis(serviceConfig config.Config, dbMarketplaceInstance repository.PgConnection) map[string]internal.Topic {
	dalOrders, err := dal.New(serviceConfig, dbMarketplaceInstance)
	if err != nil {
		log.Fatal("error getting dalOrders for transaction-eventing: ", err)
	}

	return map[string]internal.Topic{
		"orderTopic": topic.NewOrderTopic(dalOrders),
	}
}
