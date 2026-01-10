package repositories

import (
	"expense-tracker/internals/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	RemoveExpense(id uint) error
	Update(id uint, updatedData *models.Expense) error
	Search(category string, title string, sortBy string, order string) ([]models.Expense, error)
	GetTotalAmount() (float64, error)
	GetCategorySummary() ([]CategorySummary, error)
}

type sqliteRepo struct {
	db *gorm.DB
}
type CategorySummary struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &sqliteRepo{db: db}
}
func (r *sqliteRepo) Create(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *sqliteRepo) RemoveExpense(id uint) error {
	return r.db.Delete(&models.Expense{}, id).Error
}
func (r *sqliteRepo) Update(id uint, updatedData *models.Expense) error {
	return r.db.Model(&models.Expense{}).Where("id = ?", id).Updates(updatedData).Error
}

// 1. Update the signature to accept sortBy and order strings
func (r *sqliteRepo) Search(category string, title string, sortBy string, order string) ([]models.Expense, error) {
	var expenses []models.Expense
	query := r.db.Model(&models.Expense{})

	// Filtering
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	// Sorting (e.g., "amount desc")
	if sortBy != "" && order != "" {
		query = query.Order(sortBy + " " + order)
	}

	err := query.Find(&expenses).Error
	return expenses, err
}
func (r *sqliteRepo) GetTotalAmount() (float64, error) {
	var total float64

	err := r.db.Model(&models.Expense{}).Select("sum(amount)").Scan(&total).Error
	return total, err
}
func (r *sqliteRepo) GetCategorySummary() ([]CategorySummary, error) {
	var summaries []CategorySummary

	err := r.db.Model(&models.Expense{}).
		Select("category, sum(amount) as total").
		Group("category").
		Scan(&summaries).Error

	return summaries, err
}
