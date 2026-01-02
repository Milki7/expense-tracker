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
	// Business Rule: You can't update an expense to have a negative amount
	if expense.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	return s.repo.Update(id, expense)
}
