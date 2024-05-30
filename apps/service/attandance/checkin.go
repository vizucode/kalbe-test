package attandance

import (
	"context"
	"log"
	"time"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
)

func (s *attandance) CheckIn(ctx context.Context, payload domain.Attandance) (err error) {

	err = s.validator.StructCtx(ctx, &payload)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	employeeModel, err := s.db.GetEmployee(ctx, payload.EmployeeCode)
	if err != nil {
		log.Println(err)
		return err
	}

	// check if already absent this day
	now := time.Now()

	attandaces, err := s.db.GetAllAttandance(ctx, now, now.Add(time.Hour*24))
	if err != nil {
		log.Println(err)
		return err
	}

	if len(attandaces) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Already Absent In")
	}

	err = s.db.CreateAttandance(ctx, models.Attandance{
		LocationId: payload.LocationId,
		EmployeeId: int(employeeModel.ID),
		CreatedBy:  employeeModel.CreatedBy,
		AbsentIn:   now,
	})
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return
}
