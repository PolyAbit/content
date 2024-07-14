package content

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/PolyAbit/content/internal/lib/logger/sl"
	"github.com/PolyAbit/content/internal/models"
)

func (c *Content) CreateDirection(ctx context.Context, code string, name string, exams string, description string) error {
	const op = "services.content.CreateDirection"

	log := c.log.With(slog.String("op", op))

	err := c.directionProvider.SaveDirection(ctx, code, name, exams, description)

	if errors.Is(err, models.ErrDirectionExists) {
		return fmt.Errorf("%s: %w", op, ErrCodeAlreadyUsed)
	}
	if err != nil {
		log.Error("failed to save direction", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (c *Content) GetDirections(ctx context.Context) ([]models.Direction, error) {
	directions, err := c.directionProvider.GetDirections(ctx)

	if err != nil {
		c.log.Error("failed to save direction", sl.Err(err))

		return []models.Direction{}, err
	}

	return directions, nil
}

func (c *Content) DeleteDirection(ctx context.Context, directionId int64) error {
	err := c.directionProvider.DeleteDirection(ctx, directionId)

	if err != nil {
		c.log.Error("failed to delete direction", sl.Err(err))

		return err
	}

	return nil
}
