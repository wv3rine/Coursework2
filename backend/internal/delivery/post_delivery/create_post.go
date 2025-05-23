package post_delivery

import (
	"context"
	"texts/internal/usecase/post_usecase"
	"texts/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CreatePostReq struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Genre   string `json:"genre"`
	Content string `json:"content"`
	TagId   int64  `json:"tag_id"`
}

type CreatePostResp struct {
	PostID int64 `json:"post_id"`
}

func (h *PostHandlers) CreatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "PostHandlers.CreatePost"
		ctx := context.Background()

		request := CreatePostReq{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(err, spanName)
		}

		postID, err := h.postUC.CreatePost(ctx, post_usecase.CreatePostReq{
			Name:    request.Name,
			Author:  request.Author,
			Genre:   request.Genre,
			Content: request.Content,
			TagId:   request.TagId,
		})
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": CreatePostResp{
				PostID: postID,
			},
		})
	}
}
