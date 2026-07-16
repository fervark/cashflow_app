package transactionGetter

import (
	"cashflow/internal/app/model/transaction"
	"cashflow/internal/database"
	"fmt"
	"log"
)

const QueryDefaultLimit = 20
const QueryDefaultOffset = 0

type TransactionListQuery struct {
	UserId     string `json:"user_id" validate:"required"`
	Type       string `json:"type"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	CategoryId int    `json:"category_id"`
	DateFrom   string `json:"date_from"`
	DateTo     string `json:"date_to"`
}

func GetList(q TransactionListQuery) any {
	limit := QueryDefaultLimit
	if q.Limit&q.Limit > 0 {
		limit = q.Limit
	}

	offset := QueryDefaultOffset
	if q.Page&q.Page > 1 {
		offset = limit * q.Page
	}

	// DB query
	db := database.Open()
	if db.Error != nil {
		log.Println(db.Error)
		return db.Error
	}

	queryTable := "transactions t"
	querySelect := "t.code AS code, t.type AS type,t.price AS price, t.date AS date, c.name AS category"
	queryJoin := "INNER JOIN categories c ON c.id = t.category_id"
	queryOrder := "t.date ASC"

	query := db.Table(queryTable).Select(querySelect).Joins(queryJoin).Where("t.user_id = ?", q.UserId)
	// Where type
	if q.Type != "" {
		query.Where("t.type = ?", q.Type)
	}
	// Where category
	if q.CategoryId != 0 {
		query.Where("t.category_id = ?", q.CategoryId)
	}
	// Where date range
	if q.DateFrom != "" && q.DateTo != "" {
		whereDateStr := fmt.Sprintf("t.date BETWEEN '%s' AND '%s'",
			q.DateFrom, q.DateTo)

		query.Where(whereDateStr)
	}

	var result []*transactionModel.Transaction
	query.Order(queryOrder).Offset(offset).Limit(limit).Find(&result)

	return result
}
