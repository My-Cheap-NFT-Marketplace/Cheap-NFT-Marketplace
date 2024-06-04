package handler

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/internal/service/dal/model"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetANewWallet(ctx context.Context) (model.NewWallet, error)
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

func (h Handler) CreateWallet(ctx *fiber.Ctx) error {
	resp, err := h.service.GetANewWallet(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
