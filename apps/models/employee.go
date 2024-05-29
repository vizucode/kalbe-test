package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeCode  string
	Password      string
	DepartementId int
	PositionId    int
	Superior      int
	CreatedBy     string
	UpdatedBy     string

	Departement Departement
	Position    Position
}
