package attandance

import (
	"context"
	"log"
	"time"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (s *attandance) CheckOut(ctx context.Context, payload domain.Attandance) (err error) {

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

	for _, attandace := range attandaces {
		if !attandace.AbsentOut.IsZero() {
			return fiber.NewError(fiber.StatusBadRequest, "Already Absent Out")
		}
	}

	err = s.db.UpdateAttandance(ctx, []string{"updated_by", "absent_out"}, models.Attandance{
		Model:     gorm.Model{ID: uint(payload.Id)},
		UpdatedBy: employeeModel.UpdatedBy,
		AbsentOut: now,
	})
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}
