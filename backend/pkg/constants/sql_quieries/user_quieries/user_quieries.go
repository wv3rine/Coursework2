package user_quieries

import (
	"fmt"
	"slices"
	"texts/pkg/constants/utils"
)

const (
	UserTable = "texts_schema.user"
	// User table fields
	UserIDColumnName      = "user_id"
	LoginColumnName       = "login"
	PasswordColumnName    = "password"
	RoleColumnName        = "role"
	UserDeletedColumnName = "deleted"

	PostTable = "texts_schema.post"
	// Post table fields
	PostIDColumnName      = "post_id"
	PostNameColumnName    = "name"
	AuthorColumnName      = "author"
	GenreColumnName       = "genre"
	ContentColumnName     = "content"
	EditorIDColumnName    = "editor_id"
	StatusColumnName      = "status"
	PostDeletedColumnName = "deleted"

	TagTable = "texts_schema.tag"
	// Tag table fields
	TagIDColumnName          = "tag_id"
	TagNameColumnName        = "tag_name"
	NormalizedNameColumnName = "normalized_name"
	TagDeletedColumnName     = "deleted"

	CommentaryTable = "texts_schema.commentary"
	// Commentary table fields
	CommentaryIDColumnName      = "commentary_id"
	CommentaryUserIDColumnName  = "user_id"
	CommentaryContentColumnName = "commentary_content"
	CommentaryDeletedColumnName = "deleted"
)

var (
	InsertUserColumns = []string{
		"login",
		"password",
		"role",
	}

	SelectUserColumns = []string{
		"user_id",
		"login",
		"password",
		"role",
		"deleted",
	}

	SelectPostColumns = slices.Concat(
		utils.WithPrefix([]string{
			PostIDColumnName,
			PostNameColumnName,
			AuthorColumnName,
			GenreColumnName,
			ContentColumnName,
			EditorIDColumnName,
			TagIDColumnName,
			StatusColumnName,
			PostDeletedColumnName,
		},
			PostTable,
		),
		[]string{
			fmt.Sprintf("%s.%s", UserTable, LoginColumnName),
			fmt.Sprintf("%s.%s", TagTable, TagNameColumnName),
		},
	)

	InsertPostColumns = []string{
		PostNameColumnName,
		AuthorColumnName,
		GenreColumnName,
		ContentColumnName,
		EditorIDColumnName,
		TagIDColumnName,
	}

	InsertTagColumns = []string{
		TagNameColumnName,
	}

	SelectTagColumns = []string{
		TagIDColumnName,
		TagNameColumnName,
		TagDeletedColumnName,
	}

	SelectCommentaryColumns = slices.Concat(
		utils.WithPrefix([]string{
			CommentaryIDColumnName,
			UserIDColumnName,
			CommentaryContentColumnName,
			PostIDColumnName,
			CommentaryDeletedColumnName,
		},
			CommentaryTable,
		),
		[]string{
			fmt.Sprintf("%s.%s", UserTable, LoginColumnName),
		},
	)

	InsertCommentaryColumns = []string{
		UserIDColumnName,
		CommentaryContentColumnName,
	}
)
