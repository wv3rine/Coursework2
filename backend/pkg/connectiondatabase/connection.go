package connectiondatabase

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // import the pq driver
)

type DBops interface {
	StartTransaction() (*sqlx.Tx, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetPool() *sqlx.DB
	Close() error
}

type Database struct {
	db *sqlx.DB
}

func (o *Database) Close() error {
	return o.db.Close()
}
func (o *Database) StartTransaction() (*sqlx.Tx, error) {
	return o.db.Beginx()
}
func (o *Database) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return o.db.GetContext(ctx, dest, query, args...)
}
func (o *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return o.db.QueryRowContext(ctx, query, args...)
}
func (o *Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return o.db.ExecContext(ctx, query, args...)
}
func (o *Database) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return o.db.SelectContext(ctx, dest, query, args...)
}
func (o *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return o.db.QueryContext(ctx, query, args...)
}
func (o *Database) GetPool() *sqlx.DB {
	return o.db
}

func NewDB(db *sqlx.DB) *Database {
	return &Database{db: db}
}
