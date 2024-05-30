package position

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
	"gorm.io/gorm"
)

func (uc *position) UpdatePosition(ctx context.Context, payload domain.Position) (err error) {

	if err = uc.validator.StructCtx(ctx, &payload); err != nil {
		log.Println(err)
		return err
	}

	err = uc.db.UpdatePosition(ctx, []string{"position_name", "departement_id", "updated_by", "updated_at"}, models.Position{
		Model: gorm.Model{
			ID: uint(payload.Id),
		},
		PositionName:  payload.PositionName,
		DepartementId: payload.DepartementId,
		UpdatedBy:     payload.UpdatedBy,
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
