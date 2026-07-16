package transactions

import (
	transactionSaveFactory "cashflow/internal/app/factory/transaction"
	transactionGetter "cashflow/internal/app/getter/transaction"
	transactionSetter "cashflow/internal/app/setter/transaction"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Transactions struct {
}

func List(ctx *echo.Context) any {
	page, _ := strconv.Atoi(ctx.FormValue("page"))
	limit, _ := strconv.Atoi(ctx.FormValue("limit"))
	categoryId, _ := strconv.Atoi(ctx.FormValue("category_id"))

	query := transactionGetter.TransactionListQuery{
		UserId:     ctx.FormValue("user_id"),
		Type:       ctx.FormValue("type"),
		DateFrom:   ctx.FormValueOr("date_from", ""),
		DateTo:     ctx.FormValueOr("date_to", ""),
		CategoryId: categoryId,
		Page:       page,
		Limit:      limit,
	}

	// Validate data
	if err := ctx.Validate(query); err != nil {
		return err
	}

	result := transactionGetter.GetList(query)

	return result
}

func SetTransaction(ctx *echo.Context) any {
	userId, _ := strconv.Atoi(ctx.FormValue("user_id"))
	categoryId, _ := strconv.Atoi(ctx.FormValue("category_id"))

	transaction := transactionSetter.TransactionSetQuery{
		UserId:     userId,
		CategoryId: categoryId,
		Type:       ctx.FormValue("type"),
		Price:      ctx.FormValue("price"),
	}

	// Validate data
	if err := ctx.Validate(transaction); err != nil {
		return err
	}

	// Save transaction factor
	data := transactionSaveFactory.Factory{Transaction: transaction}
	transactionSaveFactory.Run(data)

	return "success"
}
