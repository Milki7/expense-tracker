package repositories

import (
	"expense-tracker/internals/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	GetAll() ([]models.Expense, error)
	RemoveExpense(id uint) error
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
func (r *sqliteRepo) RemoveExpense(id uint) error {
	return r.db.Delete(&models.Expense{}, id).Error
}
