package employee

import (
	"context"
	"log"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s *employee) UpdateEmployee(ctx context.Context, payload domain.Employee) (err error) {

	err = s.validator.StructCtx(ctx, &payload)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	bPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	err = s.db.UpdateEmployee(ctx, []string{"employee_name", "superior", "departement_id", "position_id"}, models.Employee{
		EmployeeCode:  payload.EmployeeCode,
		Password:      string(bPass),
		EmployeeName:  payload.EmployeeName,
		DepartementId: payload.DepartementId,
		PositionId:    payload.PositionId,
		Superior:      payload.Superior,
	})

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return
}
