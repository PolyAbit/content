package content

import (
	"context"
	"log/slog"

	"github.com/PolyAbit/content/internal/models"
)

type DirectionStorage interface {
	SaveDirection(ctx context.Context, code string, name string, exams string, description string) error
	GetDirections(ctx context.Context) ([]models.Direction, error)
	DeleteDirection(ctx context.Context, directionId int64) error
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
