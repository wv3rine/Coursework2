package redis_repository

import (
	"context"

	"github.com/pkg/errors"
)


func (r *UserRedisRepo) Delete(ctx context.Context, key string) error {
	_, err := r.db.Del(ctx, key).Result()
	if err != nil {
		return errors.Wrap(err, "userRedisRepo.Delete")
	}

	return nil
}