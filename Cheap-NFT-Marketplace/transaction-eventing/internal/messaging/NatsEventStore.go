package messaging

import (
	"context"
	"fmt"
	myCommon "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/common"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/model"
	"github.com/nats-io/nats.go"
	"log"
)

type NatsEventStore struct {
	conn            *nats.Conn
	natSubscription *nats.Subscription
}

func NewNatsBroker(config config.Config) (NatsEventStore, error) {
	var natsEventStore NatsEventStore
	conn, err := nats.Connect(config.Nats.Host)
	if err != nil {
		return natsEventStore, err
	}

	natsEventStore.conn = conn
	return natsEventStore, nil
}

func (es NatsEventStore) Close() {
	if es.conn != nil {
		es.conn.Close()
	}

	if es.natSubscription != nil {
		es.natSubscription.Unsubscribe()
	}
}

func (es NatsEventStore) ChanSubscribeToTopic(ctx context.Context, topic string) (<-chan interface{}, error) {
	inputEventChan := make(chan interface{})
	var err error
	ch := make(chan *nats.Msg)
	// TODO figure out how to use subscription variable from ChanSubscribe
	fmt.Println("listening topic: ", topic)
	_, err = es.conn.ChanSubscribe(topic, ch)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(inputEventChan)
		for {
			select {
			case msg := <-ch:
				inputEventChan <- msg
			}
		}
	}()

	return inputEventChan, nil
}

func (es NatsEventStore) ListeningChanSubscribedToTopic(ch <-chan interface{}) <-chan model.Message {
	out := make(chan model.Message)
	go func() {
		defer close(out)
		for message := range ch {
			var inputMsg *nats.Msg
			var ok bool

			if inputMsg, ok = message.(*nats.Msg); !ok {
				log.Fatalln("problem casting message from nats")
			}

			event := model.Message{}
			err := myCommon.DecodeMessage(inputMsg.Data, &event)
			if err != nil {
				log.Fatalln("problem decoding message from topic: ", err)
			}

			out <- event
		}
	}()

	return out
}
