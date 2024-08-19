package wallet

import (
	"context"
	"net/http"

	"github.com/hamstag/fintech/core/config"
	"github.com/hamstag/fintech/core/db"
	"github.com/hamstag/fintech/core/httpfx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func HTTPHandler(r *httpfx.Router, log *zap.Logger, s WalletService, redisClient *redis.Client, ctx context.Context) {
	r.Api().Get("/wallet/hello", func(w http.ResponseWriter, r *http.Request) {
		cfg, _ := config.New()
		db, _ := db.New(db.Params{
			Config: cfg.Config,
		})
		db.Database.FindByID("Hello")
		redisClient.Incr(ctx, "Hello")
		w.Write([]byte("Sabaidee!"))
	})
}
