package post_usecase

import (
	"context"
	"texts/internal/repository/postgres/post_repository"

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

func (u *PostUC) UpdatePost(ctx context.Context, updatePostReq UpdatePostReq) error {
	spanName := "PostUC.UpdatePost"

	err := u.postPGRepo.Update(ctx, post_repository.UpdatePostReq{
		PostId:   updatePostReq.PostId,
		Name:     updatePostReq.Name,
		Author:   updatePostReq.Author,
		Genre:    updatePostReq.Genre,
		Content:  updatePostReq.Content,
		EditorId: updatePostReq.EditorId,
		TagId:    updatePostReq.TagId,
		Status:   updatePostReq.Status,
	})
	if err != nil {
		return errors.Wrap(err, spanName)
	}

	return nil
}
