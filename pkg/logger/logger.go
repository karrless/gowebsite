package logger

import (
	"context"

	"go.uber.org/zap"
)

type KeyString string

const (
	LoggerKey KeyString = "logger"
	RequestID KeyString = "request_id"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Fatal(ctx context.Context, msg string, fields ...zap.Field)
}

type logger struct {
	logger *zap.Logger
}

func (l *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}
	l.logger.Info(msg, fields...)
}

func (l *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}
	l.logger.Error(msg, fields...)
}

func (l *logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}
	l.logger.Debug(msg, fields...)
}

func (l *logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}
	l.logger.Warn(msg, fields...)
}

func (l *logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String(string(RequestID), ctx.Value(RequestID).(string)))
	}
	l.logger.Fatal(msg, fields...)
}

func New() Logger {
	zapLogger, _ := zap.NewDevelopment()
	
	defer zapLogger.Sync()

	return &logger{logger: zapLogger}
}

func GetLoggerFromCtx(ctx context.Context) Logger {
	return ctx.Value(LoggerKey).(Logger)
}
