package middleware

import (
	"encoding/json"
	"errors"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/gofiber/fiber/v2"
)

func ValidateInputToPutNftOnSale(ctx *fiber.Ctx) error {
	body := ctx.Body()
	var input model.InputToPutNftOnSale
	err := json.Unmarshal(body, &input)
	if err != nil {
		return errors.New("input could not be unmarshalled")
	}

	if input.NftContract == "" {
		return errors.New("NftContract can not be empty")
	}

	if input.PrivateKey == "" {
		return errors.New("PrivateKey can not be empty")
	}

	if input.TokenId == "" {
		return errors.New("TokenId can not be empty")
	}

	if input.PrivateKey[:2] == "0x" {
		input.PrivateKey = input.PrivateKey[2:]
	}

	ctx.Locals("inputData", input)
	return ctx.Next()
}

func ValidateInputToNFTList(ctx *fiber.Ctx) error {
	body := ctx.Body()
	var input model.InputToGetMyNftList
	err := json.Unmarshal(body, &input)
	if err != nil {
		return errors.New("input could not be unmarshalled")
	}

	if input.NftContract == "" {
		return errors.New("NftContract can not be empty")
	}

	if input.PrivateKey == "" {
		return errors.New("PrivateKey can not be empty")
	}

	if input.PrivateKey[:2] == "0x" {
		input.PrivateKey = input.PrivateKey[2:]
	}

	ctx.Locals("inputData", input)
	return ctx.Next()
}
