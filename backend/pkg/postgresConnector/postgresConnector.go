package postgresConnector

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2"
)

type Settings struct {
	MaxOpenConns    int           `validate:"required,min=1"`
	ConnMaxLifetime time.Duration `validate:"required,min=1"`
	MaxIdleConns    int           `validate:"required,min=1"`
	ConnMaxIdleTime time.Duration `validate:"required,min=1"`
}

type Config struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	SSLMode  string `validate:"required"`
	Settings
	LogQuery bool
	AppName  string
	HookFunc func(query string, took int64, appName string)
}

type hooks struct {
	cfg *Config
}

func makeHook(cfg *Config) *hooks {
	return &hooks{cfg: cfg}
}

func (h *hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return context.WithValue(ctx, "begin", time.Now()), nil
}

func (h *hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	begin, ok := ctx.Value("begin").(time.Time)
	if !ok {
		return ctx, nil
	}
	go func(begin time.Time) {
		if h.cfg.HookFunc != nil {
			h.cfg.HookFunc(query, time.Since(begin).Milliseconds(), h.cfg.AppName)
		}
	}(begin)
	return ctx, nil
}

func GetConnection(cfg Config) (conn *sqlx.DB, err error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	if cfg.LogQuery {
		sql.Register("withHook", sqlhooks.Wrap(&stdlib.Driver{}, makeHook(&cfg)))
		conn, err = sqlx.Open("withHook", connectionURL)
	} else {
		conn, err = sqlx.Open("pgx", connectionURL)
	}
	if err != nil {
		return
	}

	conn.SetMaxOpenConns(cfg.Settings.MaxOpenConns)
	conn.SetConnMaxLifetime(cfg.Settings.ConnMaxLifetime * time.Second)
	conn.SetMaxIdleConns(cfg.Settings.MaxIdleConns)
	conn.SetConnMaxIdleTime(cfg.Settings.ConnMaxIdleTime * time.Second)

	if err = conn.Ping(); err != nil {
		return
	}

	return
}
