package main

import (
	"context"
	"gowebsite/internal/config"
	"gowebsite/pkg/db/postgres"
	"gowebsite/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New()
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	mainLogger.Info(ctx, "Application is starting...")

	cfg := config.New("")
	if cfg == nil {
		mainLogger.Fatal(ctx, "failed to load config")
	}
	mainLogger.Debug(ctx, "Config loaded", zap.Any("config", cfg))
	db, err := postgres.New(ctx, cfg.PostgresConfig)
	if err != nil {
		mainLogger.Fatal(ctx, "failed to connect to database", zap.Error(err))
	}
	mainLogger.Debug(ctx, "Database connected")
	graceChannel := make(chan os.Signal, 1)
	signal.Notify(graceChannel, syscall.SIGINT, syscall.SIGTERM)
	<-graceChannel
	db.Close()
	mainLogger.Debug(ctx, "Database connection closed")
	mainLogger.Info(ctx, "Graceful shutdown!")
}
