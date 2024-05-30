package rest

import (
	"log"
	"strconv"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (rest *rest) GetListLocation(c *fiber.Ctx) error {
	ctx := c.Context()
	resp, err := rest.location.GetAllLocation(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, resp, "Successfully get list location", true)
	return nil
}

func (rest *rest) GetRowLocation(c *fiber.Ctx) error {
	ctx := c.Context()
	ids := c.Params("id")

	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	resp, err := rest.location.GetLocation(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}
	rest.ResponseJson(c, resp, "Successfully get row location", true)
	return nil
}

func (rest *rest) CreateLocation(c *fiber.Ctx) error {
	ctx := c.Context()

	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	var payload domain.Location
	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return err
	}
	payload.CreatedBy = name

	if err := rest.location.CreateLocation(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully create location", true)
	return nil
}

func (rest *rest) UpdateLocation(c *fiber.Ctx) error {
	ctx := c.Context()

	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	ids := c.Params("id")

	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	var payload domain.Location
	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return err
	}
	payload.UpdatedBy = name
	payload.ID = id

	if err := rest.location.UpdateLocation(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully update location", true)
	return nil
}

func (rest *rest) DeleteLocation(c *fiber.Ctx) error {
	ctx := c.Context()

	ids := c.Params("id")

	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	if err := rest.location.DeleteLocation(ctx, id); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully delete location", true)
	return nil
}
