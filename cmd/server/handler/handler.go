package handler

import (
	"context"
	"encoding/json"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/config"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/server/handler/model"
	"github.com/Cheap-NFT-Marketplace/wallet-information/internal"
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
