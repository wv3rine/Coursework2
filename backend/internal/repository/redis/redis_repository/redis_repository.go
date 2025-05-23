package redis_repository

import (
	"texts/config"

	"github.com/redis/go-redis/v9"
)

type UserRedisRepo struct {
	db  *redis.Client
	cfg *config.Config
}

func NewUserRedisRepo(db *redis.Client, cfg *config.Config) *UserRedisRepo {
	return &UserRedisRepo{db: db, cfg: cfg}
}
