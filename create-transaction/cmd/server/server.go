package server

import (
	"fmt"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/cmd/config"
	"github.com/gofiber/fiber/v2"
)

type handlerIntf interface {
	SupplyMockERC20ToAccount(ctx *fiber.Ctx) error
	SupplyNftMockERC721ToAccount(ctx *fiber.Ctx) error
	CheckBalanceForAccount(ctx *fiber.Ctx) error
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
	srv.App.Post("/supply-mockerc20-to-address", srv.handler.SupplyMockERC20ToAccount)
	srv.App.Post("/supply-nft-mock-to-address", srv.handler.SupplyNftMockERC721ToAccount)
	srv.App.Post("/balance-of", srv.handler.CheckBalanceForAccount)
	return srv
}

func (srv Server) Start() error {
	err := srv.App.Listen(fmt.Sprintf(":%s", srv.config.Port))
	if err != nil {
		return err
	}

	return nil
}
