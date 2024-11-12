package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	UserName string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	DbName   string `env:"POSTGRES_DB" envDefault:"postgres"`
}

type DB struct {
	*sqlx.DB
}

func New(ctx context.Context, config PostgresConfig) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.UserName, config.Password, config.DbName, config.Host, config.Port)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if _, err := db.Conn(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return &DB{db}, nil
}
