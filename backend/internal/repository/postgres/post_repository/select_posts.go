package post_repository

import (
	"context"
	"fmt"
	"texts/pkg/constants/sql_quieries/user_quieries"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type SelectPostReq struct {
	PostIds   []int64
	Names     []string
	Authors   []string
	Genres    []string
	Contents  []string
	EditorIds []int64
	TagIds    []int64
	TagNames  []string
	Statuses  []string
	Deleted   []bool
}

type SelectPostResp struct {
	PostId      int64   `db:"post_id"`
	Name        string  `db:"name"`
	Author      string  `db:"author"`
	Genre       string  `db:"genre"`
	Content     string  `db:"content"`
	EditorId    *int64  `db:"editor_id"`
	EditorLogin *string `db:"login"`
	TagId       int64   `db:"tag_id"`
	TagName     string  `db:"tag_name"`
	Status      string  `db:"status"`
	Deleted     bool    `db:"deleted"`
}

func getConditionsSelectPost(filter SelectPostReq) squirrel.And {
	var conditions squirrel.And
	if len(filter.PostIds) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.PostIDColumnName): filter.PostIds,
		})
	}
	if len(filter.Names) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.PostNameColumnName): filter.Names,
		})
	}
	if len(filter.Authors) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.AuthorColumnName): filter.Authors,
		})
	}
	if len(filter.Genres) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.GenreColumnName): filter.Genres,
		})
	}
	if len(filter.EditorIds) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.EditorIDColumnName): filter.EditorIds,
		})
	}
	if len(filter.TagIds) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.TagIDColumnName): filter.TagIds,
		})
	}
	if len(filter.TagNames) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.TagTable, user_quieries.TagNameColumnName): filter.TagNames,
		})
	}
	if len(filter.Statuses) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.StatusColumnName): filter.Statuses,
		})
	}
	if len(filter.Deleted) != 0 {
		conditions = append(conditions, squirrel.Eq{
			fmt.Sprintf("%s.%s", user_quieries.PostTable, user_quieries.PostDeletedColumnName): filter.Deleted,
		})
	}
	return conditions
}
func (r *PostRepo) SelectPost(ctx context.Context, filter SelectPostReq) ([]SelectPostResp, error) {
	conditions := getConditionsSelectPost(filter)

	query, args, err :=
		// squirrel.Select(user_quieries.SelectPostColumns...).
		squirrel.Select([]string{
			"post.post_id AS post_id",
			"post.name AS name",
			"post.author AS author",
			"post.genre AS genre",
			"post.content AS content",
			"post.editor_id AS editor_id",
			"texts_schema.user.login AS login",
			"tag.tag_id AS tag_id",
			"tag.tag_name AS tag_name",
			"post.status AS status",
			"post.deleted AS deleted",
		}...).
			From(user_quieries.PostTable).
			LeftJoin(fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				user_quieries.UserTable,
				user_quieries.UserTable,
				user_quieries.UserIDColumnName,
				user_quieries.PostTable,
				user_quieries.EditorIDColumnName,
			)).
			Join(fmt.Sprintf(
				"%s ON %s.%s = %s.%s",
				user_quieries.TagTable,
				user_quieries.TagTable,
				user_quieries.TagIDColumnName,
				user_quieries.PostTable,
				user_quieries.TagIDColumnName,
			)).
			Where(conditions).
			PlaceholderFormat(squirrel.Dollar).
			ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "UserRepo.SelectUsers")
	}

	fmt.Println(query, args)
	var posts []SelectPostResp
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db.GetPool())
	err = tr.SelectContext(
		ctx,
		&posts,
		query,
		args...,
	)

	if err != nil {
		return nil, errors.Wrap(err, "UserRepo.SelectUsers")
	}
	return posts, nil
}
