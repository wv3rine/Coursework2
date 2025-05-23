package tag_repository

import (
	"texts/config"
	"texts/pkg/connectiondatabase"

	"github.com/avito-tech/go-transaction-manager/sqlx"
)

type TagRepo struct {
	cfg      *config.Config
	db       connectiondatabase.DBops
	txGetter *sqlx.CtxGetter
}

func NewTagPGRepo(
	cfg *config.Config,
	db connectiondatabase.DBops,
	txGetter *sqlx.CtxGetter,
) *TagRepo {
	return &TagRepo{cfg: cfg, db: db, txGetter: txGetter}
}