package config

import (
	"gowebsite/pkg/db/postgres"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.PostgresConfig
}

func New(path string) *Config {
	cfg := Config{}
	if path == "" {
		path = "./configs/.env"
	}
	err := cleanenv.ReadConfig(path, &cfg)

	if err != nil {
		return nil
	}

	return &cfg
}
