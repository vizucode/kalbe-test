package service

import (
	"context"

	"api.kalbe.crm/apps/domain"
)

type IAuth interface {
	SignIn(ctx context.Context, payload domain.AuthRequest) (token string, err error)
}

type IDepartement interface {
	GetAllDepartement(ctx context.Context) (resp []domain.Departement, err error)
	GetDepartement(ctx context.Context, id int) (resp domain.Departement, err error)
	CreateDepartement(ctx context.Context, payload domain.Departement) (err error)
	UpdateDepartement(ctx context.Context, payload domain.Departement) (err error)
	DeleteDepartement(ctx context.Context, id int) (err error)
}
