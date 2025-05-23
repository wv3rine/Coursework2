package commentary_usecase

import (
	"context"
	"texts/internal/repository/postgres/commentary_repository"
	"texts/pkg/constants/utils"

	"github.com/pkg/errors"
)

type GetCommentarysReq struct {
	CommentaryIDs      []int64
	UserIDs            []int64
	CommentaryContents []string
	PostIDs            []int64
	Deleted            []bool
}

type GetCommentaryResp struct {
	CommentaryID      int64
	UserID            int64
	Login             string
	CommentaryContent string
	PostID            int64
	Deleted           bool
}

func (u *CommentaryUC) GetCommentaries(ctx context.Context, getCommentarysReq GetCommentarysReq) ([]GetCommentaryResp, error) {
	spanName := "CommentaryUC.GetCommentaries"

	commentarys, err := u.commentaryPGRepo.SelectCommentaries(ctx, commentary_repository.SelectCommentaryReq{
		CommentaryIDs:      getCommentarysReq.CommentaryIDs,
		UserIDs:            getCommentarysReq.UserIDs,
		CommentaryContents: getCommentarysReq.CommentaryContents,
		PostIDs:            getCommentarysReq.PostIDs,
		Deleted:            getCommentarysReq.Deleted,
	})
	if err != nil {
		return []GetCommentaryResp{}, errors.Wrap(err, spanName)
	}

	return utils.MapArr(commentarys, func(commentary commentary_repository.SelectCommentaryResp) GetCommentaryResp {
		return GetCommentaryResp{
			CommentaryID:      commentary.CommentaryID,
			UserID:            commentary.UserID,
			Login:             commentary.Login,
			CommentaryContent: commentary.CommentaryContent,
			PostID:            commentary.PostID,
			Deleted:           commentary.Deleted,
		}
	}), nil
}
