package tag_repository

import (
	"context"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type InsertTagReq struct {
	TagName    string `db:"tag_name"`
}

func (r *TagRepo) InsertTag(ctx context.Context, tag InsertTagReq) (int64, error) {
	query, args, err := squirrel.Insert(user_quieries.TagTable).
		Columns(user_quieries.InsertTagColumns...).
		Values(
			tag.TagName,
		).
		Suffix("RETURNING " + user_quieries.TagIDColumnName).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "TagRepo.CreateOne")
	}

	var tagID int64
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.QueryRowContext(ctx, query, args...).Scan(&tagID)
	if err != nil {
		return 0, errors.Wrap(err, "TagRepo.CreateOne")
	}
	return tagID, nil
}