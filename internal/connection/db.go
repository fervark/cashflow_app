package connection

import (
	"cashflow/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Open() (*pgx.Conn, error) {
	conf := config.New()

	conn := &Connection{}
	conn.Host = conf.Database.Host
	conn.Port = conf.Database.Port
	conn.User = conf.Database.User
	conn.Password = conf.Database.Password
	conn.Name = conf.Database.Name

	conStr := fmt.Sprintf(	"host=%h port=%p user=%u dbname=%sn sslmode=disable",
		conn.Host, conn.Port, conn.User, conn.Name
	)

	var connect, err = pgx.Connect(context.Background(), conStr)
	if err != nil {
		log.Fatal(err)
	}

	return connect, nil
}
