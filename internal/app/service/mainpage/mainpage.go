package mainpage

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Page(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Cashflow project main page"})
}
