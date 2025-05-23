package redis_repository

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)


func (r *UserRedisRepo) Get(
	ctx context.Context,
	key string,
) (UserSession, error) {
	result := UserSession{}

	valueString, err := r.db.Get(ctx, key).Result()
	if err != nil {
		return result, errors.Wrap(err, "UserRedisRepo.Get")
	}
	err = json.Unmarshal([]byte(valueString), &result)
	if err != nil {
		return result, errors.Wrap(err, "UserRedisRepo.Get")
	}
	return result, nil
}