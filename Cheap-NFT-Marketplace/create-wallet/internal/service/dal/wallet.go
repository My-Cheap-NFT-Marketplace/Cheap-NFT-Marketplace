package dal

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/create-wallet/internal/service/dal/model"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (sc SepoliaConn) GenerateNewWallet(ctx context.Context) (model.NewWallet, error) {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		return model.NewWallet{}, err
	}

	pvData := crypto.FromECDSA(pvk)
	pvDataStr := hexutil.Encode(pvData)

	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	puDataStr := hexutil.Encode(puData)

	addressStr := crypto.PubkeyToAddress(pvk.PublicKey).Hex()

	return model.NewWallet{
		PrivateKey: pvDataStr,
		PublicKey:  puDataStr,
		Address:    addressStr,
	}, nil
}