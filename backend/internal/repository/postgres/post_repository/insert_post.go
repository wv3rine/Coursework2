package post_repository

import (
	"context"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type InsertPostReq struct {
	Name     string `db:"name"`
	Author   string `db:"author"`
	Genre    string `db:"genre"`
	Content  string `db:"content"`
	EditorId *int64 `db:"editor_id"`
	TagId    int64  `db:"tag_id"`
}

func (r *PostRepo) InsertPost(ctx context.Context, post InsertPostReq) (int64, error) {
	query, args, err := squirrel.Insert(user_quieries.PostTable).
		Columns(user_quieries.InsertUserColumns...).
		Values(
			post.Name,
			post.Author,
			post.Genre,
			post.Content,
			post.EditorId,
			post.TagId,
		).
		Suffix("RETURNING " + user_quieries.PostIDColumnName).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "PostRepo.CreateOne")
	}

	var postID int64
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.QueryRowContext(ctx, query, args...).Scan(&postID)
	if err != nil {
		return 0, errors.Wrap(err, "PostRepo.CreateOne")
	}
	return postID, nil
}
