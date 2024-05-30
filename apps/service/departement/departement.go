package departement

import (
	"api.kalbe.crm/apps/repositories"
	"github.com/go-playground/validator/v10"
)

type departement struct {
	db        repositories.IDatabaseRepositories
	validator *validator.Validate
}

func NewDepartement(
	db repositories.IDatabaseRepositories,
	validator *validator.Validate,
) *departement {
	return &departement{
		db:        db,
		validator: validator,
	}
}
