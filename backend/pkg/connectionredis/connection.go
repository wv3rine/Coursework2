package connectionredis

import (
	"context"
	"errors"
	"texts/config"

	"github.com/redis/go-redis/v9"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewDatabase(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	client := redis.NewClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
