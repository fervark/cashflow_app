package app

import (
	"cashflow/internal/app/endpoint"
	cors "cashflow/internal/middleware"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type App struct {
	server   *echo.Echo
	endpoint *endpoint.Endpoint
}

type AppValidator struct {
	validator *validator.Validate
}

func New() (*App, error) {
	app := &App{}
	app.server = echo.New()
	app.endpoint = endpoint.New()

	// Middleware
	app.server.Use(middleware.Recover(), cors.Origins())

	// Validator
	app.server.Validator = &AppValidator{validator: validator.New()}

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

func (cv *AppValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}
