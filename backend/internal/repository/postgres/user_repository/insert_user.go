package user_repository

import (
	"context"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type InsertUserReq struct {
	Login    string `db:"login"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func (r *UserRepo) InsertUser(ctx context.Context, user InsertUserReq) (int64, error) {
	query, args, err := squirrel.Insert(user_quieries.UserTable).
		Columns(user_quieries.InsertUserColumns...).
		Values(
			user.Login,
			user.Password,
			user.Role,
		).
		Suffix("RETURNING " + user_quieries.UserIDColumnName).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "UserRepo.CreateOne")
	}

	var userID int64
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.QueryRowContext(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, errors.Wrap(err, "UserRepo.CreateOne")
	}
	return userID, nil
}