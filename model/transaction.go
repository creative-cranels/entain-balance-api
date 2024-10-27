package model

import (
	"errors"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID     uint64
	Amount     float64
	Type       TransactionType `gorm:"column:transaction_type"`
	ExternalID string

	User *User
}

type TransactionType string

const (
	TransactionTypeAdd TransactionType = "ADDITION"
	TransactionTypeSub TransactionType = "SUBTRACTION"
)

func AllTransactionTypes() []TransactionType {
	return []TransactionType{
		TransactionTypeAdd,
		TransactionTypeSub,
	}
}

func (e TransactionType) IsValid() error {
	switch e {
	case TransactionTypeAdd, TransactionTypeSub:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e TransactionType) String() string {
	return string(e)
}
