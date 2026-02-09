package grpc

import (
	"fmt"
	"log/slog"
	"net"

	grpcAuth "github.com/kuaukutsu/auth/sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	auth grpcAuth.Auth,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	grpcAuth.Register(gRPCServer, auth)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (app *App) Run() error {
	const op = "app.grpc.Run"
	log := app.log.With(
		slog.String("op", op),
		slog.Int("port", app.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("add", l.Addr().String()))

	if err := app.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *App) Stop() {
	const op = "app.grpc.Stop"

	app.log.With(
		slog.String("op", op),
		slog.Int("port", app.port),
	).Info("stopping gRPC server")

	app.gRPCServer.GracefulStop()
}
