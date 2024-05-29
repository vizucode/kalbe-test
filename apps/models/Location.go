package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	LocationName string
	CreatedBy    string
	UpdatedBy    string
}
