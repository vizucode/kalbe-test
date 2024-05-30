package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeCode  string
	EmployeeName  string
	Password      string
	DepartementId int
	PositionId    int
	Superior      bool
	CreatedBy     string
	UpdatedBy     string

	Departement Departement
	Position    Position
}
