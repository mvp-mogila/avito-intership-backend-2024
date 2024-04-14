package storage

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	Close() error
	Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error)
	Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	Begin(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

type Cache interface {
	Close() error
	Set(key string, value interface{}) error
	Get(key string) ([]byte, error)
}
