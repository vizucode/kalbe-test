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
}

func NewRest(
	auth service.IAuth,
	departement service.IDepartement,
	position service.IPosition,
) *rest {
	return &rest{
		auth:        auth,
		departement: departement,
		position:    position,
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
}
