package httpfx

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/hamstag/fintech/core/config"
	"go.uber.org/fx"
)

type (
	Router struct {
		*chi.Mux

		api *chi.Mux
	}

	RouterParams struct {
		fx.In

		Config *config.Config
	}

	RouterResult struct {
		fx.Out

		Router *Router
	}
)

func NewRouter(p RouterParams) (RouterResult, error) {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		cors.AllowAll().Handler,
		middleware.RequestID,
		middleware.RealIP,
		middleware.Heartbeat("/ping"),
		middleware.Compress(5),
		middleware.Logger,
		middleware.Recoverer,
	)

	apiRouter := chi.NewRouter()
	r.Mount(p.Config.APIPrefix, apiRouter)

	return RouterResult{Router: &Router{r, apiRouter}}, nil
}

func (r *Router) Api() *chi.Mux {
	return r.api
}
