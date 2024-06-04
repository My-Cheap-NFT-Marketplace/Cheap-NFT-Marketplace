package internal

import (
	"context"
	"fmt"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/model"
	"log"
)

type MessagingIntf interface {
	Close()
	ChanSubscribeToTopic(ctx context.Context, topic string) (<-chan interface{}, error)
	ListeningChanSubscribedToTopic(ch <-chan interface{}) <-chan model.Message
}

type Topic interface {
	Exec(message model.Message)
}

type Server struct {
	Config    config.Config
	Messaging MessagingIntf
	Topics    map[string]Topic
}

func NewServer(config config.Config, messaging MessagingIntf, topics map[string]Topic) Server {
	return Server{
		Config:    config,
		Messaging: messaging,
		Topics:    topics,
	}
}

func (s Server) StartListenService(ctx context.Context) {
	for _, topic := range s.Config.Topics {
		ch, err := s.Messaging.ChanSubscribeToTopic(ctx, topic)
		if err != nil {
			log.Fatalln("problem subscribing to topic: ", err)
		}

		eventMessageCh := s.Messaging.ListeningChanSubscribedToTopic(ch)
		s.DiscoverTopic(eventMessageCh)
	}
}

func (s Server) DiscoverTopic(eventMessageCh <-chan model.Message) {
	go func() {
		for eventMessage := range eventMessageCh {
			exec, ok := s.Topics[eventMessage.Topic]
			if !ok {
				log.Fatalln(fmt.Sprintf("topic %s not loaded in transaction-eventing service", eventMessage.Topic))
			}

			exec(eventMessage)
		}
	}()
}
