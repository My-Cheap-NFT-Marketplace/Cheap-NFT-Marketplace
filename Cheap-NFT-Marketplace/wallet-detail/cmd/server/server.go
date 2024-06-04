package server

import (
	"fmt"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/cmd/config"
	"github.com/gofiber/fiber/v2"
)

type handlerIntf interface {
	WalletDetail(ctx *fiber.Ctx) error
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
	srv.App.Post("/wallet-detail", srv.handler.WalletDetail)
	return srv
}

func (srv Server) Start() error {
	err := srv.App.Listen(fmt.Sprintf(":%s", srv.config.Port))
	if err != nil {
		return err
	}

	return nil
}
