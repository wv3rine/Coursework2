package commentary_repository

import (
	"context"
	"fmt"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type SelectCommentaryReq struct {
	CommentaryIDs      []int64  `db:"commentary_id"`
	UserIDs            []int64  `db:"user_id"`
	CommentaryContents []string `db:"commentary_content"`
	PostIDs            []int64  `db:"post_id"`
	Deleted            []bool   `db:"deleted"`
}

type SelectCommentaryResp struct {
	CommentaryID      int64  `db:"commentary_id"`
	UserID            int64  `db:"user_id"`
	Login             string `db:"login"`
	CommentaryContent string `db:"commentary_content"`
	PostID            int64  `db:"post_id"`
	Deleted           bool   `db:"deleted"`
}

func getConditionsSelectCommentary(filter SelectCommentaryReq) squirrel.And {
	var conditions squirrel.And
	if len(filter.CommentaryIDs) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.CommentaryTable, user_quieries.CommentaryIDColumnName): filter.CommentaryIDs,
		})
	}
	if len(filter.UserIDs) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.UserTable, user_quieries.UserIDColumnName): filter.UserIDs,
		})
	}
	if len(filter.CommentaryContents) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.CommentaryTable, user_quieries.CommentaryContentColumnName): filter.CommentaryContents,
		})
	}
	if len(filter.PostIDs) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.CommentaryTable, user_quieries.PostIDColumnName): filter.PostIDs,
		})
	}
	if len(filter.Deleted) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.CommentaryTable, user_quieries.CommentaryDeletedColumnName): filter.Deleted,
		})
	}
	return conditions
}
func (r *CommentaryRepo) SelectCommentaries(ctx context.Context, filter SelectCommentaryReq) ([]SelectCommentaryResp, error) {
	conditions := getConditionsSelectCommentary(filter)

	query, args, err := squirrel.Select(user_quieries.SelectCommentaryColumns...).
		From(user_quieries.CommentaryTable).
		Join(fmt.Sprintf(
			"%s ON %s.%s = %s.%s",
			user_quieries.UserTable,
			user_quieries.CommentaryTable,
			user_quieries.UserIDColumnName,
			user_quieries.UserTable,
			user_quieries.UserIDColumnName,
		)).
		Where(conditions).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "CommentaryRepo.SelectUsers")
	}

	var Commentaries []SelectCommentaryResp
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.SelectContext(
		ctx,
		&Commentaries,
		query,
		args...,
	)

	if err != nil {
		return nil, errors.Wrap(err, "CommentaryRepo.SelectUsers")
	}
	return Commentaries, nil
}
