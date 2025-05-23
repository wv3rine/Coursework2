package post_delivery

import (
	"context"
	"texts/internal/domain"
	"texts/internal/usecase/post_usecase"
	"texts/pkg/constants"
	"texts/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type UpdatePostReq struct {
	PostId int64 `json:"post_id,required"`
}

func (h *PostHandlers) ApprovePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "PostHandlers.ApprovePost"
		ctx := context.Background()

		authHeaders, ok := c.Locals("authHeaders").(domain.AuthHeaders)
		if !ok {
			return errors.New("no authHeaders")
		}

		request := UpdatePostReq{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(err, spanName)
		}

		err := h.postUC.UpdatePost(ctx, post_usecase.UpdatePostReq{
			PostId:   request.PostId,
			EditorId: &authHeaders.UserID,
			Status:   &constants.StatusApproved,
		})
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": "successfuly updated",
		})
	}
}
