package model

type AddNFTToSell struct {
	TokenId         *string `json:"tokenId"`
	Owner           *string `json:"owner"`
	ContractAddress *string `json:"contractAddress"`
	Creator         *string `json:"creator"`
	TokenStandard   *string `json:"tokenStandard"`
}

type GetNFTs struct {
	TokenId *string `json:"tokenId"`
	Owner   *string `json:"owner"`
	Limit   *string `default:"20" json:"limit"`
	Offset  *string `default:"0" json:"offset"`
}
