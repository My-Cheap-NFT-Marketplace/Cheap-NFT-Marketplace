package model

type NftToSell struct {
	TokenId         string `json:"tokenId" db:"tokenId"`
	Owner           string `json:"owner" db:"owner"`
	ContractAddress string `json:"contractAddress" db:"contractAddress"`
	Creator         string `json:"creator" db:"creator"`
	TokenStandard   string `json:"tokenStandard" db:"tokenStandard"`
	Status          string `json:"status" db:"status"`
	CreatedAt       string `json:"createdAt" db:"createdAt"`
	UpdatedAt       string `json:"UpdatedAt" db:"UpdatedAt"`
}

type ExecResult struct {
	LastInsertId int64 `json:"lastInsertId"`
	RowsAffected int64 `json:"rowsAffected"`
}
