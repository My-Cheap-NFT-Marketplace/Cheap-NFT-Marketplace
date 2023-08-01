package model

type AddNFTToSell struct {
	TokenId         string `json:"tokenId"`
	Owner           string `json:"owner"`
	ContractAddress string `json:"contractAddress"`
	Creator         string `json:"creator"`
	TokenStandard   string `json:"tokenStandard"`
}
