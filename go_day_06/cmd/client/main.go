package main

import (
	grpcclient "articles/internal/clients/blog/grpc"
	"articles/internal/config"
	"articles/pkg/logger/logsetup"
	"articles/pkg/logger/sl"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	const op = "cmd/client/New"

	cfg := config.MustLoad()

	log := logsetup.SetupLogger(cfg.Env)
	log.Info("starting client", slog.Any("config", cfg.Clients.GRPCClient))

	cc, err := grpcclient.NewConnection(
		context.Background(),
		log,
		cfg.GRPC.Host+":"+strconv.Itoa(cfg.GRPC.Port),
		cfg.Clients.GRPCClient.RetriesCount,
		cfg.Clients.GRPCClient.Timeout,
	)
	if err != nil {
		log.Error("failed to connect to server", op, sl.Err(err))
		os.Exit(1)
	}
	defer cc.Close()

	go grpcclient.New(cc, log, cfg).Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sgl := <-stop
	log.Info("stopping client", slog.String("signal", sgl.String()))
	log.Info("client stopped")
}
