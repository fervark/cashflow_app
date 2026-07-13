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
	app := &App{}
	app.service = service.New()
	app.endpoint = endpoint.New(app.service)
	app.server = echo.New()

	// Middleware
	app.server.Use(middleware.Recover(), cors.Origins())

	// Routes
	app.server.GET("/", app.endpoint.GetMain)
	//a.server.GET("/transaction-list", getTransactions)
	//a.server.GET("/movement-stats", getCashMovementStats)
	//a.server.POST("/set/transaction", setTransaction)
	//a.server.POST("/set/category", setCategory)

	return app, nil
}

func (app *App) Run() error {
	fmt.Println("Server started")

	sc := echo.StartConfig{Address: ":8080"}
	if err := sc.Start(context.Background(), app.server); err != nil {
		app.server.Logger.Error("failed to start server", "error", err)
	}

	return nil
}
