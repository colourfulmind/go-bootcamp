package grpc

import (
	"articles/internal/grpc/articles"
	"articles/internal/grpc/auth"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, authService auth.Auth, articlesService articles.Articles, port int) *App {
	GRPCSever := grpc.NewServer()
	auth.Register(GRPCSever, authService)
	articles.Register(GRPCSever, articlesService)
	return &App{
		log:        log,
		GRPCServer: GRPCSever,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "internal/app/grpc/Run"
	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err = a.GRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "internal/app/grpc/Stop"

	a.log.With(slog.String("op", op)).Info(
		"stopping gRPC server", slog.Int("port", a.port),
	)

	a.GRPCServer.GracefulStop()
}
