package handler

import (
	"context"
	"errors"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	NFTListForAddress(ctx context.Context, input model.InputToGetMyNftListConverted) ([]interface{}, error)
	PutMyNftOnSale(ctx context.Context, input model.InputToPutNftOnSaleConverted) (interface{}, error)
	BuyNftOnSale(ctx context.Context, input model.InputToBuyNftConverted) (interface{}, error)
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
	inputData, ok := ctx.Locals("inputData").(model.InputToGetMyNftListConverted)
	if !ok {
		err := errors.New("error getting validated input from context")
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]error{"error": err})
	}

	resp, err := h.service.NFTListForAddress(ctx.Context(), inputData)
	if err != nil {
		msg := map[string]error{"error": err}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h Handler) PutNftOnSale(ctx *fiber.Ctx) error {
	inputData, ok := ctx.Locals("inputData").(model.InputToPutNftOnSaleConverted)
	if !ok {
		err := errors.New("error getting validated input from context")
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]error{"error": err})
	}

	resp, err := h.service.PutMyNftOnSale(ctx.Context(), inputData)
	if err != nil {
		msg := map[string]error{"error": err}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h Handler) BuyNft(ctx *fiber.Ctx) error {
	inputData, ok := ctx.Locals("inputData").(model.InputToBuyNftConverted)
	if !ok {
		err := errors.New("error getting validated input from context")
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]error{"error": err})
	}

	resp, err := h.service.BuyNftOnSale(ctx.Context(), inputData)
	if err != nil {
		msg := map[string]error{"error": err}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
