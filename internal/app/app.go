package app

import (
	"log/slog"

	grpcapp "github.com/PolyAbit/content/internal/app/grpc"
	"github.com/PolyAbit/content/internal/services/content"
	"github.com/PolyAbit/content/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	httpPort int,
	storagePath string,
) *App {
	storage, err := sqlite.New(storagePath)

	if err != nil {
		panic(err)
	}

	contentService := content.New(log, storage)

	grpcApp := grpcapp.New(log, contentService, grpcPort, httpPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
