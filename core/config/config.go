package config

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"
)

type Config struct {
	Host        string `env:"HOST" envDefault:"0.0.0.0"`
	Port        int    `env:"PORT" envDefault:"8080"`
	Address     string `env:"ADDRESS,expand" envDefault:"$HOST:${PORT}"`
	APIPrefix   string `env:"API_PREFIX" envDefault:"/api"`
	RedisURL    string `env:"REDIS_URL,notEmpty"`
	DatabaseURL string `env:"DATABASE_URL,notEmpty"`
}

type Result struct {
	fx.Out

	Config *Config
}

func New() (Result, error) {
	cfg, err := env.ParseAsWithOptions[Config](env.Options{
		RequiredIfNoDef: true,
	})
	return Result{Config: &cfg}, err
}
