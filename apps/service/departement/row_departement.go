package departement

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
)

func (uc *departement) GetDepartement(ctx context.Context, id int) (resp domain.Departement, err error) {

	departementModel, err := uc.db.GetDepartement(ctx, id)
	if err != nil {
		log.Println(err)
		return domain.Departement{}, err
	}

	resp = domain.Departement{
		Id:              int(departementModel.Model.ID),
		DepartementName: departementModel.DepartementName,
		CreatedBy:       departementModel.CreatedBy,
		UpdatedBy:       departementModel.UpdatedBy,
		CreatedAt:       departementModel.Model.CreatedAt,
		UpdatedAt:       departementModel.Model.UpdatedAt,
	}

	return resp, nil
}
