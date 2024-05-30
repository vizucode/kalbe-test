package rest

import (
	"net/http"

	"api.kalbe.crm/apps/domain"
	middleware "api.kalbe.crm/apps/middlewares"
	"api.kalbe.crm/apps/service"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	auth        service.IAuth
	departement service.IDepartement
	position    service.IPosition
	location    service.ILocation
	employee    service.IEmployee
	attandace   service.IAttandance
}

func NewRest(
	auth service.IAuth,
	departement service.IDepartement,
	position service.IPosition,
	location service.ILocation,
	employee service.IEmployee,
	attandance service.IAttandance,
) *rest {
	return &rest{
		auth:        auth,
		departement: departement,
		position:    position,
		location:    location,
		employee:    employee,
		attandace:   attandance,
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
	v1.Get("/departements", middleware.VerifyToken, rest.GetListDepartement)
	v1.Get("/departement/:id", middleware.VerifyToken, rest.GetRowDepartement)
	v1.Post("/departement", middleware.VerifyToken, rest.CreateDepartement)
	v1.Put("/departement/:id", middleware.VerifyToken, rest.UpdateDepartement)
	v1.Delete("/departement/:id", middleware.VerifyToken, rest.DeleteDepartement)

	/*
		Position
	*/
	v1.Get("/positions", middleware.VerifyToken, rest.GetListPosition)
	v1.Get("/position/:id", middleware.VerifyToken, rest.GetRowPosition)
	v1.Post("/position", middleware.VerifyToken, rest.CreatePosition)
	v1.Put("/position/:id", middleware.VerifyToken, rest.UpdatePosition)
	v1.Delete("/position/:id", middleware.VerifyToken, rest.DeletePosition)

	/*
		Location
	*/
	v1.Get("/locations", middleware.VerifyToken, rest.GetListLocation)
	v1.Get("/location/:id", middleware.VerifyToken, rest.GetRowLocation)
	v1.Post("/location", middleware.VerifyToken, rest.CreateLocation)
	v1.Put("/location/:id", middleware.VerifyToken, rest.UpdateLocation)
	v1.Delete("/location/:id", middleware.VerifyToken, rest.DeleteLocation)

	/*
		Employee
	*/
	v1.Get("/employees", middleware.VerifyToken, rest.GetAllEmployee)
	v1.Get("/employee/:employee_code", middleware.VerifyToken, rest.GetEmployee)
	v1.Post("/employee", middleware.VerifyToken, rest.CreateEmployee)
	v1.Put("/employee/:employee_code", middleware.VerifyToken, rest.UpdateEmployee)
	v1.Delete("/employee/:employee_code", middleware.VerifyToken, rest.DeleteEmployee)

	/*
		Attandance
	*/
	v1.Get("/attandances", middleware.VerifyToken, rest.GetAllAttandance)
	v1.Get("/attandance/:id", middleware.VerifyToken, rest.GetRowAttandance)
	v1.Post("/attandance/checkin", middleware.VerifyToken, rest.CheckIn)
	v1.Post("/attandance/checkout/:id", middleware.VerifyToken, rest.CheckOut)
}
