package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salasi00/go-market/internal/handlers"
	"github.com/salasi00/go-market/internal/repositories"
	"github.com/salasi00/go-market/internal/services"
	"github.com/salasi00/go-market/pkg/database"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepository)
	h := handlers.NewHandlerProduct(productService)

	e.POST("/products", h.CreateProduct)
	e.GET("/products", h.FindProduct)
	e.GET("/product/:id", h.GetProduct)
}
