package attandance

import (
	"context"
	"log"
	"time"

	"api.kalbe.crm/apps/domain"
)

func (s *attandance) GetAllAttandance(ctx context.Context, startDate time.Time, endDate time.Time) (resp []domain.Attandance, err error) {

	attandanceList, err := s.db.GetAllAttandance(ctx, startDate, endDate)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	for _, attandance := range attandanceList {
		resp = append(resp, domain.Attandance{
			Id:              int(attandance.ID),
			LocationId:      int(attandance.LocationId),
			EmployeeCode:    attandance.Employee.EmployeeCode,
			EmployeeName:    attandance.Employee.EmployeeName,
			PositionName:    attandance.Employee.Position.PositionName,
			DepartementName: attandance.Employee.Departement.DepartementName,
			LocationName:    attandance.Location.LocationName,
			CreatedBy:       attandance.CreatedBy,
			UpdatedBy:       attandance.UpdatedBy,
			AbsentIn:        attandance.AbsentIn,
			AbsentOut:       attandance.AbsentOut,
			CreatedAt:       attandance.CreatedAt,
			UpdatedAt:       attandance.UpdatedAt,
		})
	}

	return resp, nil
}
