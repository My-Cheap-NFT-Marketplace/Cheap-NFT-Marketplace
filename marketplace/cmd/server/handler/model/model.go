package model

type InputToGetMyNftList struct {
	PrivateKey  string `json:"privateKey"`
	NftContract string `json:"nftContract,omitempty"`
}

type InputToPutNftOnSale struct {
	PrivateKey  string `json:"privateKey,omitempty"`
	NftContract string `json:"nftContract,omitempty"`
	TokenId     string `json:"tokenId,omitempty"`
}
