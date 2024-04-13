package storage

import (
	"context"
	"database/sql"
)

type Storage interface {
	Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error)
	Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, q string, args ...interface{}) error
}
