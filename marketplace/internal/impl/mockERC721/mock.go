package mockERC721

import (
	"context"
	myCommon "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/common"
	mockContract "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/common/contract/nfterc721/mock/built"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/cmd/config"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/impl"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const tokenStandardERC721 = "ERC-721"

type MockERC721Impl struct {
	config   config.Config
	conn     *ethclient.Client
	contract *mockContract.MockERC721
}

func New(config config.Config) (MockERC721Impl, error) {
	var mockERC721Impl MockERC721Impl
	conn, err := ethclient.Dial(config.Url)
	if err != nil {
		return mockERC721Impl, err
	}

	instance, err := mockContract.NewMockERC721(common.HexToAddress(config.NftContracts.MockERC721), conn)
	return MockERC721Impl{
		config:   config,
		conn:     conn,
		contract: instance,
	}, nil
}

func (dal MockERC721Impl) BalanceOf(ctx context.Context, privateKey string) (*big.Int, error) {
	userAddress, err := myCommon.GetUserAddress(privateKey)
	if err != nil {
		return nil, err
	}

	balance, err := dal.contract.BalanceOf(&bind.CallOpts{}, userAddress)
	if err != nil {
		return nil, err
	}

	return balance, nil

}

func (dal MockERC721Impl) TokenOfOwnerByIndex(ctx context.Context, privateKey string, index *big.Int) (*big.Int, error) {
	userAddress, err := myCommon.GetUserAddress(privateKey)
	if err != nil {
		return nil, err
	}

	tokenId, err := dal.contract.TokenOfOwnerByIndex(&bind.CallOpts{}, userAddress, index)
	if err != nil {
		return nil, err
	}

	return tokenId, nil
}

func (dal MockERC721Impl) PutNftOnSale(ctx context.Context, privateKey string, tokenId *big.Int) (impl.TransactionOutputObj, error) {
	var trx impl.TransactionOutputObj
	trxObj, err := myCommon.CreateTransactionObject(ctx, dal.conn, privateKey, &dal.config.AuctionContracts.Marketplace, nil, nil)
	if err != nil {
		return trx, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(trxObj.PrivateKey, trxObj.ChainID)
	if err != nil {
		return trx, err
	}

	transaction, err := dal.contract.Approve(auth, trxObj.To, tokenId)
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

func (dal MockERC721Impl) BuildNtfObject(ctx context.Context, tokenID *big.Int) (impl.TokenObj, error) {
	var nftObj impl.TokenObj
	callOpts := &bind.CallOpts{}
	name, err := dal.contract.Name(callOpts)
	if err != nil {
		return nftObj, err
	}

	symbol, err := dal.contract.Symbol(callOpts)
	if err != nil {
		return nftObj, err
	}

	address, err := dal.contract.OwnerOf(callOpts, tokenID)
	if err != nil {
		return nftObj, err
	}

	tokenUri, err := dal.contract.TokenURI(callOpts, tokenID)
	if err != nil {
		return nftObj, err
	}

	nftObj.NameContract = name
	nftObj.Symbol = symbol
	nftObj.OwnerOf = address.String()
	nftObj.TokenId = tokenID.String()
	nftObj.TokenStandard = tokenStandardERC721
	nftObj.Creator = dal.config.NftContracts.MockERC721
	nftObj.TokenUri = tokenUri
	return nftObj, nil
}
