package db

import (
	"github.com/hamstag/fintech/core/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	Database struct {
		cfg *config.Config
		log *zap.Logger
	}

	Params struct {
		fx.In

		Config *config.Config
		Logger *zap.Logger `optional:"true"`
	}

	Result struct {
		fx.Out

		Database *Database
	}
)

func New(p Params) (Result, error) {
	log := p.Logger
	if log == nil {
		log = zap.NewNop()
	}

	return Result{
		Database: &Database{
			cfg: p.Config,
			log: log,
		},
	}, nil
}

func (d *Database) FindByID(id string) (string, error) {
	d.log.Info(d.cfg.APIPrefix)
	return "FindByID", nil
}
