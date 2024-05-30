package service

import (
	"context"
	"time"

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

type ILocation interface {
	GetAllLocation(ctx context.Context) (resp []domain.Location, err error)
	GetLocation(ctx context.Context, id int) (resp domain.Location, err error)
	CreateLocation(ctx context.Context, payload domain.Location) (err error)
	UpdateLocation(ctx context.Context, payload domain.Location) (err error)
	DeleteLocation(ctx context.Context, id int) (err error)
}

type IEmployee interface {
	GetAllEmployee(ctx context.Context) (resp []domain.Employee, err error)
	GetRowEmployee(ctx context.Context, employeeCode string) (resp domain.Employee, err error)
	CreateEmployee(ctx context.Context, payload domain.Employee) (err error)
	UpdateEmployee(ctx context.Context, payload domain.Employee) (err error)
	DeleteEmployee(ctx context.Context, employeeCode string) (err error)
}

type IAttandance interface {
	CheckIn(ctx context.Context, payload domain.Attandance) (err error)
	CheckOut(ctx context.Context, payload domain.Attandance) (err error)
	GetAllAttandance(ctx context.Context, startDate time.Time, endDate time.Time) (resp []domain.Attandance, err error)
	GetRowAttandance(ctx context.Context, id int) (resp domain.Attandance, err error)
}
