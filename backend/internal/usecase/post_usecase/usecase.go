package post_usecase

import (
	"context"
	"texts/config"
	"texts/internal/repository/postgres/post_repository"

	"github.com/avito-tech/go-transaction-manager/trm"
)

type (
	PostRepo interface {
		InsertPost(ctx context.Context, post post_repository.InsertPostReq) (int64, error)
		SelectPost(ctx context.Context, filter post_repository.SelectPostReq) ([]post_repository.SelectPostResp, error)
		Update(ctx context.Context, post post_repository.UpdatePostReq) error
	}
)

type PostUC struct {
	cfg        *config.Config
	postPGRepo PostRepo
	trManager  trm.Manager
}

func NewPostUC(
	cfg *config.Config,
	postPGRepo PostRepo,
	trManager trm.Manager,
) *PostUC {
	return &PostUC{
		cfg:        cfg,
		trManager:  trManager,
		postPGRepo: postPGRepo,
	}
}
