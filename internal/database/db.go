package database

import (
	"cashflow/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() *gorm.DB {
	conf := config.New()
	conStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Name,
		conf.Database.Password)

	db, _ := gorm.Open(postgres.Open(conStr), &gorm.Config{})

	return db
}
