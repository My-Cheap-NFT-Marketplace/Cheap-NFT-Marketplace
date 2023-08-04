package platform

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

//this methods are going go be moved to common folder

func GetUserAddress(privateKey string) (common.Address, error) {
	var userAddress common.Address
	pvKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return userAddress, err
	}

	publicKey := pvKey.Public().(*ecdsa.PublicKey)
	userAddress = crypto.PubkeyToAddress(*publicKey)
	return userAddress, nil
}

type NewTransactionObj struct {
	PrivateKey  *ecdsa.PrivateKey
	fromAddress common.Address
	Nonce       *big.Int
	GasPrice    *big.Int
	GasLimit    uint64
	ChainID     *big.Int
	TokenAmount *big.Int
	Value       *big.Int
	To          common.Address
}

func CreateTransactionObject(ctx context.Context, conn *ethclient.Client, privateKey string, to *string, tokenAmount *string, value *string) (NewTransactionObj, error) {
	var transactionObj NewTransactionObj

	pvKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return transactionObj, err
	}

	userAddress, err := GetUserAddress(privateKey)
	if err != nil {
		return transactionObj, err
	}

	nonceResult, _ := conn.PendingNonceAt(ctx, userAddress)
	nonce := new(big.Int)
	nonce.SetUint64(nonceResult)

	gasPrice, err := conn.SuggestGasPrice(ctx)
	if err != nil {
		return transactionObj, err
	}

	// todo check how to obtain an appropiate gasLimit
	gasLimit := uint64(500521)

	chainID, _ := conn.NetworkID(ctx)
	if err != nil {
		return transactionObj, err
	}

	transactionObj.PrivateKey = pvKey
	transactionObj.fromAddress = userAddress
	transactionObj.Nonce = nonce
	transactionObj.GasPrice = gasPrice
	transactionObj.GasLimit = gasLimit
	transactionObj.ChainID = chainID

	if to != nil {
		toAddress := common.HexToAddress(*to)
		transactionObj.To = toAddress
	}
	if tokenAmount != nil {
		amount := new(big.Int)
		amount.SetString(*tokenAmount, 10)
		transactionObj.TokenAmount = amount
	}

	if value != nil {
		amount := new(big.Int)
		amount.SetString(*value, 10)
		transactionObj.Value = amount
	}

	return transactionObj, nil
}
