package app

import (
	"cashflow/internal/app/endpoint"
	cors "cashflow/internal/middleware"
	"context"
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type App struct {
	endpoint *endpoint.Endpoint
	server   *echo.Echo
}

func New() (*App, error) {
	app := &App{}
	app.endpoint = endpoint.New()
	app.server = echo.New()

	// Middleware
	app.server.Use(middleware.Recover(), cors.Origins())

	// Routes
	app.server.GET("/", app.endpoint.MainPage)
	app.server.POST("/set/category", app.endpoint.SetCategory)
	app.server.GET("/transactions", app.endpoint.GetTransactions)
	app.server.POST("/set/transaction", app.endpoint.SetTransaction)
	app.server.GET("/movement-stats", app.endpoint.GetCashMovementStats)

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
