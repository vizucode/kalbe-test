package rest

import (
	"net/http"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/service"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	auth        service.IAuth
	departement service.IDepartement
	position    service.IPosition
	location    service.ILocation
	employee    service.IEmployee
}

func NewRest(
	auth service.IAuth,
	departement service.IDepartement,
	position service.IPosition,
	location service.ILocation,
	employee service.IEmployee,
) *rest {
	return &rest{
		auth:        auth,
		departement: departement,
		position:    position,
		location:    location,
		employee:    employee,
	}
}

func (rest *rest) ResponseJson(ctx *fiber.Ctx, data interface{}, message string, status bool) error {
	return ctx.Status(http.StatusOK).JSON(&domain.Response{
		Data:    data,
		Message: message,
		Status:  status,
	})
}

func (rest *rest) RegisterRoute(c *fiber.App) {
	v1 := c.Group("/api/v1")

	v1.Post("/auth/signin", rest.SignIn)

	/*
		Departement
	*/
	v1.Get("/departements", rest.GetListDepartement)
	v1.Get("/departement/:id", rest.GetRowDepartement)
	v1.Post("/departement", rest.CreateDepartement)
	v1.Put("/departement/:id", rest.UpdateDepartement)
	v1.Delete("/departement/:id", rest.DeleteDepartement)

	/*
		Position
	*/
	v1.Get("/positions", rest.GetListPosition)
	v1.Get("/position/:id", rest.GetRowPosition)
	v1.Post("/position", rest.CreatePosition)
	v1.Put("/position/:id", rest.UpdatePosition)
	v1.Delete("/position/:id", rest.DeletePosition)

	/*
		Location
	*/
	v1.Get("/locations", rest.GetListLocation)
	v1.Get("/location/:id", rest.GetRowLocation)
	v1.Post("/location", rest.CreateLocation)
	v1.Put("/location/:id", rest.UpdateLocation)
	v1.Delete("/location/:id", rest.DeleteLocation)

	/*
		Employee
	*/
	v1.Get("/employees", rest.GetAllEmployee)
	v1.Get("/employee/:employee_code", rest.GetEmployee)
	v1.Post("/employee", rest.CreateEmployee)
	v1.Put("/employee/:employee_code", rest.UpdateEmployee)
	v1.Delete("/employee/:employee_code", rest.DeleteEmployee)
}
