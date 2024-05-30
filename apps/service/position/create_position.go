package position

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
)

func (uc *position) CreatePosition(ctx context.Context, payload domain.Position) (err error) {

	if err = uc.validator.StructCtx(ctx, &payload); err != nil {
		log.Println(err)
		return err
	}

	err = uc.db.CreatePosition(ctx, models.Position{
		DepartementId: payload.DepartementId,
		PositionName:  payload.PositionName,
		CreatedBy:     payload.CreatedBy,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
