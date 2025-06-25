package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

type Config struct {
	SdkPath string `env:"SDK_PATH"`
}

func InitiateConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(homeDir, "android-sdk-server.env")

	if err := godotenv.Load(envPath); err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
