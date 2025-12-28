package repositories

import (
	"expense-tracker/internals/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Created(expense *models.Expense) error
	GetAll() ([]models.Expense, error)
}
type sqliteRepo struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &sqliteRepo{db: db}
}
func (r *sqliteRepo) Create(expense *models.Expense) error {
	return r.db.Create(expense).Error
}
func (r *sqliteRepo) GetAll() ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.db.Find(&expenses).Error
	return expenses, err
}
