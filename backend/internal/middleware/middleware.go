package middleware

import (
	"context"
	"texts/config"
	"texts/internal/repository/redis/redis_repository"
)

type (
	UserRedisRepo interface {
		Get(
			ctx context.Context,
			key string,
		) (redis_repository.UserSession, error)
	}
)

type MDWManager struct {
	cfg           *config.Config
	userRedisRepo UserRedisRepo
}

func NewMiddlewareManager(
	cfg *config.Config,
	userRedisRepo UserRedisRepo,
) *MDWManager {
	return &MDWManager{
		cfg:           cfg,
		userRedisRepo: userRedisRepo,
	}
}
