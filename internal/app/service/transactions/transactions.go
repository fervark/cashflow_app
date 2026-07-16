package transactions

import (
	transactionGetter "cashflow/internal/app/getter/transaction"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Transactions struct {
}

func List(ctx *echo.Context) any {
	page, _ := strconv.Atoi(ctx.FormValue("page"))
	limit, _ := strconv.Atoi(ctx.FormValue("limit"))

	query := transactionGetter.TransactionListQuery{
		UserId: ctx.FormValue("user_id"),
		Type:   ctx.FormValue("type"),
		Page:   page,
		Limit:  limit,
	}

	// Validate data
	if err := ctx.Validate(query); err != nil {
		return err
	}

	result := transactionGetter.GetList(query)

	return result
}

func SetTransaction(ctx *echo.Context) any {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Set transaction."})
}
