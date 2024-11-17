package config

import (
	"gowebsite/pkg/db/postgres"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.PostgresConfig
	RESTServerPort string `env:"REST_SERVER_PORT" envDefault:"8080"`
	RESTServerHost string `env:"REST_SERVER_HOST" envDefault:"localhost"`
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
