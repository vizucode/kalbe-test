package models

import "gorm.io/gorm"

type Departement struct {
	gorm.Model
	DepartementName string
	CreatedBy       string
	UpdatedBy       string

	Positions []Position
}
