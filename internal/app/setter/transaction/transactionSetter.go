package transactionSetter

import "time"

type TransactionSetQuery struct {
	UserId     int    `json:"user_id" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Price      string `json:"price" validate:"required"`
}

type TransactionSet struct {
	Code       string    `json:"code" validate:"required"`
	UserId     int       `json:"user_id" validate:"required"`
	CategoryId int       `json:"category_id" validate:"required"`
	Type       string    `json:"type" validate:"required"`
	Price      string    `json:"price" validate:"required"`
	Date       time.Time `json:"date" validate:"required"`
}
