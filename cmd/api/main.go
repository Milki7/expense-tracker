package main

import (
	"expense-tracker/config"
	"expense-tracker/internals/handlers"
	"expense-tracker/internals/repositories"
	"expense-tracker/internals/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	expenseRepo := repositories.NewExpenseRepository(config.DB)
	expenseService := services.NewExpenseService(expenseRepo)
	expenseHandler := handlers.NewExpensehandler(expenseService)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/expenses", expenseHandler.CreateExpense)
		api.GET("/expenses", expenseHandler.GetExpenses)
		api.DELETE("/expenses/:id", expenseHandler.DeleteExpense)
		api.PUT("/expenses/:id", expenseHandler.UpdateExpense)
	}

	r.Run(":8080")
}
