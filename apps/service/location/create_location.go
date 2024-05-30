package location

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
)

func (s *location) CreateLocation(ctx context.Context, payload domain.Location) (err error) {
	if err = s.validator.StructCtx(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	if err = s.db.CreateLocation(ctx, models.Location{
		LocationName: payload.LocationName,
		CreatedBy:    payload.CreatedBy,
		UpdatedBy:    payload.UpdatedBy,
	}); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
