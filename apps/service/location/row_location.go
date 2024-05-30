package location

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (s *location) GetLocation(ctx context.Context, id int) (resp domain.Location, err error) {

	location, err := s.db.GetLocation(ctx, id)
	if err != nil {
		log.Println(err)
		return domain.Location{}, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	resp = domain.Location{
		ID:           int(location.ID),
		LocationName: location.LocationName,
		CreatedBy:    location.CreatedBy,
		CreatedAt:    location.CreatedAt,
		UpdatedBy:    location.UpdatedBy,
		UpdatedAt:    location.UpdatedAt,
	}

	return resp, err
}
