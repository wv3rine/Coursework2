package post_repository

import (
	"context"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type UpdatePostReq struct {
	PostId   int64
	Name     *string `db:"name"`
	Author   *string `db:"author"`
	Genre    *string `db:"genre"`
	Content  *string `db:"content"`
	EditorId *int64  `db:"editor_id"`
	TagId    *int64  `db:"tag_id"`
	Status   *string `db:"status"`
}

func getPostUpdateData(post UpdatePostReq) (squirrel.UpdateBuilder, error) {
	sb := squirrel.Update(user_quieries.PostTable)
	if post.Name != nil {
		sb = sb.Set(user_quieries.PostNameColumnName, post.Name)
	}
	if post.Author != nil {
		sb = sb.Set(user_quieries.AuthorColumnName, post.Author)
	}
	if post.Genre != nil {
		sb = sb.Set(user_quieries.GenreColumnName, post.Genre)
	}
	if post.Content != nil {
		sb = sb.Set(user_quieries.ContentColumnName, post.Content)
	}
	if post.EditorId != nil {
		sb = sb.Set(user_quieries.EditorIDColumnName, post.EditorId)
	}
	if post.TagId != nil {
		sb = sb.Set(user_quieries.TagIDColumnName, post.TagId)
	}
	if post.Status != nil {
		sb = sb.Set(user_quieries.StatusColumnName, post.Status)
	}
	return sb, nil
}

func (r *PostRepo) Update(ctx context.Context, post UpdatePostReq) error {
	sb, err := getPostUpdateData(post)
	if err != nil {
		return errors.Wrap(err, "PostRepo.Update")
	}

	query, args, err := sb.
		Where(squirrel.Eq{
			user_quieries.PostIDColumnName: post.PostId,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "PostRepo.Update")
	}

	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	_, err = tr.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "PostRepo.Update")
	}
	return nil
}
