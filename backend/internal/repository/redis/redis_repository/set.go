package redis_repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type UserSession struct {
	ID int64 `json:"id"`
	Role string `json:"role"`
	UserAgent string `json:"user_agent"`
}

func (r *UserRedisRepo) Set(ctx context.Context, key string, user UserSession, ttl time.Duration) error {
	sessionBytes, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "UserRedisRepo.Put")
	}

	_, err = r.db.Set(ctx, key, sessionBytes, ttl).Result()
	if err != nil {
		return errors.Wrap(err, "UserRedisRepo.Put")
	}
	return nil
}
