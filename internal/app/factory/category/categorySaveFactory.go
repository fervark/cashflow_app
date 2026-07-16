package categorySaveFactory

import (
	categoryModel "cashflow/internal/app/model/category"
	"cashflow/internal/database"
	"log"
)

type Factory struct {
	Category categoryModel.Category
}

func Run(f Factory) any {
	// DB connection
	db := database.Open()
	if db.Error != nil {
		log.Println(db.Error)
		return db.Error
	}

	// Create category
	result := db.Create(&f.Category)
	if result.Error != nil {
		log.Fatalf("Error creating category: %s", result.Error)
		return result.Error
	}

	return result
}
