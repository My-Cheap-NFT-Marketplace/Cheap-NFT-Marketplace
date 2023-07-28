package model

type BalanceDetails struct {
	BalanceAt        BalanceDetail `json:"balanceAt,omitempty"`
	PendingBalanceAt BalanceDetail `json:"pendingBalanceAt,omitempty"`
}

type BalanceDetail struct {
	Denomination Denomination `json:"denomination,omitempty"`
	Error        string       `json:"error,omitempty"`
}

type Denomination struct {
	Exp18 string `json:"wei,omitempty"`  // wei
	Exp9  string `json:"gwei,omitempty"` // gwei
	Exp0  string `json:"eth,omitempty"`  // eth

}
