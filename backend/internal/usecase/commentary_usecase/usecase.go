package commentary_usecase

import (
	"context"
	"texts/config"
	"texts/internal/repository/postgres/commentary_repository"

	"github.com/avito-tech/go-transaction-manager/trm"
)

type (
	CommentaryRepo interface {
		InsertCommentary(ctx context.Context, commentary commentary_repository.InsertCommentaryReq) (int64, error)
		SelectCommentaries(ctx context.Context, filter commentary_repository.SelectCommentaryReq) ([]commentary_repository.SelectCommentaryResp, error)
	}
)

type CommentaryUC struct {
	cfg              *config.Config
	commentaryPGRepo CommentaryRepo
	trManager        trm.Manager
}

func NewCommentaryUC(
	cfg *config.Config,
	commentaryPGRepo CommentaryRepo,
	trManager trm.Manager,
) *CommentaryUC {
	return &CommentaryUC{
		cfg:              cfg,
		trManager:        trManager,
		commentaryPGRepo: commentaryPGRepo,
	}
}
