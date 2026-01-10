package handlers

import (
	"encoding/csv"
	"expense-tracker/internals/models"
	"expense-tracker/internals/services"
	"fmt"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}
	if err := h.services.AddExpense(&expense); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, expense)

}
func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	category := c.Query("category")
	title := c.Query("title")

	// Capture sorting from URL or use defaults
	sortBy := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	// Call the unified search service
	expenses, err := h.services.SearchExpenses(category, title, sortBy, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search/Sort failed"})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := h.services.RemoveExpense(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.ParseUint(paramID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.UpdateExpense(uint(id), &expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
}
func (h *ExpenseHandler) GetSummary(c *gin.Context) {
	total, err := h.services.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not calculate summary"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_spending": total,
		"currency":       "USD",
	})
}
func (h *ExpenseHandler) GetCategorySummary(c *gin.Context) {
	stats, err := h.services.GetCategoryStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
func (h *ExpenseHandler) ExportCSV(c *gin.Context) {
	expenses, err := h.services.SearchExpenses("", "", "created_at", "desc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=expenses.csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Amount", "Category", "Date"})

	for _, e := range expenses {
		row := []string{
			fmt.Sprintf("%d", e.ID),
			e.Title,
			fmt.Sprintf("%.2f", e.Amount),
			e.Category,
			e.CreatedAt.Format("2006-01-02"),
		}
		writer.Write(row)
	}
}
