package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

// Config represents the application configuration
type Config struct {
	DatabasePassword     string `env:"DB_PASSWORD,required"`
	GoogleClientSecret   string `env:"GOOGLE_CLIENT_SECRET,default="`
	FacebookClientSecret string `env:"FACEBOOK_CLIENT_SECRET,default="`
	JWTSecret            string `env:"JWT_SECRET,required"`
}

// Load loads configuration from environment variables using envconfig
func Load() (*Config, error) {
	ctx := context.Background()
	var cfg Config

	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
