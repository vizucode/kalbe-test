package position

import (
	"api.kalbe.crm/apps/repositories"
	"github.com/go-playground/validator/v10"
)

type position struct {
	db        repositories.IDatabaseRepositories
	validator *validator.Validate
}

func NewPosition(
	db repositories.IDatabaseRepositories,
	validator *validator.Validate,
) *position {
	return &position{
		db:        db,
		validator: validator,
	}
}
