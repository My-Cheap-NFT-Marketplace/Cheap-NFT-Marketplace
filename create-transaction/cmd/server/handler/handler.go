package handler

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/cmd/config"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetTransferNftToken(ctx context.Context) interface{}
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

func (h Handler) TransferNft(ctx *fiber.Ctx) error {
	resp := h.service.GetTransferNftToken(ctx.Context())
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
