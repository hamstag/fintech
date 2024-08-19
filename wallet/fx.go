package wallet

import (
	"go.uber.org/fx"
)

var Module = fx.Module("wallet",
	fx.Provide(
		fx.Annotate(
			NewWalletService,
			fx.As(new(WalletService)),
		),
	),
	fx.Invoke(HTTPHandler),
)
