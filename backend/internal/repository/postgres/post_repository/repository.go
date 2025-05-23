package post_repository

import (
	"texts/config"
	"texts/pkg/connectiondatabase"

	"github.com/avito-tech/go-transaction-manager/sqlx"
)

type PostRepo struct {
	cfg      *config.Config
	db       connectiondatabase.DBops
	txGetter *sqlx.CtxGetter
}

func NewPostPGRepo(
	cfg *config.Config,
	db connectiondatabase.DBops,
	txGetter *sqlx.CtxGetter,
) *PostRepo {
	return &PostRepo{cfg: cfg, db: db, txGetter: txGetter}
}