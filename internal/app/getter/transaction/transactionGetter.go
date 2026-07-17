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

type TransactionStatsQuery struct {
	UserId   string `json:"user_id" validate:"required"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
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
		whereDateStr := fmt.Sprintf("t.date::date BETWEEN '%s' AND '%s'",
			q.DateFrom, q.DateTo)

		query.Where(whereDateStr)
	}

	var result []*transactionModel.Transaction
	query.Order(queryOrder).Offset(offset).Limit(limit).Find(&result)

	return result
}

func GetStats(q TransactionStatsQuery) any {

	// Where query
	whereStr := fmt.Sprintf("t.user_id = %s", q.UserId)
	if q.DateFrom != "" && q.DateTo != "" {
		whereStr = fmt.Sprintf("%s AND t.date::date BETWEEN '%s' AND '%s'",
			whereStr, q.DateFrom, q.DateTo)
	}

	query := fmt.Sprintf(`
		with _t as (
			select *
			from transactions t 
			where %s
		)
		, _ti as (
			select t.*
			from _t t
			where t.type = 'income'
		)
		, _te as (
			select t.*
			from _t t
			where t.type = 'expense'
		)
		, _teg as (
			select 
				t.category_id as category_id,
				sum(t.price) as value
			from _te t
			group by t.category_id 
		)
		, _tdg as (
			select 
				t.type as type,
			    
				sum(t.price) as sum
			from _t t
			group by t.date::date, t.type
		)
		select 
			t.user_id as user_id,
			CASE WHEN any_value(ti.id) is not null 
				THEN sum(ti.price::numeric) 
				ELSE 0 
				end as income_sum,
			CASE WHEN any_value(te.id) is not null 
				THEN sum(te.price::numeric) 
				ELSE 0 	
				end as expense_sum,
			CASE WHEN any_value(ti.id) is not null and any_value(te.id) is not null 
				THEN SUM(ti.price::numeric) - SUM(te.price::numeric) 
				ELSE 0 
				end as diff_sum,
			( select json_agg(teg.*) as value from _teg teg ) as category_expense_group,
			( select json_agg(tdg.*) as value from _tdg tdg where tdg.type = 'income' ) as date_income_group,
			( select json_agg(tdg.*) as value from _tdg tdg where tdg.type = 'expense' ) as date_expense_group
		from _t t
		left join _ti ti on ti.id = t.id
		left join _te te on te.id = t.id
		group by t.user_id
		`, whereStr)

	db := database.SqlOpen()
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	result := []transactionModel.TransactionStats{}
	for rows.Next() {
		var data transactionModel.TransactionStats
		if err := rows.Scan(&data.UserId, &data.IncomeSum, &data.ExpenseSum, &data.DiffSum, &data.DateIncomeGroup, &data.CategoryExpenseGroup, &data.DateExpenseGroup); err != nil {
			log.Fatal(err)
		}

		result = append(result, data)
	}

	return result
}
