package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kuaukutsu/auth/sso/internal/app"
	"github.com/kuaukutsu/auth/sso/internal/config"
)

const (
	envLocal = "local"
	envProd  = "production"
)

func main() {
	cfg := config.NewConfig()
	log := setupLogger(cfg.Env)

	log.Debug("application", slog.Any("cfg", cfg))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go func() {
		if err := application.GRPCServer.Run(); err != nil {
			log.Error("application.grpc", slog.Attr{
				Key:   "error",
				Value: slog.StringValue(err.Error()),
			})
		}
	}()

	interrupt := make(chan os.Signal, 3)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	interruptSignal := <-interrupt

	application.GRPCServer.Stop()

	log.Info("application.grpc: stopped", slog.String("signal", interruptSignal.String()))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
