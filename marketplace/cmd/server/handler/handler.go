package handler

import (
	"context"
	"encoding/json"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	input "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/service/dal/repository/model"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	NFTListForAWallet(ctx context.Context, input input.GetNFTs) ([]model.NftToSell, error)
	UserAddNFTToSell(ctx context.Context, input input.AddNFTToSell) (model.NftToSell, error)
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
	body := ctx.Body()
	var input input.GetNFTs
	err := json.Unmarshal(body, &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	resp, err := h.service.NFTListForAWallet(ctx.Context(), input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h Handler) AddNFTToSell(ctx *fiber.Ctx) error {
	body := ctx.Body()
	var input input.AddNFTToSell
	err := json.Unmarshal(body, &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	resp, err := h.service.UserAddNFTToSell(ctx.Context(), input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
