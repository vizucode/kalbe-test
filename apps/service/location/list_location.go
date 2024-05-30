package location

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (s *location) GetAllLocation(ctx context.Context) (resp []domain.Location, err error) {

	locationList, err := s.db.GetAllLocation(ctx)
	if err != nil {
		log.Println(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	for _, v := range locationList {
		resp = append(resp, domain.Location{
			ID:           int(v.ID),
			LocationName: v.LocationName,
			CreatedBy:    v.CreatedBy,
			CreatedAt:    v.CreatedAt,
			UpdatedBy:    v.UpdatedBy,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return resp, nil
}
