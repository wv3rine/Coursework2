package user_repository

import (
	"context"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type SelectUserReq struct {
	UserIDs []int64  `db:"user_id"`
	Logins  []string `db:"login"`
	Roles   []string `db:"role"`
	Deleted []bool   `db:"deleted"`
}

type SelectUserResp struct {
	UserID   int64  `db:"user_id"`
	Login    string `db:"login"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Deleted  bool   `db:"deleted"`
}

func getConditionsSelectUser(filter SelectUserReq) squirrel.And {
	var conditions squirrel.And
	if len(filter.UserIDs) != 0 {
		conditions = append(conditions, squirrel.Eq{
			user_quieries.UserIDColumnName: filter.UserIDs,
		})
	}
	if len(filter.Logins) != 0 {
		conditions = append(conditions, squirrel.Eq{
			user_quieries.LoginColumnName: filter.Logins,
		})
	}
	if len(filter.Roles) != 0 {
		conditions = append(conditions, squirrel.Eq{
			user_quieries.RoleColumnName: filter.Roles,
		})
	}
	if len(filter.Deleted) != 0 {
		conditions = append(conditions, squirrel.Eq{
			user_quieries.UserDeletedColumnName: filter.Deleted,
		})
	}
	return conditions
}
func (r *UserRepo) SelectUsers(ctx context.Context, filter SelectUserReq) ([]SelectUserResp, error) {
	conditions := getConditionsSelectUser(filter)

	query, args, err := squirrel.Select(user_quieries.SelectUserColumns...).
		From(user_quieries.UserTable).
		Where(conditions).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "UserRepo.SelectUsers")
	}

	var users []SelectUserResp
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.SelectContext(
		ctx,
		&users,
		query,
		args...,
	)

	if err != nil {
		return nil, errors.Wrap(err, "UserRepo.SelectUsers")
	}
	return users, nil
}
