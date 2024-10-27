package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID      uint64  `gorm:"primarykey"`
	Balance float64 `gorm:"type:numeric"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
