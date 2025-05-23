package middleware

import (
	"texts/internal/domain"
	"texts/pkg/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

const (
	authHeadersCtx = "authHeaders"
)

func (m *MDWManager) NonAuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeaders := domain.AuthHeaders{
			UserAgent: c.Get(fiber.HeaderUserAgent),
		}

		c.Locals(authHeadersCtx, authHeaders)

		if c.Cookies(constants.SessionKey) != "" {
			return errors.New("already has the session")
		}

		if err := c.Next(); err != nil {
			return errors.Wrap(err, "MDWManager.NonAuthedMiddleware")
		}
		return nil
	}
}
