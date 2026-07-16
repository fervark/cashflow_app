package movement_stats

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type MovementStats struct {
}

func GetStatistic(ctx *echo.Context) error {
	
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Get cash movement statistic."})
}
