package request

type TransactionRequest struct {
	State         string `json:"role"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionId"`
}

func (rr *TransactionRequest) Validate() error {
	return nil
}
