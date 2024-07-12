package app

import (
	"log/slog"

	grpcapp "github.com/PolyAbit/content/internal/app/grpc"
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
	// storage, err := sqlite.New(storagePath)

	// if err != nil {
	// 	panic(err)
	// }

	// authService := auth.New(log, storage, tokenTTL, tokenSecret)

	grpcApp := grpcapp.New(log, nil, grpcPort, httpPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
