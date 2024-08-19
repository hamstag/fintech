package redisfx

import (
	"context"

	"github.com/hamstag/fintech/core/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Config  *config.Config
	Context context.Context
}

type Result struct {
	fx.Out

	Client *redis.Client
}

func New(p Params) (Result, error) {
	opts, err := redis.ParseURL(p.Config.RedisURL)

	if err != nil {
		return Result{}, err
	}

	client := redis.NewClient(opts)

	if err = client.Ping(p.Context).Err(); err != nil {
		return Result{}, err
	}

	return Result{Client: client}, nil
}
