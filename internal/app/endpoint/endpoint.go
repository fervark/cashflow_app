package endpoint

import (
	categories "cashflow/internal/app/service/category"
	"cashflow/internal/app/service/mainpage"
	"cashflow/internal/app/service/movement_stats"
	"cashflow/internal/app/service/transactions"

	"github.com/labstack/echo/v5"
)

type Endpoint struct {
}

func New() *Endpoint {
	return &Endpoint{}
}

// MainPage Main
func (r *Endpoint) MainPage(ctx *echo.Context) error {
	return mainpage.New().Page(ctx)
}

// SetCategory Categories
func (r *Endpoint) SetCategory(ctx *echo.Context) error {
	return categories.New().SetCategory(ctx)
}

// GetTransactions Transactions
func (r *Endpoint) GetTransactions(ctx *echo.Context) error {
	return transactions.New().List(ctx)
}

func (r *Endpoint) SetTransaction(ctx *echo.Context) error {
	return transactions.New().SetTransaction(ctx)
}

// GetCashMovementStats Stats
func (r *Endpoint) GetCashMovementStats(ctx *echo.Context) error {
	return movement_stats.New().Stats(ctx)
}
