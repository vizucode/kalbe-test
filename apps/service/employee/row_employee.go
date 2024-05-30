package employee

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
)

func (s *employee) GetRowEmployee(ctx context.Context, employeeCode string) (resp domain.Employee, err error) {
	employee, err := s.db.GetEmployee(ctx, employeeCode)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return domain.Employee{
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
	}, nil
}
