package departement

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
)

func (uc *departement) CreateDepartement(ctx context.Context, payload domain.Departement) (err error) {
	if err = uc.validator.StructCtx(ctx, &payload); err != nil {
		log.Println(err)
		return err
	}

	err = uc.db.CreateDepartement(ctx, models.Departement{
		DepartementName: payload.DepartementName,
		CreatedBy:       payload.CreatedBy,
		UpdatedBy:       payload.UpdatedBy,
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
