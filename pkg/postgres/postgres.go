package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // postgres driver
	"github.com/jmoiron/sqlx"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/config"
	st "github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/storage"
)

type PgxDatabase struct {
	dsn    string
	client *sqlx.DB
}

func NewPgxDatabase(cfg config.PostgresConfig) (st.Database, error) {
	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		hostPort,
		cfg.Database,
		cfg.Sslmode,
	)
	dbClient, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	log.Println("postgres connection opened ...")
	dbClient.SetMaxOpenConns(cfg.MaxOpenConns)
	dbClient.SetConnMaxIdleTime(time.Second * time.Duration(cfg.MaxIdleTime))
	return &PgxDatabase{
		dsn:    dsn,
		client: dbClient,
	}, nil
}

func (db *PgxDatabase) Close() error {
	log.Println("postgres connection closing...")
	return db.client.Close()
}

func (db *PgxDatabase) Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error) {
	log.Printf("QUERY: %s\nARGS: %v\n", q, args)
	return db.client.ExecContext(ctx, sqlx.Rebind(sqlx.DOLLAR, q), args...)
}

func (db *PgxDatabase) Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error {
	log.Printf("QUERY: %s\nARGS: %v\n", q, args)
	return db.client.GetContext(ctx, dest, sqlx.Rebind(sqlx.DOLLAR, q), args...)
}

func (db *PgxDatabase) Select(ctx context.Context, dest interface{}, q string, args ...interface{}) error {
	log.Printf("QUERY: %s\nARGS: %v\n", q, args)
	return db.client.SelectContext(ctx, dest, sqlx.Rebind(sqlx.DOLLAR, q), args...)
}

func (db *PgxDatabase) Begin(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return db.client.BeginTxx(ctx, opts)
}
