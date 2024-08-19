package redis

import (
	"context"

	"github.com/hamstag/fintech/core/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.Config, ctx context.Context) (*redis.Client, error) {
	opts, err := redis.ParseURL(cfg.RedisURL)

	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	if err = client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
