package post_delivery

import (
	"context"
	"texts/config"
	"texts/internal/usecase/post_usecase"

	"github.com/gofiber/fiber/v2"
)

type (
	PostUC interface {
		CreatePost(ctx context.Context, createPostReq post_usecase.CreatePostReq) (int64, error)
		GetPosts(ctx context.Context, getPostsReq post_usecase.GetPostsReq) ([]post_usecase.GetPostResp, error)
		UpdatePost(ctx context.Context, updatePostReq post_usecase.UpdatePostReq) error
	}

	Handlers interface {
		GetPosts() fiber.Handler
		CreatePost() fiber.Handler
		ApprovePost() fiber.Handler
		RejectPost() fiber.Handler
	}
)

type PostHandlers struct {
	postUC PostUC
	cfg    *config.Config
}

func NewPostHandler(postUC PostUC, cfg *config.Config) *PostHandlers {
	return &PostHandlers{
		cfg:    cfg,
		postUC: postUC,
	}
}
