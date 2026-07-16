package transactionModel

import (
	"time"
)

type Transaction struct {
	Code     string
	Type     string
	Price    string
	Date     time.Time
	Category string
}
