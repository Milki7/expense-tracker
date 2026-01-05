package models

import (
	"time"
)

type Expense struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Title string `json:"title" gorm:"not null" binding:"required"`

	Amount float64 `json:"amount" gorm:"not null" binding:"required,gt=0"`

	Category string `json:"category" binding:"required"`

	CreatedAt time.Time `json:"created_at"`
}
