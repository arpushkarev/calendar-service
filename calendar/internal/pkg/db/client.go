package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	DB() *DB
	Close() error
}

type client struct {
	db *DB
}

func NewClient(ctx context.Context, config *pgxpool.Config) (Client, error) {
	dbc, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	return &client{
		db: &DB{pool: dbc},
	}, nil
}

func (c *client) DB() *DB {
	return c.db
}

func (c *client) Close() error {
	if c != nil {
		c.db.CLose()
	}

	return nil
}
