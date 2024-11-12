package config

import (
	"gowebsite/pkg/db/postgres"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.PostgresConfig
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("./configs/.env", &cfg)

	if err != nil {
		return nil
	}

	return &cfg
}
