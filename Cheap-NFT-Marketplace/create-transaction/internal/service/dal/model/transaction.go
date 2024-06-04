package model

type TransactionOutput struct {
	Tx          string `json:"tx"`
	GasPrice    string `json:"gasPrice"`
	GasFeeCap   string `json:"gasFeeCap"`
	GasTipCap   string `json:"gasTipCap"`
	To          string `json:"to"`
	TokenAmount string `json:"TokenAmount"`
}
