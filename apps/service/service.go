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

type IPosition interface {
	GetAllPosition(ctx context.Context) (resp []domain.Position, err error)
	GetPosition(ctx context.Context, id int) (resp domain.Position, err error)
	CreatePosition(ctx context.Context, payload domain.Position) (err error)
	UpdatePosition(ctx context.Context, payload domain.Position) (err error)
	DeletePosition(ctx context.Context, id int) (err error)
}
