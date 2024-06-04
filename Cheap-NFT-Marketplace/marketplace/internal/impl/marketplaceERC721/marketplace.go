package marketplaceERC721

import (
	"context"
	"crypto/ecdsa"
	myCommon "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/common"
	contract "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/common/contract/trade/Marketplace/built"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/impl"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type MarketPlaceImpl struct {
	config   config.Config
	conn     *ethclient.Client
	contract *contract.Marketplace
}

func New(config config.Config) (MarketPlaceImpl, error) {
	var marketPlace MarketPlaceImpl
	conn, err := ethclient.Dial(config.Url)
	if err != nil {
		return marketPlace, err
	}

	marketContract := common.HexToAddress(config.AuctionContracts.Marketplace)
	instance, err := contract.NewMarketplace(marketContract, conn)
	if err != nil {
		return marketPlace, err
	}

	marketPlace.conn = conn
	marketPlace.config = config
	marketPlace.contract = instance
	return marketPlace, nil
}

func (dal MarketPlaceImpl) PutOrderToBuyNft(ctx context.Context, privateKey *ecdsa.PrivateKey, tokenId *big.Int, bid *big.Int) (impl.TransactionOutputObj, error) {
	var trx impl.TransactionOutputObj
	trxObj, err := myCommon.CreateTransactionObject(ctx, dal.conn, privateKey)
	if err != nil {
		return trx, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, trxObj.ChainID)
	if err != nil {
		return trx, err
	}
	auth.Nonce = trxObj.Nonce
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = trxObj.GasLimit // in units
	auth.GasPrice = trxObj.GasPrice

	marketplaceAuctionData := contract.MarketplaceAuctionData{
		CollectionAddress: common.HexToAddress(dal.config.NftContracts.MockERC721),
		Erc20Address:      common.HexToAddress(dal.config.TokenContracts.MockERC20),
		TokenId:           tokenId,
		Bid:               bid,
	}

	encodedData := encodemarketplaceAuctionData(marketplaceAuctionData)

	messageHash := crypto.Keccak256Hash(encodedData)

	bidderSig, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return trx, err
	}

	sellerKey := "9fdebc6a799893cba2ea2bf5e46a0088e0c929e341c23e48c4c223a3a96a4c79"
	ownerPrivateKey, err := crypto.HexToECDSA(sellerKey)
	if err != nil {
		return trx, err
	}

	ownerApprovedSig, err := crypto.Sign(messageHash.Bytes(), ownerPrivateKey)

	transaction, err := dal.contract.FinishAuction(auth, marketplaceAuctionData, bidderSig, ownerApprovedSig)
	if err != nil {
		return trx, err
	}

	trx.Tx = transaction.Hash().String()
	trx.GasPrice = transaction.GasPrice().String()
	trx.GasTipCap = transaction.GasTipCap().String()
	trx.GasFeeCap = transaction.GasFeeCap().String()
	trx.To = transaction.To().String()
	return trx, nil
}

func encodemarketplaceAuctionData(data contract.MarketplaceAuctionData) []byte {
	return append(append(append(data.CollectionAddress.Bytes(), data.Erc20Address.Bytes()...), data.TokenId.Bytes()...), data.Bid.Bytes()...)
}
