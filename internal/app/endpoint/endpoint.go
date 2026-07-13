package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Service interface {
	TransactionList() any
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (r *Endpoint) GetMain(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Start test!"})
}

//func GetTransactions(ctx *echo.Context) error {
//	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get transaction list."})
//}

//func GetCashMovementStats(ctx *echo.Context) error {
//	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get cash movement statistic."})
//}

//func SetTransaction(ctx *echo.Context) error {
//	return ctx.JSON(http.StatusOK, map[string]string{"message": "Set transaction."})
//}

//func SetCategory(ctx *echo.Context) error {
//	return ctx.JSON(http.StatusOK, map[string]string{"message": "Set category transaction."})
//}
