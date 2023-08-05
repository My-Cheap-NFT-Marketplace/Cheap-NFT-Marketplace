package middleware

import (
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/server/handler/model"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"math/big"
)

func ConvertInputToBuyNft(ctx *fiber.Ctx) error {
	inputData, _ := ctx.Locals("rawInputData").(model.RawInputToBuyNft)
	var convertedInput model.InputToBuyNftConverted
	convertedInput.AuctionContract = *inputData.AuctionContract

	pvKey, err := crypto.HexToECDSA(*inputData.PrivateKey)
	if err != nil {
		return err
	}
	convertedInput.PrivateKey = pvKey

	tokenId := new(big.Int)
	tokenId.SetString(*inputData.TokenId, 10)
	convertedInput.TokenId = tokenId

	bid := new(big.Int)
	bid.SetString(*inputData.Bid, 10)
	convertedInput.TokenId = bid
	ctx.Locals("inputData", convertedInput)
	return ctx.Next()
}
