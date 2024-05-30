package attandance

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
)

func (s *attandance) GetRowAttandance(ctx context.Context, id int) (resp domain.Attandance, err error) {

	attandance, err := s.db.GetAttandance(ctx, id)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	resp = domain.Attandance{
		Id:              int(attandance.ID),
		LocationId:      int(attandance.LocationId),
		EmployeeCode:    attandance.Employee.EmployeeCode,
		EmployeeName:    attandance.Employee.EmployeeName,
		PositionName:    attandance.Employee.Position.PositionName,
		DepartementName: attandance.Employee.Departement.DepartementName,
		LocationName:    attandance.Location.LocationName,
		CreatedBy:       attandance.CreatedBy,
		AbsentIn:        attandance.AbsentIn,
		AbsentOut:       attandance.AbsentOut,
		CreatedAt:       attandance.CreatedAt,
		UpdatedAt:       attandance.UpdatedAt,
	}

	return resp, nil
}
