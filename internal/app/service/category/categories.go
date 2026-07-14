package categories

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Categories struct {
}

func New() *Categories {
	return &Categories{}
}

func (s *Categories) List(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get category list."})
}

func (s *Categories) SetCategory(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Set category."})
}
