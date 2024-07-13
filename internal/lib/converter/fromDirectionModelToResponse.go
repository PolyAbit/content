package converter

import (
	"github.com/PolyAbit/content/internal/models"
	contentv1 "github.com/PolyAbit/protos/gen/go/content"
)

func FromDirectionModelToResponse(direction models.Direction) *contentv1.Direction {
	return &contentv1.Direction{
		Id:          direction.Id,
		Code:        direction.Code,
		Name:        direction.Name,
		Exams:       direction.Exams,
		Description: direction.Description,
	}
}
