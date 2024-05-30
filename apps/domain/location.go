package domain

import "time"

type Location struct {
	ID           int       `json:"id"`
	LocationName string    `json:"location_name" validate:"required"`
	CreatedBy    string    `json:"created_by"`
	UpdatedBy    string    `json:"updated_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
