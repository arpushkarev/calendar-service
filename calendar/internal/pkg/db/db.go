package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

type Query struct {
	Name     string
	QueryRaw string
}

func (db *DB) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	return db.pool.Query(ctx, q.QueryRaw, args...)
}

func (db *DB) CLose() error {
	if db.pool != nil {
		db.pool.Close()
	}

	return nil
}
