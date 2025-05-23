package commentary_repository

import (
	"context"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type InsertCommentaryReq struct {
	UserID            int64  `db:"user_id"`
	CommentaryContent string `db:"commentary_content"`
}

func (r *CommentaryRepo) InsertCommentary(ctx context.Context, commentary InsertCommentaryReq) (int64, error) {
	query, args, err := squirrel.Insert(user_quieries.CommentaryTable).
		Columns(user_quieries.InsertCommentaryColumns...).
		Values(
			commentary.UserID,
			commentary.CommentaryContent,
		).
		Suffix("RETURNING " + user_quieries.CommentaryTable).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "CommentaryRepo.CreateOne")
	}

	var commentaryID int64
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.QueryRowContext(ctx, query, args...).Scan(&commentaryID)
	if err != nil {
		return 0, errors.Wrap(err, "CommentaryRepo.CreateOne")
	}
	return commentaryID, nil
}
