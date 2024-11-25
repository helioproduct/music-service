package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

type SlogLogger struct {
	logger *slog.Logger
	env    string
}

func NewSlogLogger(env string) *SlogLogger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, nil))

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, nil),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, nil),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, nil),
		)
	}

	return &SlogLogger{
		logger: log,
		env:    env,
	}
}

func New(level string) *SlogLogger {
	var lvl slog.Level
	switch level {
	case "debug":
		lvl = slog.LevelDebug
	case "info":
		lvl = slog.LevelInfo
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: lvl})
	return &SlogLogger{
		logger: slog.New(handler),
	}
}

func (l *SlogLogger) Debug(message string, args ...interface{}) {
	l.log(slog.LevelDebug, message, args...)
}

func (l *SlogLogger) Info(message string, args ...interface{}) {
	l.log(slog.LevelInfo, message, args...)
}

func (l *SlogLogger) Warn(message string, args ...interface{}) {
	l.log(slog.LevelWarn, message, args...)
}

func (l *SlogLogger) Error(message string, args ...interface{}) {
	l.log(slog.LevelError, message, args...)
}

func (l *SlogLogger) Fatal(message string, args ...interface{}) {
	l.log(slog.LevelError, message, args...)
	os.Exit(1)
}

func (l *SlogLogger) log(level slog.Level, message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Log(context.TODO(), level, message)
	} else {
		l.logger.Log(context.TODO(), level, message, "args", args)
	}
}
