package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("./configs/.env", &cfg)

	if err != nil {
		return nil
	}

	return &cfg
}
