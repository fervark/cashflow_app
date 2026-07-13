package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	s := echo.New()
	fmt.Println("Server started")

	// *******************
	// ** Middlewares
	// *******************x
	s.Use(middleware.RequestLogger())
	s.Use(middleware.Recover())
	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// *******************
	// ** Routes
	// *******************x

	// Get routes
	s.GET("/", Handler)
	s.GET("/transaction-list", getTransactions)
	s.GET("/movement-stats", getCashMovementStats)

	// Set routes
	s.POST("/set/transaction", setTransaction)
	s.POST("/set/category", setCategory)

	// *******************
	// ** Run server
	// *******************x
	sc := echo.StartConfig{Address: ":8080"}
	if err := sc.Start(context.Background(), s); err != nil {
		s.Logger.Error("failed to start server", "error", err)
	}
}

func Handler(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Start test!"})
}

func getTransactions(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get transaction list."})
}

func getCashMovementStats(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get cash movement statistic."})
}

func setTransaction(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Set transaction."})
}

func setCategory(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Set category transaction."})
}
