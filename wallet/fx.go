package wallet

import (
	"go.uber.org/fx"
)

var Module = fx.Module("wallet",
	fx.Provide(New),
	fx.Invoke(HTTPHandler),
)
