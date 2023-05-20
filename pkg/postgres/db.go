package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func OpenDB(host string, port int, user, password, dbname string) (*pgxpool.Pool, error) {
	// postgres://postgres:postgres@localhost:5432/greenlight?sslmode=disable
	poolCfg, err := configurePool(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, host, port, dbname))
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	// Return the sql.DB connection pool.
	return conn, nil
}

func configurePool(dsn string) (*pgxpool.Config, error) {
	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	return poolCfg, nil
}
