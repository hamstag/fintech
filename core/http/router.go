package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/hamstag/fintech/core/config"
)

type Router struct {
	*chi.Mux

	api *chi.Mux
}

func NewRouter(cfg *config.Config) (*Router, error) {
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
	r.Mount(cfg.APIPrefix, apiRouter)

	return &Router{r, apiRouter}, nil
}

func (r *Router) Api() *chi.Mux {
	return r.api
}
