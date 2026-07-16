package transactionGetter

import (
	"cashflow/internal/app/model/transaction"
	"cashflow/internal/database"
	"log"
)

const QueryDefaultLimit = 20
const QueryDefaultOffset = 0

type TransactionListQuery struct {
	UserId string `json:"user_id" validate:"required"`
	Type   string `json:"type"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
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
	if q.Type != "" {
		query.Where("t.type = ?", q.Type)
	}

	var result []*transactionModel.Transaction
	query.Order(queryOrder).Offset(offset).Limit(limit).Find(&result)

	return result
}
