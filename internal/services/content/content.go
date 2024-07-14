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

type ProfileStorage interface {
	GetProfile(ctx context.Context, useId int64) (models.Profile, error)
	UpdateProfile(ctx context.Context, useId int64, fistName string, middleName string, lastName string) (models.Profile, error)
}

type Content struct {
	log               *slog.Logger
	directionProvider DirectionStorage
	profileProvider   ProfileStorage
}

func New(log *slog.Logger, directionProvider DirectionStorage, profileProvider ProfileStorage) *Content {
	return &Content{
		log:               log,
		directionProvider: directionProvider,
		profileProvider:   profileProvider,
	}
}
