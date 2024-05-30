package repositories

import (
	"context"

	"api.kalbe.crm/apps/models"
)

type IDatabaseRepositories interface {
	/*
		Departement
	*/
	GetAllDepartement(ctx context.Context) (resp []models.Departement, err error)
	GetDepartement(ctx context.Context, id int) (resp models.Departement, err error)
	CreateDepartement(ctx context.Context, payload models.Departement) (err error)
	UpdateDepartement(ctx context.Context, selectedField []string, payload models.Departement) (err error)
	DeleteDepartement(ctx context.Context, id int) (err error)

	/*
		Position
	*/
	GetAllPosition(ctx context.Context) (resp []models.Position, err error)
	GetPosition(ctx context.Context, id int) (resp models.Position, err error)
	CreatePosition(ctx context.Context, payload models.Position) (err error)
	UpdatePosition(ctx context.Context, selectedField []string, payload models.Position) (err error)
	DeletePosition(ctx context.Context, id int) (err error)

	/*
		Location
	*/
	// GetAllLocation(ctx context.Context) (resp []models.Location, err error)
	// GetLocation(ctx context.Context, id int) (resp models.Location, err error)
	// CreateLocation(ctx context.Context, payload models.Location) (err error)
	// UpdateLocation(ctx context.Context, selectedString []string, payload models.Location) (err error)
	// DeleteLocation(ctx context.Context, id int) (err error)

	/*
		Employee
	*/
	GetAllEmployee(ctx context.Context) (resp []models.Employee, err error)
	GetEmployee(ctx context.Context, employeeCode string) (resp models.Employee, err error)
	CreateEmployee(ctx context.Context, payload models.Employee) (err error)
	UpdateEmployee(ctx context.Context, selectedString []string, payload models.Location) (err error)
	DeleteEmployee(ctx context.Context, id int) (err error)

	/*
		Attandance
	*/
	// GetAllAttandance(ctx context.Context, startDate time.Time, endDate time.Time) (resp []models.Attandance, err error)
	// GetAttandance(ctx context.Context, id int) (resp models.Attandance, err error)
	// CreateAttandance(ctx context.Context, payload models.Attandance) (err error)
	// UpdateAttandance(ctx context.Context, selectedField []string, payload models.Attandance) (err error)
	// DeleteAttandance(ctx context.Context, id int) (err error)
}
