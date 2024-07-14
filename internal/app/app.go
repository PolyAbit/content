package app

import (
	"log/slog"

	grpcapp "github.com/PolyAbit/content/internal/app/grpc"
	"github.com/PolyAbit/content/internal/config"
	"github.com/PolyAbit/content/internal/services/content"
	"github.com/PolyAbit/content/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	storage, err := sqlite.New(cfg.StoragePath)

	if err != nil {
		panic(err)
	}

	contentService := content.New(log, storage)

	grpcApp := grpcapp.New(log, contentService, cfg.GRPC.Port, cfg.GRPC.GatewayPort, cfg.JwtSecret)

	return &App{
		GRPCServer: grpcApp,
	}
}
