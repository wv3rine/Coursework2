package user_delivery

import (
	"context"
	"texts/config"
	"texts/internal/usecase/user_usecase"

	"github.com/gofiber/fiber/v2"
)

type (
	UserUC interface {
		SignUp(ctx context.Context, user user_usecase.SignUpUserReq) (user_usecase.Session, error)
		SignIn(ctx context.Context, signInUserReq user_usecase.SignInUserReq) (user_usecase.Session, error)
		CheckSession(
			ctx context.Context,
			sessionKey string,
		) (user_usecase.Session, error)
		SignOut(ctx context.Context, signOutUserReq user_usecase.SignOutUserReq) error
	}

	Handlers interface {
		CheckSession() fiber.Handler
		SignIn() fiber.Handler
		SignUp() fiber.Handler
		SignOut() fiber.Handler
	}
)

type UserHandlers struct {
	userUC UserUC
	cfg    *config.Config
}

func NewUserHandler(userUC UserUC, cfg *config.Config) *UserHandlers {
	return &UserHandlers{
		cfg:    cfg,
		userUC: userUC,
	}
}
