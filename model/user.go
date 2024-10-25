package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID      uint64 `gorm:"primarykey"`
	Balance int64  `gorm:"type:integer"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
