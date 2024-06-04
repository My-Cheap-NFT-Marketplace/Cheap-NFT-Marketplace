package handler

import (
	"context"
	"encoding/json"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/cmd/server/handler/model"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/wallet-detail/internal"
	"github.com/gofiber/fiber/v2"
	"math/big"
)

type Service interface {
	GetWalletDetail(ctx context.Context, address string, blockNumber *big.Int) internal.WalletDetail
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
	body := ctx.Body()
	var input model.InputWalletDetail
	_ = json.Unmarshal(body, &input)

	resp := h.service.GetWalletDetail(ctx.Context(), input.Wallet.Address, input.Wallet.BlockNumber)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
