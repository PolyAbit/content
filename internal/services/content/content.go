package content

import (
	"context"
	"log/slog"
)

type DirectionStorage interface {
	SaveDirection(ctx context.Context, code string, name string, exams string, description string) (error)
}

type Content struct {
	log               *slog.Logger
	directionProvider DirectionStorage
}

func New(log *slog.Logger, directionProvider DirectionStorage) *Content {
	return &Content{
		log:               log,
		directionProvider: directionProvider,
	}
}
