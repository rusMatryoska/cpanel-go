package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/exp/slog"
)

type Database struct {
	DBConnURL string
	ConnPool  *pgxpool.Pool
	Log       *slog.Logger
}

func (db *Database) GetDBConnection(ctx context.Context) (*pgxpool.Pool, string) {
	pool, err := pgxpool.Connect(ctx, db.DBConnURL)
	if err != nil {
		return nil, fmt.Sprint(err)
	} else {
		return pool, ""
	}
}

func (db *Database) Ping(ctx context.Context) error {
	err := db.ConnPool.Ping(ctx)
	return err
}

func (db *Database) Exec(ctx context.Context, query string) (pgconn.CommandTag, error) {
	res, err := db.ConnPool.Exec(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TODO: optimization work with postgresql
