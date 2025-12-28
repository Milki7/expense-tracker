package services

import (
	"errors"
	"expense-tracker/internals/models"
	"expense-tracker/internals/repositories"
)

type ExpenseService interface {
	AddExpense(expense *models.Expense) error
	FetchAllExpense() ([]models.Expense, error)
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
	return s.repo.Created(expense)
}

func (s *expenseService) FetchAllExpense() ([]models.Expense, error) {
	return s.repo.GetAll()
}
