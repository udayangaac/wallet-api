package operation

// RetrieveWalletBalanceResp a model for common retrieve wallet balance response.
type RetrieveWalletBalanceResp struct {
	Balance float32 `json:"balance"`
}

// NewRetrieveWalletBalanceResp create an instance of retrieve wallet balance response.
func NewRetrieveWalletBalanceResp(balance float32) RetrieveWalletBalanceResp {
	return RetrieveWalletBalanceResp{
		Balance: balance,
	}
}
