package departement

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
)

func (uc *departement) GetAllDepartement(ctx context.Context) (resp []domain.Departement, err error) {

	departementModelList, err := uc.db.GetAllDepartement(ctx)
	if err != nil {
		log.Println(err)
		return []domain.Departement{}, err
	}

	for _, departementModel := range departementModelList {
		departement := domain.Departement{
			Id:              int(departementModel.Model.ID),
			DepartementName: departementModel.DepartementName,
			CreatedBy:       departementModel.CreatedBy,
			UpdatedBy:       departementModel.UpdatedBy,
			CreatedAt:       departementModel.Model.CreatedAt,
			UpdatedAt:       departementModel.Model.UpdatedAt,
		}
		resp = append(resp, departement)
	}

	return resp, nil
}
