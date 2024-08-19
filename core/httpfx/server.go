package httpfx

import (
	"context"
	"net"
	"net/http"

	"github.com/hamstag/fintech/core/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	ServerParams struct {
		fx.In

		LC     fx.Lifecycle
		Config *config.Config
		Logger *zap.Logger
		Router *Router
	}

	ServerResult struct {
		fx.Out

		Server *http.Server
	}
)

func NewServer(p ServerParams) ServerResult {
	srv := &http.Server{Addr: p.Config.Address, Handler: p.Router}
	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			p.Logger.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return ServerResult{Server: srv}
}
