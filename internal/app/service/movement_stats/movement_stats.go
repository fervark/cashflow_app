package movement_stats

import (
	transactionGetter "cashflow/internal/app/getter/transaction"

	"github.com/labstack/echo/v5"
)

type MovementStats struct {
}

func GetStatistic(ctx *echo.Context) any {
	query := transactionGetter.TransactionStatsQuery{
		UserId:   ctx.FormValue("user_id"),
		DateFrom: ctx.FormValueOr("date_from", ""),
		DateTo:   ctx.FormValueOr("date_to", ""),
	}

	// Validate data
	if err := ctx.Validate(query); err != nil {
		return err
	}

	result := transactionGetter.GetStats(query)

	return result
}
