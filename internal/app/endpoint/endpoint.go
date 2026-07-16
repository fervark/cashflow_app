package endpoint

import (
	categories "cashflow/internal/app/service/category"
	"cashflow/internal/app/service/mainpage"
	"cashflow/internal/app/service/movement_stats"
	"cashflow/internal/app/service/transactions"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Endpoint struct {
}

func New() *Endpoint {
	return &Endpoint{}
}

// MainPage Main
//
//	Page data
func (r *Endpoint) MainPage(ctx *echo.Context) error {
	res := mainpage.Page(ctx)
	return ctx.JSON(http.StatusOK, res)
}

// SetCategory Categories
// Create new category (POST)
func (r *Endpoint) SetCategory(ctx *echo.Context) error {
	res := categories.SetCategory(ctx)
	return ctx.JSON(http.StatusOK, res)
}

// GetTransactions Transactions
// Get transaction list
func (r *Endpoint) GetTransactions(ctx *echo.Context) error {
	res := transactions.List(ctx)
	return ctx.JSON(http.StatusOK, res)
}

// Set transaction
// Create new transaction (POST)
func (r *Endpoint) SetTransaction(ctx *echo.Context) error {
	res := transactions.SetTransaction(ctx)
	return ctx.JSON(http.StatusOK, res)
}

// GetCashMovementStats Stats
// Get movement statistics
func (r *Endpoint) GetCashMovementStats(ctx *echo.Context) error {
	res := movement_stats.GetStatistic(ctx)
	return ctx.JSON(http.StatusOK, res)
}
