package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salasi00/go-market/internal/handlers"
	"github.com/salasi00/go-market/internal/repositories"
	"github.com/salasi00/go-market/internal/services"
	"github.com/salasi00/go-market/pkg/database"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.NewTransactionRepository(database.DB)
	transactionService := services.NewTransactionService(transactionRepository)
	h := handlers.NewHandlerTransaction(transactionService)

	e.GET("/transactions/:user_id", h.FindTransaction)
	e.POST("/transaction", h.CreateTransaction)
}
