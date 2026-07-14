package database

import (
	"cashflow/config"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Open() *sql.DB {
	conf := config.New()
	ctx := context.Background()
	conStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)

	dbConfig, err := pgxpool.ParseConfig(conStr)
	if err != nil {
		log.Fatalf("Failed to parse database config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		log.Fatalf("Failed to create database connection pool: %v", err)
	}
	defer pool.Close()

	fmt.Println("Database connection established successfully")
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	//var connect, err = pgx.Connect(context.Background(), conStr)
	//if err != nil {
	//	log.Fatal("Error database connection: ", err)
	//}
	//defer connect.Close(context.Background())

	return db
}
