package common

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetUserAddress(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*publicKey)
}

type NewTransactionObj struct {
	fromAddress common.Address
	Nonce       *big.Int
	GasPrice    *big.Int
	GasLimit    uint64
	ChainID     *big.Int
}

func CreateTransactionObject(ctx context.Context, conn *ethclient.Client, privateKey *ecdsa.PrivateKey) (NewTransactionObj, error) {
	var transactionObj NewTransactionObj

	userAddress := GetUserAddress(privateKey)
	nonceResult, _ := conn.PendingNonceAt(ctx, userAddress)
	nonce := big.NewInt(int64(nonceResult))
	gasPrice, err := conn.SuggestGasPrice(ctx)
	if err != nil {
		return transactionObj, err
	}
	// todo check how to obtain an appropiate gasLimit
	gasLimit := uint64(5000521)
	chainID, err := conn.NetworkID(ctx)
	if err != nil {
		return transactionObj, err
	}

	transactionObj.fromAddress = userAddress
	transactionObj.Nonce = nonce
	transactionObj.GasPrice = gasPrice
	transactionObj.GasLimit = gasLimit
	transactionObj.ChainID = chainID
	return transactionObj, nil
}
