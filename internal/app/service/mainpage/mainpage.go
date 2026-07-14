package mainpage

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type MainPage struct {
}

func New() *MainPage {
	return &MainPage{}
}

func (s *MainPage) Page(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Cashflow project main page"})
}
