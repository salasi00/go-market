package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	dto "github.com/salasi00/go-market/internal/dto/result"
	transactiondto "github.com/salasi00/go-market/internal/dto/transaction"
	"github.com/salasi00/go-market/internal/metrics"
	"github.com/salasi00/go-market/internal/models"
	"github.com/salasi00/go-market/internal/services"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewHandlerTransaction(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) FindTransaction(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("user_id"))
	transactions, err := h.transactionService.FindTransaction(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.TransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transaction := models.Transaction{
		Total:     request.Total,
		UserID:    request.UserID,
		ProductID: request.ProductID,
	}

	data, err := h.transactionService.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	metrics.TpsCounter.Inc()
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
