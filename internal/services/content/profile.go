package content

import (
	"context"

	"github.com/PolyAbit/content/internal/lib/logger/sl"
	"github.com/PolyAbit/content/internal/models"
)

func (c *Content) GetProfile(ctx context.Context, userId int64) (models.Profile, error) {
	profile, err := c.profileProvider.GetProfile(ctx, userId)

	if err != nil {
		c.log.Error("failed to get profile", sl.Err(err))

		return models.Profile{}, err
	}

	return profile, nil
}

func (c *Content) UpdateProfile(ctx context.Context, userId int64, fistName string, middleName string, lastName string) (models.Profile, error) {
	profile, err := c.profileProvider.UpdateProfile(ctx, userId, fistName, middleName, lastName)

	if err != nil {
		c.log.Error("failed to update profile", sl.Err(err))

		return models.Profile{}, err
	}

	return profile, nil
}
