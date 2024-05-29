package rest

import (
	"net/http"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/service"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	auth service.IAuth
}

func NewRest(
	auth service.IAuth,
) *rest {
	return &rest{
		auth: auth,
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
	// v1 := c.Group("/api/v1")
}
