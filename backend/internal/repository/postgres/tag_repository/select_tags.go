package tag_repository

import (
	"context"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type SelectTagReq struct {
	TagIDs   []int64  `db:"tag_id"`
	TagNames    []string `db:"tag_name"`
	Deleted  []bool   `db:"deleted"`
}

type SelectTagResp struct {
	TagID   int64  `db:"tag_id"`
	TagName    string `db:"tag_name"`
	Deleted  bool   `db:"deleted"`
}

func getConditionsSelectTag(filter SelectTagReq) squirrel.And {
	var conditions squirrel.And
	if len(filter.TagIDs) != 0 {
		conditions = append(conditions, squirrel.Eq{
			user_quieries.TagIDColumnName: filter.TagIDs,
		})
	}
	if len(filter.TagNames) != 0 {
		conditions = append(conditions, squirrel.Eq{
			user_quieries.TagNameColumnName: filter.TagNames,
		})
	}
	if len(filter.Deleted) != 0 {
		conditions = append(conditions, squirrel.Eq{
			user_quieries.TagDeletedColumnName: filter.Deleted,
		})
	}
	return conditions
}
func (r *TagRepo) SelectTags(ctx context.Context, filter SelectTagReq) ([]SelectTagResp, error) {
	conditions := getConditionsSelectTag(filter)

	query, args, err := squirrel.Select(user_quieries.SelectTagColumns...).
		From(user_quieries.TagTable).
		Where(conditions).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "TagRepo.SelectUsers")
	}

	var tags []SelectTagResp
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.GetContext(
		ctx,
		&tags,
		query,
		args...,
	)

	if err != nil {
		return nil, errors.Wrap(err, "TagRepo.SelectUsers")
	}
	return tags, nil
}