package main

import (
	"context"
	"gowebsite/pkg/logger"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New()
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	mainLogger.Debug(ctx, "hello world")
}
