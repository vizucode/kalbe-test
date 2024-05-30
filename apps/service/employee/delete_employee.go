package employee

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (s *employee) DeleteEmployee(ctx context.Context, employeeCode string) (err error) {

	err = s.db.DeleteEmployee(ctx, employeeCode)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return
}
