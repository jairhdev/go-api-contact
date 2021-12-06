package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type impl interface {
	NewConnect() error
	GetCtx() context.Context
	GetConn() *pgxpool.Pool
	CloseConn()
}

type service struct {
	impl
}

func NewService(database impl) service {
	return service{database}
}
