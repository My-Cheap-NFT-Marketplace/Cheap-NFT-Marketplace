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

func ValidateInputToBuyNft(ctx *fiber.Ctx) error {
	body := ctx.Body()
	var input model.RawInputToBuyNft
	err := json.Unmarshal(body, &input)
	if err != nil {
		return errors.New("input could not be unmarshalled")
	}
	if input.AuctionContract == nil {
		return errors.New("auction contract can not be empty")
	}

	if input.PrivateKey == nil {
		return errors.New("privateKey can not be empty")
	}

	pvkey := *input.PrivateKey
	if pvkey[0:2] == "0x" {
		*input.PrivateKey = pvkey[2:]
	}

	if input.TokenId == nil {
		return errors.New("tokenId can not be empty")
	}

	if input.Bid == nil {
		return errors.New("bid can not be empty")
	}

	ctx.Locals("rawInputData", input)
	return ctx.Next()
}
