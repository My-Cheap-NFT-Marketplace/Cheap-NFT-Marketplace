package dal

import (
	"context"
	"crypto/ecdsa"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/cmd/server/handler/model"
	dalModel "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-transaction/internal/service/dal/model"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

const MockERC721Contract = "0xFCE9b92eC11680898c7FE57C4dDCea83AeabA3ff"

func (sc SepoliaConn) ExecSupplyMechanismsToAddMockERC721ToAccount(ctx context.Context, input model.AddTokenMockERCM721ToAddress) (dalModel.TransactionOutput, error) {
	contractAddress := common.HexToAddress(MockERC721Contract)
	contractInstance, err := mockERC721.NewMockERC721(contractAddress, sc.conn)
	if err != nil {
		return dalModel.TransactionOutput{}, err
	}

	privateKey, err := crypto.HexToECDSA(input.PrivateKey)
	if err != nil {
		return dalModel.TransactionOutput{}, err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	nonceResult, _ := sc.conn.PendingNonceAt(context.Background(), fromAddress)
	nonce := new(big.Int)
	nonce.SetUint64(nonceResult)
	gasPrice, err := sc.conn.SuggestGasPrice(context.Background())
	if err != nil {
		return dalModel.TransactionOutput{}, err
	}

	//todo delete hardcoded value for gasLimit
	//////////////
	//gasPrice := new(big.Int)
	//gasPrice.SetString("3015761820", 10)
	gasLimit := uint64(500521) // in units gas
	//////////////

	value := new(big.Int)
	value.SetUint64(0)

	chainID, _ := sc.conn.NetworkID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return dalModel.TransactionOutput{}, err
	}

	auth.Nonce = nonce
	auth.Value = value       // in wei
	auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice

	transaction, err := contractInstance.Mint(auth, fromAddress)
	if err != nil {
		return dalModel.TransactionOutput{}, err
	}

	return dalModel.TransactionOutput{
		Tx:        transaction.Hash().String(),
		GasPrice:  transaction.GasPrice().String(),
		GasTipCap: transaction.GasTipCap().String(),
		GasFeeCap: transaction.GasFeeCap().String(),
		To:        transaction.To().String(),
	}, nil
}
