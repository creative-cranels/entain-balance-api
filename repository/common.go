package repository

import (
	"gorm.io/gorm"
)

// Exec - simple function that returns transaction object if exists
func Exec(db *gorm.DB, tx []*gorm.DB) *gorm.DB {
	if len(tx) > 0 {
		return tx[0]
	}
	return db
}
