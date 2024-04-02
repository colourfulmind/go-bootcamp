package main

import (
	"articles/internal/app"
	"articles/internal/config"
	"articles/internal/logo"
	"articles/pkg/logger/logsetup"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logo.CreateFile()

	cfg := config.MustLoad()

	log := logsetup.SetupLogger(cfg.Env)
	log.Info("starting application", slog.Any("config", cfg))

	application, err := app.New(log, cfg)
	if err != nil {
		return
	}
	go application.Server.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sgl := <-stop
	log.Info("stopping application", slog.String("signal", sgl.String()))
	application.Server.Stop()
	log.Info("application stopped")
}
