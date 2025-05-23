package user_usecase

import (
	"context"
	"texts/config"
	"texts/internal/repository/postgres/user_repository"
	"texts/internal/repository/redis/redis_repository"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
)


type (
	UserRepo interface {
		InsertUser(ctx context.Context, user user_repository.InsertUserReq) (int64, error)
		SelectUsers(ctx context.Context, filter user_repository.SelectUserReq) ([]user_repository.SelectUserResp, error)
	}

	UserRedisRepo interface {
		Get(
			ctx context.Context,
			key string,
		) (redis_repository.UserSession, error)
		Set(ctx context.Context, key string, user redis_repository.UserSession, ttl time.Duration) error
		Delete(ctx context.Context, key string) error
	}

)

type repos struct {
	userPGRepo       UserRepo
	userRedisRepo    UserRedisRepo
}

type UserUC struct {
	cfg *config.Config
	repos
	trManager trm.Manager
}

func NewUserUC(
	cfg *config.Config,
	userPGRepo UserRepo,
	userRedisRepo    UserRedisRepo,
	trManager trm.Manager,
) *UserUC {
	return &UserUC{
		cfg:       cfg,
		trManager: trManager,
		repos: repos{
			userPGRepo:       userPGRepo,
			userRedisRepo:    userRedisRepo,
		},
	}
}