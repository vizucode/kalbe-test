package rest

import (
	"log"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (rest *rest) CreateEmployee(c *fiber.Ctx) error {
	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	var payload domain.Employee

	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.CreatedBy = name

	err := rest.employee.CreateEmployee(c.Context(), payload)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	rest.ResponseJson(c, nil, "Successfully create employee", true)

	return nil
}

func (rest *rest) UpdateEmployee(c *fiber.Ctx) error {
	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}
	var payload domain.Employee

	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.UpdatedBy = name
	payload.EmployeeCode = c.Params("employee_code")

	err := rest.employee.UpdateEmployee(c.Context(), payload)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	rest.ResponseJson(c, nil, "Successfully update employee", true)

	return nil
}

func (rest *rest) DeleteEmployee(c *fiber.Ctx) error {
	err := rest.employee.DeleteEmployee(c.Context(), c.Params("employee_code"))
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	rest.ResponseJson(c, nil, "Successfully delete employee", true)

	return nil
}

func (rest *rest) GetAllEmployee(c *fiber.Ctx) error {
	resp, err := rest.employee.GetAllEmployee(c.Context())
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	rest.ResponseJson(c, resp, "Successfully get all employee", true)

	return nil
}

func (rest *rest) GetEmployee(c *fiber.Ctx) error {
	resp, err := rest.employee.GetRowEmployee(c.Context(), c.Params("employee_code"))
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	rest.ResponseJson(c, resp, "Successfully get employee", true)

	return nil
}
