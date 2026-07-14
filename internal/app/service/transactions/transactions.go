package transactions

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Transactions struct {
}

func New() *Transactions {
	return &Transactions{}
}

func (s *Transactions) List(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get transaction list."})
}

func (s *Transactions) SetTransaction(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Set transaction."})
}
