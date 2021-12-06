package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ctx = context.Background()
var conn *pgxpool.Pool

// **********************************

func (data db) NewConnect() error {
	var err error

	if conn == nil {
		url := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			os.Getenv("DB_POSTGRES_USER"),
			os.Getenv("DB_POSTGRES_PWD"),
			os.Getenv("DB_POSTGRES_HOST"),
			os.Getenv("DB_POSTGRES_PORT"),
			os.Getenv("DB_POSTGRES_NAME"),
		)

		conn, err = pgxpool.Connect(ctx, url)
	}
	return err
}

func (data db) GetCtx() context.Context {
	return ctx
}

func (data db) GetConn() *pgxpool.Pool {
	return conn
}

func (data db) CloseConn() {
	conn.Close()
}
