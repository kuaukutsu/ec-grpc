package app

import (
	"log/slog"
	"time"

	"github.com/kuaukutsu/auth/sso/internal/app/grpc"
	"github.com/kuaukutsu/auth/sso/internal/services/auth"
	"github.com/kuaukutsu/auth/sso/internal/storage/memory"
)

type App struct {
	GRPCServer *grpc.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTtl time.Duration,
) *App {
	storage := memory.New()
	authService := auth.New(log, storage, storage, storage, tokenTtl)

	grpc := grpc.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpc,
	}
}
