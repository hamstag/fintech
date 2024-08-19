package config

import (
	"os"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("RedisURL", "")
	os.Setenv("DATABASE_URL", "")

	cfg, err := env.ParseAsWithOptions[Config](env.Options{
		RequiredIfNoDef: true,
	})

	assert.Error(t, err)
	assert.Equal(t, "127.0.0.1", cfg.Host)
}
