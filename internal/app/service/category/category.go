package category

import (
	categorySaveFactory "cashflow/internal/app/factory/category"
	categoryModel "cashflow/internal/app/model/category"

	"github.com/labstack/echo/v5"
)

func SetCategory(ctx *echo.Context) any {
	category := categoryModel.Category{Name: ctx.FormValue("name")}

	// Validate data
	if err := ctx.Validate(category); err != nil {
		return err
	}

	// Save category factor
	data := categorySaveFactory.Factory{Category: category}
	result := categorySaveFactory.Run(data)

	return result
}
