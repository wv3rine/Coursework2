package commentary_delivery

import (
	"context"
	"texts/internal/domain"
	"texts/internal/usecase/commentary_usecase"
	"texts/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CreateCommentaryReq struct {
	CommentaryContent string `json:"commentary_content"`
}

type CreateCommentaryResp struct {
	CommentaryID int64 `json:"commentary_id"`
}

func (h *CommentaryHandlers) CreateCommentary() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "CommentaryHandlers.CreateCommentary"
		ctx := context.Background()

		authHeaders, ok := c.Locals("authHeaders").(domain.AuthHeaders)
		if !ok {
			return errors.New("no authHeaders")
		}

		request := CreateCommentaryReq{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(err, spanName)
		}

		commentaryID, err := h.commentaryUC.CreateCommentary(ctx, commentary_usecase.CreateCommentaryReq{
			UserId:            authHeaders.UserID,
			CommentaryContent: request.CommentaryContent,
		})
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": CreateCommentaryResp{
				CommentaryID: commentaryID,
			},
		})
	}
}
