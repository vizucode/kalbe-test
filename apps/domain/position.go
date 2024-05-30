package domain

import "time"

type Position struct {
	Id            int       `json:"id"`
	DepartementId int       `json:"departement_id" validate:"required"`
	PositionName  string    `json:"position_name" validate:"required"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedBy     string    `json:"updated_by"`
	UpdatedAt     time.Time `json:"updated_at"`
}
