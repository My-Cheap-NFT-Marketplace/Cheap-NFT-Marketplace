package handler

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	NFTListForAWallet(ctx context.Context) error
}

type Handler struct {
	config  config.Config
	service Service
}

func New(config config.Config, service Service) Handler {
	return Handler{
		config:  config,
		service: service,
	}
}

func (h Handler) NFTList(ctx *fiber.Ctx) error {
	err := h.service.NFTListForAWallet(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(map[string]string{"message": "connection is ok!"})
}
