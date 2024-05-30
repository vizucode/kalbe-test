package domain

import "time"

type Employee struct {
	Id            int       `json:"id"`
	EmployeeCode  string    `json:"employee_code"`
	EmployeeName  string    `json:"employee_name" validate:"required"`
	Password      string    `json:"password,omitempty" validate:"required"`
	DepartementId int       `json:"departement_id" valdidate:"required"`
	PositionId    int       `json:"position_id" validate:"required"`
	Superior      bool      `json:"superior" validate:"required"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
