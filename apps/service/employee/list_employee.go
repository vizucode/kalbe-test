package employee

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
)

func (s *employee) GetAllEmployee(ctx context.Context) (resp []domain.Employee, err error) {

	employeeList, err := s.db.GetAllEmployee(ctx)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	for _, employee := range employeeList {
		resp = append(resp, domain.Employee{
			Id:            int(employee.ID),
			EmployeeCode:  employee.EmployeeCode,
			EmployeeName:  employee.EmployeeName,
			DepartementId: int(employee.DepartementId),
			PositionId:    int(employee.PositionId),
			Superior:      employee.Superior,
			CreatedBy:     employee.CreatedBy,
			UpdatedBy:     employee.UpdatedBy,
			CreatedAt:     employee.CreatedAt,
			UpdatedAt:     employee.UpdatedAt,
		})
	}

	return resp, nil
}
