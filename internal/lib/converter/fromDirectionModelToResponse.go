package converter

import (
	"github.com/PolyAbit/content/internal/models"
	contentv1 "github.com/PolyAbit/protos/gen/go/content"
)

func ConvertDirection(direction models.Direction) *contentv1.Direction {
	return &contentv1.Direction{
		Id:          direction.Id,
		Code:        direction.Code,
		Name:        direction.Name,
		Exams:       direction.Exams,
		Description: direction.Description,
	}
}

func ConvertProfile(profile models.Profile) *contentv1.Profile {
	return &contentv1.Profile{
		UserId:     profile.UserId,
		FirstName:  profile.FirstName,
		MiddleName: profile.MiddleName,
		LastName:   profile.LastName,
	}
}
