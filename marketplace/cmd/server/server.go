package server

import (
	"fmt"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/middleware"
	"github.com/gofiber/fiber/v2"
)

type handlerIntf interface {
	NFTList(ctx *fiber.Ctx) error
	PutNftOnSale(ctx *fiber.Ctx) error
	BuyNft(ctx *fiber.Ctx) error
}

type Server struct {
	config  config.Config
	App     *fiber.App
	handler handlerIntf
}

func NewFiberServer(config config.Config, handler handlerIntf) Server {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      config.ServiceName,
	})

	return Server{
		App:     app,
		config:  config,
		handler: handler,
	}
}

func (srv Server) AddRoutes() Server {
	srv.App.Post("/my-nft-list", middleware.ValidateInputToNFTList, srv.handler.NFTList)
	srv.App.Post("/put-nft-on-sale", middleware.ValidateInputToPutNftOnSale, srv.handler.PutNftOnSale)
	srv.App.Post("/put-order-to-buy-nft", middleware.ValidateInputToBuyNft, middleware.ConvertInputToBuyNft, srv.handler.BuyNft)

	return srv
}

func (srv Server) Start() error {
	err := srv.App.Listen(fmt.Sprintf(":%s", srv.config.Port))
	if err != nil {
		return err
	}

	return nil
}
