package handler

import (
	"context"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/config"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetWalletDetail(ctx context.Context) (interface{}, error)
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

func (h Handler) WalletDetail(ctx *fiber.Ctx) error {
	resp, err := h.service.GetWalletDetail(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
