package commentary_delivery

import (
	"context"
	"texts/internal/usecase/commentary_usecase"
	"texts/pkg/constants/utils"
	"texts/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type GetCommentariesReq struct {
	PostID int64 `json:"post_id"`
}

type GetCommentaryResp struct {
	CommentaryID      int64  `json:"commentary_id"`
	UserID            int64  `json:"user_id"`
	Login             string `json:"login"`
	CommentaryContent string `json:"commentary_content"`
	Deleted           bool   `json:"deleted"`
}

type GetCommentariesResp struct {
	Commentaries []GetCommentaryResp `json:"commentaries"`
}

func (h *CommentaryHandlers) GetCommentaries() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "CommentaryHandlers.GetCommentaries"
		ctx := context.Background()

		request := GetCommentariesReq{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(err, spanName)
		}

		commentaries, err := h.commentaryUC.GetCommentaries(ctx, commentary_usecase.GetCommentarysReq{
			PostIDs: []int64{request.PostID},
		})
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": GetCommentariesResp{
				Commentaries: utils.MapArr(commentaries, func(commentary commentary_usecase.GetCommentaryResp) GetCommentaryResp {
					return GetCommentaryResp{
						CommentaryID:      commentary.CommentaryID,
						UserID:            commentary.UserID,
						Login:             commentary.Login,
						CommentaryContent: commentary.CommentaryContent,
						Deleted:           commentary.Deleted,
					}
				}),
			},
		})
	}
}
