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
	if tr.State == "" || tr.Amount == "" || tr.TransactionID == "" {
		return errors.New("invalid data")
	}

	if tr.State != "win" && tr.State != "lose" {
		return errors.New("invalid state value")
	}

	if _, err := utils.AtoiFloat64(tr.Amount); err != nil {
		return err
	}

	return nil
}
