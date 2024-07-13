package content

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/PolyAbit/content/internal/models"
)

func (c *Content) CreateDirection(ctx context.Context, code string, name string, exams string, description string) (models.Direction, error) {
	const op = "services.content.CreateDirection"

	log := c.log.With(slog.String("op", op))


	newDirection, err := c.directionProvider.SaveDirection(ctx, code, name, exams, description)
	if errors.Is(err, models.ErrDirectionExists) {
		return models.Direction{}, fmt.Errorf("%s: %w", op, ErrCodeAlreadyUsed)
	}
	if err != nil {
		log.Error("failed to save direction", err)

		return models.Direction{}, fmt.Errorf("%s: %w", op, err)
	}

	return newDirection, nil
}
