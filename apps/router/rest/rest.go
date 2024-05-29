package rest

import (
	"net/http"

	"api.kalbe.crm/apps/domain"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
}

func NewRest() *rest {
	return &rest{}
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
