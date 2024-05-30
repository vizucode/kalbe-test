package rest

import (
	"log"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (s *rest) SignIn(c *fiber.Ctx) error {

	ctx := c.Context()

	payload := domain.AuthRequest{}
	err := c.BodyParser(&payload)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message)
	}

	token, err := s.auth.SignIn(ctx, payload)
	if err != nil {
		return err
	}

	s.ResponseJson(c, token, "success", true)
	return nil
}
