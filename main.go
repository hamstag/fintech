package main

import (
	"net/http"

	"github.com/hamstag/fintech/core"
	"github.com/hamstag/fintech/wallet"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		core.Module,
		wallet.Module,
		fx.Invoke(func(*http.Server) {}),
	)
	app.Run()
}
