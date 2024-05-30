package domain

import "time"

type Departement struct {
	Id              int       `json:"id"`
	DepartementName string    `json:"departement_name" validate:"required"`
	CreatedBy       string    `json:"created_by"`
	UpdatedBy       string    `json:"updated_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
