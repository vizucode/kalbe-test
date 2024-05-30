package rest

import (
	"log"
	"strconv"
	"time"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (s *rest) CheckIn(c *fiber.Ctx) error {
	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	employeeCode, ok := c.Locals("employee_code").(string)
	if !ok {
		employeeCode = ""
	}

	var payload domain.Attandance
	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.CreatedBy = name
	payload.EmployeeCode = employeeCode

	err := s.attandace.CheckIn(c.Context(), payload)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	s.ResponseJson(c, nil, "Successfully create attandance", true)

	return nil
}

func (s *rest) CheckOut(c *fiber.Ctx) error {
	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	employeeCode, ok := c.Locals("employee_code").(string)
	if !ok {
		employeeCode = ""
	}

	var payload domain.Attandance
	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.CreatedBy = name
	payload.EmployeeCode = employeeCode

	err := s.attandace.CheckOut(c.Context(), payload)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	s.ResponseJson(c, nil, "Successfully checkout attandance", true)

	return nil
}

func (s *rest) GetAllAttandance(c *fiber.Ctx) error {

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)

	attandaces, err := s.attandace.GetAllAttandance(c.Context(), startDate, endDate)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	s.ResponseJson(c, attandaces, "Successfully get all attandance", true)

	return nil
}

func (s *rest) GetRowAttandance(c *fiber.Ctx) error {
	ctx := c.Context()

	ids := c.Params("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	attandanceRow, err := s.attandace.GetRowAttandance(ctx, id)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	s.ResponseJson(c, attandanceRow, "Successfully get attandance", true)

	return nil
}
