package main

import (
	"log/slog"

	"github.com/PolyAbit/content/internal/config"
	"github.com/PolyAbit/content/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env).With(slog.String("env", cfg.Env))

	log.Info("init config and logger")
}
