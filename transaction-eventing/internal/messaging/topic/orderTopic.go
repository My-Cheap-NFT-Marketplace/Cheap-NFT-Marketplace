package topic

import (
	"context"
	"errors"
	"fmt"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/model"
	dbModel "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/topic/dal/repository/model"
	"log"
)

type DalOrderIntrf interface {
	CreateOrderToSell(ctx context.Context, input dbModel.OrderItem) error
	UpdateStatusOrder(ctx context.Context, input dbModel.OrderItem) error
}

type OrderTopic struct {
	dalOrder DalOrderIntrf
}

func NewOrderTopic(db DalOrderIntrf) OrderTopic {
	return OrderTopic{
		dalOrder: db,
	}
}

func (ot OrderTopic) Exec(message model.Message) {
	switch message.EventType {
	case "create":
		ot.createOrderToPutNFTOnSale(message)
	case "update":
		ot.updateStatusOder(message)
	default:
		log.Fatalln(fmt.Sprintf("event type %s for order topic does not exists", message.EventType))
	}
}

func (ot OrderTopic) createOrderToPutNFTOnSale(input model.Message) {
	var item dbModel.OrderItem
	var ok bool
	if item, ok = input.Message.(dbModel.OrderItem); !ok {
		log.Println(errors.New("message to add in order table is not OrderItem object"))
	}
	err := ot.dalOrder.CreateOrderToSell(context.Background(), item)
	if err != nil {
		log.Println(err)
	}
}

func (ot OrderTopic) updateStatusOder(input model.Message) {
	var item dbModel.OrderItem
	var ok bool
	if item, ok = input.Message.(dbModel.OrderItem); !ok {
		log.Println(errors.New("message to update in order table is not OrderItem object"))
	}
	err := ot.dalOrder.UpdateStatusOrder(context.Background(), item)
	if err != nil {
		log.Println(err)
	}
}
