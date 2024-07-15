package app

import (
	"context"
	"log/slog"

	grpcapp "github.com/PolyAbit/content/internal/app/grpc"
	grpcauth "github.com/PolyAbit/content/internal/clients/auth/grpc"
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

	contentService := content.New(log, storage, storage)

	authClient, err := grpcauth.New(context.Background(), log, cfg.Clients.Auth.Address, cfg.Clients.Auth.Timeout, int(cfg.Clients.Auth.RetriesCount))

	if err != nil {
		panic(err)
	}

	grpcApp := grpcapp.New(log, contentService, authClient, cfg.GRPC.Port, cfg.GRPC.GatewayPort, cfg.JwtSecret)

	return &App{
		GRPCServer: grpcApp,
	}
}
