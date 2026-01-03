package repositories

import (
	"expense-tracker/internals/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	GetAll() ([]models.Expense, error)
	RemoveExpense(id uint) error
	Update(id uint, updatedData *models.Expense) error
	Search(category string, title string) ([]models.Expense, error)
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
func (r *sqliteRepo) Update(id uint, updatedData *models.Expense) error {
	return r.db.Model(&models.Expense{}).Where("id = ?", id).Updates(updatedData).Error
}
func (r *sqliteRepo) Search(category string, title string) ([]models.Expense, error) {
	var expenses []models.Expense
	query := r.db.Model(&models.Expense{})

	// If the user provided a category, filter by it
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// If the user provided a title, search for titles containing that text
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	err := query.Find(&expenses).Error
	return expenses, err
}
