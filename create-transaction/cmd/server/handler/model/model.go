package model

type AddTokenMockERC20ToAddress struct {
	PrivateKey string `json:"privateKey"`
	Amount     string `json:"amount"`
}

type AddTokenMockERCM721ToAddress struct {
	PrivateKey string `json:"privateKey"`
}

type GetBalanceForAddress struct {
	Address string `json:"address"`
}
