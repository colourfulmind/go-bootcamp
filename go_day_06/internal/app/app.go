package app

import (
	grpcserver "articles/internal/app/grpc"
	"articles/internal/config"
	"articles/internal/services/articles"
	"articles/internal/services/auth"
	gormdb "articles/internal/storage/gorm"
	"articles/pkg/logger/sl"
	"fmt"
	"log/slog"
)

type App struct {
	Server *grpcserver.App
}

func New(log *slog.Logger, cfg *config.Config) (*App, error) {
	const op = "internal/app/New"
	storage, err := gormdb.New(cfg.Postgres)
	if err != nil {
		log.Error("error occurred while connecting to database", op, sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	authService := auth.New(log, storage, storage, cfg.TokenTTL)
	articlesService := articles.New(log, storage, storage, cfg.TokenTTL)

	return &App{
		Server: grpcserver.New(log, authService, articlesService, cfg.GRPC.Port),
	}, nil
}
