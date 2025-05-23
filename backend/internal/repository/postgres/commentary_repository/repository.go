package commentary_repository

import (
	"texts/config"
	"texts/pkg/connectiondatabase"

	"github.com/avito-tech/go-transaction-manager/sqlx"
)

type CommentaryRepo struct {
	cfg      *config.Config
	db       connectiondatabase.DBops
	txGetter *sqlx.CtxGetter
}

func NewCommentaryPGRepo(
	cfg *config.Config,
	db connectiondatabase.DBops,
	txGetter *sqlx.CtxGetter,
) *CommentaryRepo {
	return &CommentaryRepo{cfg: cfg, db: db, txGetter: txGetter}
}