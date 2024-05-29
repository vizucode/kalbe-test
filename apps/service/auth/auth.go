package auth

import (
	"api.kalbe.crm/apps/repositories"
	"github.com/go-playground/validator/v10"
)

type auth struct {
	db         repositories.IDatabaseRepositories
	validation *validator.Validate
}

func NewAuth(
	db repositories.IDatabaseRepositories,
	validate *validator.Validate,
) *auth {
	return &auth{
		db:         db,
		validation: validate,
	}
}
