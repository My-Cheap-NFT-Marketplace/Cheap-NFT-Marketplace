package impl

type TokenObj struct {
	OwnerOf       string `json:"ownerOf,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	TokenUri      string `json:"tokenUri,omitempty"`
	Creator       string `json:"contractAddress,omitempty"`
	TokenId       string `json:"tokenId,omitempty"`
	TokenStandard string `json:"tokenStandard,omitempty"`
	NameContract  string `json:"nameContract,omitempty"`
}

type TransactionOutputObj struct {
	Tx          string `json:"tx,omitempty"`
	GasPrice    string `json:"gasPrice,omitempty"`
	GasFeeCap   string `json:"gasFeeCap,omitempty"`
	GasTipCap   string `json:"gasTipCap,omitempty"`
	To          string `json:"to,omitempty"`
	TokenAmount string `json:"TokenAmount,omitempty"`
}
