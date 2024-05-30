package attandance

import (
	"api.kalbe.crm/apps/repositories"
	"github.com/go-playground/validator/v10"
)

type attandance struct {
	db        repositories.IDatabaseRepositories
	validator *validator.Validate
}

func NewAttandance(
	db repositories.IDatabaseRepositories,
	validator *validator.Validate,
) *attandance {
	return &attandance{
		db:        db,
		validator: validator,
	}
}
