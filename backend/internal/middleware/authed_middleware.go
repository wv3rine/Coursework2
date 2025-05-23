package middleware

import (
	"context"
	"texts/internal/domain"
	"texts/pkg/constants"
	"texts/pkg/cookie"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (m *MDWManager) AuthedMiddleware(role *string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeaders := domain.AuthHeaders{
			UserAgent:  c.Get(fiber.HeaderUserAgent),
			SessionKey: c.Cookies(constants.SessionKey),
		}

		if authHeaders.SessionKey == "" {
			return errors.New("no session")
		}

		cachedSession, err := m.userRedisRepo.Get(context.Background(), authHeaders.SessionKey)
		if err != nil {
			cookie.ClearCookie(c, constants.SessionKey, constants.CookieDomain)
			return err
		}

		if role != nil {
			if cachedSession.Role != *role {
				return errors.New("forbidden")
			}
		}

		authHeaders.UserID = cachedSession.ID
		c.Locals(authHeadersCtx, authHeaders)

		if err := c.Next(); err != nil {
			return errors.Wrap(err, "MDWManager.NonAuthedMiddleware")
		}
		return nil
	}
}
