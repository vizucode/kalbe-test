package location

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
	"gorm.io/gorm"
)

func (s *location) UpdateLocation(ctx context.Context, payload domain.Location) (err error) {

	if err = s.validator.StructCtx(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	if err = s.db.UpdateLocation(ctx, []string{"location_name", "updated_By"}, models.Location{
		Model:        gorm.Model{ID: uint(payload.ID)},
		LocationName: payload.LocationName,
		UpdatedBy:    payload.UpdatedBy,
	}); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
