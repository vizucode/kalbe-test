package departement

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
	"gorm.io/gorm"
)

func (uc *departement) UpdateDepartement(ctx context.Context, payload domain.Departement) (err error) {

	if err = uc.validator.StructCtx(ctx, &payload); err != nil {
		log.Println(err)
		return err
	}

	err = uc.db.UpdateDepartement(ctx, []string{"departement_name", "updated_by", "updated_at"}, models.Departement{
		Model: gorm.Model{
			ID: uint(payload.Id),
		},
		DepartementName: payload.DepartementName,
		UpdatedBy:       payload.UpdatedBy,
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
