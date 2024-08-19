package http

import (
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	fx.Provide(
		NewRouter,
		NewServer,
	),
	fx.Invoke(func(r *Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, render.M{
				"message": "Hello",
			})
		})
	}),
)
