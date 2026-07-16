package transactionSaveFactory

import (
	transactionSetter "cashflow/internal/app/setter/transaction"
	"cashflow/internal/database"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type Factory struct {
	Transaction transactionSetter.TransactionSetQuery
}

func Run(f Factory) any {
	// DB connection
	db := database.Open()
	if db.Error != nil {
		log.Println(db.Error)
		return db.Error
	}

	// Create transaction
	code := generateHashCode(f)
	data := transactionSetter.TransactionSet{
		Code:       code,
		UserId:     f.Transaction.UserId,
		CategoryId: f.Transaction.CategoryId,
		Type:       f.Transaction.Type,
		Price:      f.Transaction.Price,
		Date:       time.Now(),
	}

	result := db.Table("transactions").Create(&data)
	if result.Error != nil {
		log.Fatalf("Error creating transaction: %s", result.Error)
		return result.Error
	}

	return result
}

func generateHashCode(f Factory) string {
	date := time.Now()
	key := fmt.Sprintf("U%s D%s",
		f.Transaction.UserId, date.Format("2026-01-01 00:00:00"))

	h := sha1.New()
	h.Write([]byte(key))
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	return sha1Hash
}
