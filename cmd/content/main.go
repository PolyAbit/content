package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/PolyAbit/content/internal/app"
	"github.com/PolyAbit/content/internal/config"
	"github.com/PolyAbit/content/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env).With(slog.String("env", cfg.Env))

	log.Info("init config and logger")

	application := app.New(log, cfg)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("code", sign.String()))

	application.GRPCServer.Stop()

	log.Info("application stopped")
}
