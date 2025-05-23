package user_delivery

import (
	"context"
	"texts/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CheckSessionRes struct {
	UserID int64 `json:"user_id"`
	Role string `json:"role"`
}

func (h *UserHandlers) CheckSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "userHandler.CheckSession"
		ctx := context.Background()

		authHeaders, ok := c.Locals("authHeaders").(domain.AuthHeaders)
		if !ok {
			return errors.New("no authHeaders")
		}

		session, err := h.userUC.CheckSession(ctx,authHeaders.SessionKey)
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": SingUpRes{
				UserID: session.UserID,
				Role: session.Role,
			},
		})
	}
}