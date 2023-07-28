package model

type NonceDetails struct {
	PendingNonceAt NonceDetail `json:"pendingNonceAt,omitempty"`
	NonceAt        NonceDetail `json:"nonceAt,omitempty"`
}

type NonceDetail struct {
	Count uint64 `json:"count,omitempty"`
	Error string `json:"error,omitempty"`
}
