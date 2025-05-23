package user_repository

import (
	"texts/config"
	"texts/pkg/connectiondatabase"

	"github.com/avito-tech/go-transaction-manager/sqlx"
)

type UserRepo struct {
	cfg      *config.Config
	db       connectiondatabase.DBops
	txGetter *sqlx.CtxGetter
}

func NewUserPGRepo(
	cfg *config.Config,
	db connectiondatabase.DBops,
	txGetter *sqlx.CtxGetter,
) *UserRepo {
	return &UserRepo{cfg: cfg, db: db, txGetter: txGetter}
}