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
}

func NewRest(
	auth service.IAuth,
	departement service.IDepartement,
) *rest {
	return &rest{
		auth:        auth,
		departement: departement,
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
}
