package app

import (
	"cashflow/internal/app/endpoint"
	cors "cashflow/internal/app/middleware"
	"cashflow/internal/app/service"
	"context"
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	server   *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.service = service.New()
	a.endpoint = endpoint.New(a.service)
	a.server = echo.New()

	// Middleware
	a.server.Use(middleware.Recover(), cors.Origins())

	// Routes
	a.server.GET("/", a.endpoint.GetMain)
	//a.server.GET("/transaction-list", getTransactions)
	//a.server.GET("/movement-stats", getCashMovementStats)
	//a.server.POST("/set/transaction", setTransaction)
	//a.server.POST("/set/category", setCategory)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server started")

	sc := echo.StartConfig{Address: ":8080"}
	if err := sc.Start(context.Background(), a.server); err != nil {
		a.server.Logger.Error("failed to start server", "error", err)
	}

	return nil
}
