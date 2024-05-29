package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	DepartementId int
	PositionName  string
	CreatedBy     string
	UpdatedBy     string

	Departement Departement
}
