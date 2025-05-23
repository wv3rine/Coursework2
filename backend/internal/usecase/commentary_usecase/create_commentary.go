package commentary_usecase

import (
	"context"
	"texts/internal/repository/postgres/commentary_repository"

	"github.com/pkg/errors"
)

type CreateCommentaryReq struct {
	UserId            int64
	CommentaryContent string
}

func (u *CommentaryUC) CreateCommentary(ctx context.Context, createCommentaryReq CreateCommentaryReq) (int64, error) {
	spanName := "CommentaryUC.CreateCommentary"

	commentaryID, err := u.commentaryPGRepo.InsertCommentary(ctx, commentary_repository.InsertCommentaryReq{
		UserID:            createCommentaryReq.UserId,
		CommentaryContent: createCommentaryReq.CommentaryContent,
	})
	if err != nil {
		return 0, errors.Wrap(err, spanName)
	}

	return commentaryID, nil
}
