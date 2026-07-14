package movement_stats

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type MovementStats struct {
}

func New() *MovementStats {
	return &MovementStats{}
}

func (s *MovementStats) Stats(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get cash movement statistic."})
}
