package location

import (
	"api.kalbe.crm/apps/repositories"
	"github.com/go-playground/validator/v10"
)

type location struct {
	db        repositories.IDatabaseRepositories
	validator *validator.Validate
}

func NewLocation(
	db repositories.IDatabaseRepositories,
	validator *validator.Validate,
) *location {
	return &location{
		db:        db,
		validator: validator,
	}
}
