package models

import (
	"time"
)

type Expense struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Title    string  `json:"title" gorm:"not null"`
	Amount   float64 `json:"amount" gorm:"not null"`
	Category string  `json:"category"`

	CreatedAt time.Time `json:"created_at"`
}
