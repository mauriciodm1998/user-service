package repository

import (
	"context"
	"errors"
	"user-service/internal/config"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	cfg           = &config.Cfg
	ErrorNotFound = errors.New("entity not found")
)

func New() PgxIface {
	db, err := pgxpool.Connect(context.Background(), cfg.DB.ConnectionString)
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) (err error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Ping(context.Context) error
}
