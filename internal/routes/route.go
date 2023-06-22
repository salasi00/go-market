package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	ProductRoutes(e)
	TransactionRoutes(e)

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
