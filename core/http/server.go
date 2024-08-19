package http

import (
	"context"
	"net"
	"net/http"

	"github.com/hamstag/fintech/core/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewServer(lc fx.Lifecycle, cfg *config.Config, log *zap.Logger, r *Router) *http.Server {
	srv := &http.Server{Addr: cfg.Address, Handler: r}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
