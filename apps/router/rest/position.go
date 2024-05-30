package rest

import (
	"log"
	"strconv"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (rest *rest) CreatePosition(c *fiber.Ctx) error {

	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	var payload domain.Position
	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	payload.CreatedBy = name

	ctx := c.Context()
	if err := rest.position.CreatePosition(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully create position", true)

	return nil
}

func (rest *rest) UpdatePosition(c *fiber.Ctx) error {
	name, ok := c.Locals("name").(string)
	if !ok {
		name = ""
	}

	var payload domain.Position
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
	if err := rest.position.UpdatePosition(ctx, payload); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully update position", true)

	return nil
}

func (rest *rest) DeletePosition(c *fiber.Ctx) error {
	ids := c.Params("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}
	ctx := c.Context()
	if err := rest.position.DeletePosition(ctx, id); err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, nil, "Successfully delete position", true)

	return nil
}

func (rest *rest) GetRowPosition(c *fiber.Ctx) error {

	ids := c.Params("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Message)
	}

	ctx := c.Context()
	resp, err := rest.position.GetPosition(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, resp, "Successfully get position", true)

	return nil
}

func (rest *rest) GetListPosition(c *fiber.Ctx) error {
	ctx := c.Context()
	resp, err := rest.position.GetAllPosition(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	rest.ResponseJson(c, resp, "Successfully get position list", true)

	return nil
}
