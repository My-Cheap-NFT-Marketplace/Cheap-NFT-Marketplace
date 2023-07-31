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

const MockERC20Contract = "0xbd65c58D6F46d5c682Bf2f36306D461e3561C747"

func (sc SepoliaConn) ExecSupplyMechanismsToAddMockERC20ToAccount(ctx context.Context, input model.AddTokenMockERC20ToAddress) (dalModel.TransactionOutput, error) {
	contractAddress := common.HexToAddress(MockERC20Contract)
	contractInstance, err := MockERC20.NewMockERC20(contractAddress, sc.conn)
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

	TokenAmount := new(big.Int)
	TokenAmount.SetString(input.Amount, 10)

	//todo delete hardcoded value for gasLimit
	//////////////
	//gasPrice := new(big.Int)
	//gasPrice.SetString("3015761820", 10)
	gasLimit := uint64(500521) // in units gas
	//////////////

	value := new(big.Int)
	value.SetUint64(0)

	//auth := bind.NewKeyedTransactor(privateKey)
	chainID, _ := sc.conn.NetworkID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return dalModel.TransactionOutput{}, err
	}

	auth.Nonce = nonce
	auth.Value = value       // in wei
	auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice

	transaction, err := contractInstance.Mint(auth, fromAddress, TokenAmount)
	if err != nil {
		return dalModel.TransactionOutput{}, err
	}

	return dalModel.TransactionOutput{
		Tx:          transaction.Hash().String(),
		GasPrice:    transaction.GasPrice().String(),
		GasTipCap:   transaction.GasTipCap().String(),
		GasFeeCap:   transaction.GasFeeCap().String(),
		To:          transaction.To().String(),
		TokenAmount: input.Amount,
	}, nil
}

func (sc SepoliaConn) ExecGetBalanceForAccount(ctx context.Context, input model.GetBalanceForAddress) (interface{}, error) {
	contractAddress := common.HexToAddress(MockERC20Contract)
	address := common.HexToAddress(input.Address)
	//todo set context
	callOpts := &bind.CallOpts{}

	resp, err := MockERC20.NewMockERC20(contractAddress, sc.conn)
	if err != nil {
		return nil, err
	}

	balance, err := resp.BalanceOf(callOpts, address)
	if err != nil {
		return nil, err
	}

	//todo set differente exp for balance
	return map[string]string{"balance": balance.String()}, nil
}
