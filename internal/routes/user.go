package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salasi00/go-market/internal/handlers"
	"github.com/salasi00/go-market/internal/repositories"
	"github.com/salasi00/go-market/internal/services"
	"github.com/salasi00/go-market/pkg/database"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	h := handlers.NewHandlerUser(userService)

	e.POST("/users", h.CreateUser)
	e.GET("/user/:id", h.GetUser)
}
