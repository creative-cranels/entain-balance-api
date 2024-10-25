package request

import (
	"balance-api/utils"
	"errors"
)

type TransactionRequest struct {
	State         string `json:"state"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionId"`
}

func (tr *TransactionRequest) Validate() error {
	if tr.State != "win" && tr.State != "lose" {
		return errors.New("invalid state value")
	}

	_, err := utils.AtoiFloat64(tr.Amount)
	if err != nil {
		return err
	}

	if tr.TransactionID == "" {
		return errors.New("invalid transactionId value")
	}
	return nil
}
