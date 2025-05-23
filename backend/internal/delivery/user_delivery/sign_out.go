package user_delivery

import (
	"context"

	"texts/internal/usecase/user_usecase"
	"texts/pkg/constants"
	"texts/pkg/cookie"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)


func (h *UserHandlers) SignOut() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		sessionKey := c.Cookies(constants.SessionKey)
		if sessionKey == "" {
			return errors.New("no cookies")
		}

		err := h.userUC.SignOut(ctx, user_usecase.SignOutUserReq{
			SessionKey: sessionKey,
		})
		if err != nil {
			return err
		}

		cookie.ClearCookie(c, constants.SessionKey, constants.CookieDomain)

		return c.JSON(fiber.Map{
			"data": "Successfully Log out",
		})
	}
}