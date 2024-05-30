package rest

import (
	"log"
	"strconv"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (rest *rest) CreateDepartement(c *fiber.Ctx) error {
	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	var payload domain.Departement

	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.CreatedBy = name

	ctx := c.Context()
	if err := rest.departement.CreateDepartement(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully create departement", true)

	return nil
}

func (rest *rest) UpdateDepartement(c *fiber.Ctx) error {

	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	var payload domain.Departement

	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.UpdatedBy = name

	ids := c.Params("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.Id = id

	ctx := c.Context()
	if err := rest.departement.UpdateDepartement(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully update departement", true)

	return nil
}

func (rest *rest) DeleteDepartement(c *fiber.Ctx) error {
	ids := c.Params("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	ctx := c.Context()
	if err := rest.departement.DeleteDepartement(ctx, id); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully delete departement", true)

	return nil
}

func (rest *rest) GetListDepartement(c *fiber.Ctx) error {

	ctx := c.Context()
	resp, err := rest.departement.GetAllDepartement(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, resp, "Successfully get departement list", true)

	return nil
}

func (rest *rest) GetRowDepartement(c *fiber.Ctx) error {

	ids := c.Params("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	ctx := c.Context()
	resp, err := rest.departement.GetDepartement(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, resp, "Successfully get departement", true)

	return nil
}
