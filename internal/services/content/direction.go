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
