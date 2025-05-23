package user_delivery

import (
	"context"
	"texts/internal/domain"
	"texts/internal/usecase/user_usecase"
	"texts/pkg/constants"
	"texts/pkg/cookie"
	"texts/pkg/reqvalidator"
	"time"

	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
)


type SignInReq struct {
	Login string `json:"login,required"`
	Password string `json:"password,required"`
}

type SingInRes struct {
	UserID int64 `json:"user_id"`
	Role string `json:"role"`
}

func (h *UserHandlers) SignIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "userHandler.SignIn"
		ctx := context.Background()

		authHeaders, ok := c.Locals("authHeaders").(domain.AuthHeaders)
		if !ok {
			return errors.New("no authHeaders")
		}

		request := SignUpReq{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(err, spanName)
		}

		session, err := h.userUC.SignIn(ctx, user_usecase.SignInUserReq{
			Login: request.Login,
			Password: request.Password,
			UserAgent: authHeaders.UserAgent,
		})
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		cookie.ClearCookie(c, constants.SessionKey, constants.CookieDomain)

		cookie.SetCookie(c, cookie.CookieData{
			Name:    constants.SessionKey,
			Value:   session.SessionKey,
			Expires: time.Now().Add(session.TTL),
			Domain:  constants.CookieDomain,
		})
		return c.JSON(fiber.Map{
			"data": SingUpRes{
				UserID: session.UserID,
				Role: session.Role,
			},
		})
	}
}