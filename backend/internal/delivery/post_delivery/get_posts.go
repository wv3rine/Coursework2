package post_delivery

import (
	"context"
	"texts/internal/usecase/post_usecase"
	"texts/pkg/constants/utils"
	"texts/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type GetPostsReq struct {
	PostIds   []int64  `json:"post_ids"`
	Names     []string `json:"names"`
	Authors   []string `json:"authors"`
	Genres    []string `json:"genres"`
	Contents  []string `json:"contents"`
	EditorIds []int64  `json:"editor_ids"`
	TagIds    []int64  `json:"tag_ids"`
	TagNames  []string `json:"tag_names"`
	Statuses  []string `json:"statuses"`
	Deleted   []bool   `json:"deleted"`
}

type GetPostRes struct {
	PostId      int64   `json:"post_id"`
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Genre       string  `json:"genre"`
	Content     string  `json:"content"`
	EditorId    *int64  `json:"editor_id"`
	EditorLogin *string `json:"login"`
	TagId       int64   `json:"tag_id"`
	TagName     string  `json:"tag_name"`
	Status      string  `json:"status"`
	Deleted     bool    `json:"deleted"`
}

type GetPostsRes struct {
	Posts []GetPostRes `json:"posts"`
}

func (h *PostHandlers) GetPosts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "PostHandlers.GetPosts"
		ctx := context.Background()

		request := GetPostsReq{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(err, spanName)
		}

		posts, err := h.postUC.GetPosts(ctx, post_usecase.GetPostsReq{
			PostIds:   request.PostIds,
			Names:     request.Names,
			Authors:   request.Authors,
			Genres:    request.Genres,
			Contents:  request.Contents,
			EditorIds: request.EditorIds,
			TagIds:    request.TagIds,
			TagNames:  request.TagNames,
			Statuses:  request.Statuses,
			Deleted:   request.Deleted,
		})
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": GetPostsRes{
				Posts: utils.MapArr(posts, func(post post_usecase.GetPostResp) GetPostRes {
					return GetPostRes{
						PostId:      post.PostId,
						Name:        post.Name,
						Author:      post.Author,
						Genre:       post.Genre,
						Content:     post.Content,
						EditorId:    post.EditorId,
						EditorLogin: post.EditorLogin,
						TagId:       post.TagId,
						TagName:     post.TagName,
						Status:      post.Status,
						Deleted:     post.Deleted,
					}
				}),
			},
		})
	}
}
