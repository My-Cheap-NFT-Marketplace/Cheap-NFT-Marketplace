package model

type QueryNft struct {
	Limit  *string `json:"limit"`
	Offset *string `json:"offset"`
	OrderItem
}

type OrderItem struct {
	Trx             *string `json:"trx" db:"trx"`
	TokenId         *string `json:"tokenId" db:"token_id"`
	Owner           *string `json:"owner" db:"owner"`
	ContractAddress *string `json:"contractAddress" db:"contract_address"`
	Creator         *string `json:"creator" db:"creator"`
	TokenStandard   *string `json:"tokenStandard" db:"token_standard"`
	Status          *string `json:"status" db:"status"`
	CreatedAt       *string `json:"createdAt" db:"created_at"`
	UpdatedAt       *string `json:"UpdatedAt" db:"updated_at"`
}

type ExecResult struct {
	LastInsertId int64 `json:"lastInsertId"`
	RowsAffected int64 `json:"rowsAffected"`
}
