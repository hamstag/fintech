package core

import (
	"context"

	"github.com/hamstag/fintech/core/config"
	"github.com/hamstag/fintech/core/db"
	"github.com/hamstag/fintech/core/http"
	"github.com/hamstag/fintech/core/redis"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("core",
	fx.Provide(
		fx.Annotate(
			func() context.Context {
				return context.Background()
			},
			fx.As(new(context.Context)),
		),
		zap.NewProduction,
	),
	config.Module,
	redis.Module,
	db.Module,
	http.Module,
)
