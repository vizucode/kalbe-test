package domain

import "time"

type Attandance struct {
	Id              int       `json:"id"`
	EmployeeCode    string    `json:"employee_code"`
	LocationId      int       `json:"location_id" validate:"required"`
	AbsentIn        time.Time `json:"absent_in"`
	AbsentOut       time.Time `json:"absent_out"`
	PositionName    string    `json:"position_name"`
	EmployeeName    string    `json:"employee_name"`
	LocationName    string    `json:"location_name"`
	DepartementName string    `json:"departement_name"`
	UpdatedBy       string    `json:"updated_by"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
