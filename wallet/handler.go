package wallet

import (
	"context"
	stdhttp "net/http"

	"github.com/hamstag/fintech/core/config"
	"github.com/hamstag/fintech/core/http"
	"github.com/redis/go-redis/v9"
)

func HTTPHandler(r *http.Router, s WalletService, redisClient *redis.Client, ctx context.Context, cfg *config.Config) {
	r.Api().Get("/wallet/hello", func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		redisClient.Incr(ctx, "Hello")
		w.Write([]byte("Sabaidee!"))
	})
}
