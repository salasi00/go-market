package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/salasi00/go-market/internal/metrics"
	"github.com/salasi00/go-market/internal/routes"
	"github.com/salasi00/go-market/pkg/database"
)

func main() {
	e := echo.New()

	database.NewDBConnection()
	database.RunMigration()

	metrics.RegisterMetrics()

	port := "8080"

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("Server running on port " + port)
	e.Logger.Fatal(e.Start(":" + port))
}
