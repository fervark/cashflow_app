package database

import (
	"cashflow/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

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

func SqlOpen() *sql.DB {
	conf := config.New()
	conStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	return db
}
