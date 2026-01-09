package services

import (
	"errors"
	"expense-tracker/internals/models"
	"expense-tracker/internals/repositories"
)

type ExpenseService interface {
	AddExpense(expense *models.Expense) error
	FetchAllExpenses() ([]models.Expense, error)
	RemoveExpense(id uint) error
	UpdateExpense(id uint, expense *models.Expense) error
	SearchExpenses(category, title, sortBy, order string) ([]models.Expense, error)
	GetSummary() (float64, error)
	GetCategoryStats() ([]repositories.CategorySummary, error)
}
type expenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}
func (s *expenseService) AddExpense(expense *models.Expense) error {
	if expense.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if expense.Title == "" {
		return errors.New("Expense title is required")
	}
	if len(expense.Category) < 3 {
		return errors.New("category name is too short")
	}
	return s.repo.Create(expense)
}

func (s *expenseService) FetchAllExpenses() ([]models.Expense, error) {
	return s.repo.GetAll()
}
func (s *expenseService) RemoveExpense(id uint) error {
	if id == 0 {
		return errors.New("Invalid expense ID")
	}
	return s.repo.RemoveExpense(id)

}
func (s *expenseService) UpdateExpense(id uint, expense *models.Expense) error {
	if id == 0 {
		return errors.New("invalid ID")
	}

	if expense.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	return s.repo.Update(id, expense)
}

func (s *expenseService) SearchExpenses(category string, title string) ([]models.Expense, error) {
	return s.repo.Search(category, title)
	return s.repo.Search(category, title, sortBy, order)
}
func (s *expenseService) GetSummary() (float64, error) {
	return s.repo.GetTotalAmount()
}
func (s *expenseService) GetCategoryStats() ([]repositories.CategorySummary, error) {
	return s.repo.GetCategorySummary()
}
func (s *expenseService) FetchSortedExpenses(sortBy string, order string) ([]models.Expense, error) {

	allowedColumns := map[string]bool{"amount": true, "title": true, "created_at": true, "category": true}

	if !allowedColumns[sortBy] {
		sortBy = "created_at"
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	return s.repo.GetAllSorted(sortBy, order)
}
