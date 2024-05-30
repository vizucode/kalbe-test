package psqldatabase

import (
	"context"
	"log"

	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
)

func (psql *psql) GetAllLocation(ctx context.Context) (resp []models.Location, err error) {
	if tx := psql.db.Model(&models.Location{}).Find(&resp); tx.Error != nil {
		log.Println(tx.Error)
		return nil, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (psql *psql) GetLocation(ctx context.Context, id int) (resp models.Location, err error) {

	if tx := psql.db.Model(&models.Location{}).Where("id = ?", id).First(&resp); tx.Error != nil {
		log.Println(tx.Error)
		return models.Location{}, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (psql *psql) CreateLocation(ctx context.Context, payload models.Location) (err error) {
	if tx := psql.db.Model(&models.Location{}).Create(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (psql *psql) UpdateLocation(ctx context.Context, selectedString []string, payload models.Location) (err error) {
	if tx := psql.db.Model(&models.Location{}).Select(selectedString).Where("id = ?", payload.ID).Updates(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}
	return nil
}

func (psql *psql) DeleteLocation(ctx context.Context, id int) (err error) {
	if tx := psql.db.Model(&models.Location{}).Delete(&models.Location{}, id); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}
