package handlers

import (
	"expense-tracker/internals/models"
	"expense-tracker/internals/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	services services.ExpenseService
}

func NewExpensehandler(services services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{services: services}

}
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}
	if err := h.services.AddExpense(&expense); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, expense)

}
func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	expenses, err := h.services.FetchAllExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch expenses"})
		return
	}
	c.JSON(http.StatusOK, expenses)
}
