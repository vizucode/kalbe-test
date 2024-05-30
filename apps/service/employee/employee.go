package employee

import (
	"api.kalbe.crm/apps/repositories"
	"github.com/go-playground/validator/v10"
)

type employee struct {
	db        repositories.IDatabaseRepositories
	validator *validator.Validate
}

func NewEmployee(
	db repositories.IDatabaseRepositories,
	validator *validator.Validate,
) *employee {
	return &employee{
		db:        db,
		validator: validator,
	}
}
