package handler

import (
	"context"
	"encoding/json"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/cmd/server/handler/model"
	dalModel "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/internal/service/dal/model"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetSupplyMockERC20ToAccount(ctx context.Context, input model.AddTokenMockERC20ToAddress) (dalModel.TransactionOutput, error)
	GetSupplyNftMockERC721ToAccount(ctx context.Context, input model.AddTokenMockERCM721ToAddress) (dalModel.TransactionOutput, error)
	GetBalanceForAccount(ctx context.Context, input model.GetBalanceForAddress) (interface{}, error)
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

func (h Handler) SupplyMockERC20ToAccount(ctx *fiber.Ctx) error {
	var input model.AddTokenMockERC20ToAddress
	body := ctx.Body()
	err := json.Unmarshal(body, &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	resp, err := h.service.GetSupplyMockERC20ToAccount(ctx.Context(), input)
	if err != nil {
		msg := map[string]string{"error": err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h Handler) SupplyNftMockERC721ToAccount(ctx *fiber.Ctx) error {
	var input model.AddTokenMockERCM721ToAddress
	body := ctx.Body()
	err := json.Unmarshal(body, &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	resp, err := h.service.GetSupplyNftMockERC721ToAccount(ctx.Context(), input)
	if err != nil {
		msg := map[string]string{"error": err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h Handler) CheckBalanceForAccount(ctx *fiber.Ctx) error {
	var input model.GetBalanceForAddress
	body := ctx.Body()
	err := json.Unmarshal(body, &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	resp, err := h.service.GetBalanceForAccount(ctx.Context(), input)
	if err != nil {
		msg := map[string]string{"error": err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
