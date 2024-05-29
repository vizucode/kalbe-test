package models

import (
	"time"

	"gorm.io/gorm"
)

type Attandance struct {
	gorm.Model
	EmployeeId int
	LocationId int
	AbsentIn   time.Time
	AbsentOut  time.Time
	UpdatedBy  string
	CreatedBy  string

	Employee Employee
	Location Location
}
