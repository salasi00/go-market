package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	productdto "github.com/salasi00/go-market/internal/dto/product"
	dto "github.com/salasi00/go-market/internal/dto/result"
	"github.com/salasi00/go-market/internal/models"
	"github.com/salasi00/go-market/internal/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewHandlerProduct(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	request := new(productdto.ProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	product := models.Product{
		Name:   request.Name,
		Desc:   request.Desc,
		Price:  request.Price,
		UserID: request.UserID,
	}

	product, err = h.productService.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	product, _ = h.productService.GetProduct(product.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})

}

func (h *ProductHandler) FindProduct(c echo.Context) error {
	products, err := h.productService.FindProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: products})

}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.productService.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}
