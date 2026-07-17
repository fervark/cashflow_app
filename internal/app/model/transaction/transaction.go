package transactionModel

import (
	"time"
)

const TransactionTypeIncome = "income"
const TransactionTypeExpense = "expense"

type Transaction struct {
	Code     string
	Type     string
	Price    string
	Date     time.Time
	Category string
}

type TransactionStats struct {
	UserId               int
	IncomeSum            string
	ExpenseSum           string
	DiffSum              string
	CategoryExpenseGroup string
	DateIncomeGroup      string
	DateExpenseGroup     string
}
