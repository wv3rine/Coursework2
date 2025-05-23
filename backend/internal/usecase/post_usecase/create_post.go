package post_usecase

import (
	"context"
	"texts/internal/repository/postgres/post_repository"

	"github.com/pkg/errors"
)

type CreatePostReq struct {
	Name     string
	Author   string
	Genre    string
	Content  string
	EditorId *int64
	TagId    int64
}

func (u *PostUC) CreatePost(ctx context.Context, createPostReq CreatePostReq) (int64, error) {
	spanName := "PostUC.CreatePost"

	postID, err := u.postPGRepo.InsertPost(ctx, post_repository.InsertPostReq{
		Name:     createPostReq.Name,
		Author:   createPostReq.Author,
		Genre:    createPostReq.Genre,
		Content:  createPostReq.Content,
		EditorId: createPostReq.EditorId,
		TagId:    createPostReq.TagId,
	})
	if err != nil {
		return 0, errors.Wrap(err, spanName)
	}

	return postID, nil
}
