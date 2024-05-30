package psqldatabase

import (
	"context"
	"log"
	"net/http"

	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
)

func (uc *psql) GetAllPosition(ctx context.Context) (resp []models.Position, err error) {
	if tx := uc.db.Model(&models.Position{}).Find(&resp); tx.Error != nil {
		log.Println(tx.Error)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (uc *psql) GetPosition(ctx context.Context, id int) (resp models.Position, err error) {
	if tx := uc.db.Model(&models.Position{}).Where("id = ?", id).First(&resp); tx.Error != nil {
		log.Println(tx.Error)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (uc *psql) CreatePosition(ctx context.Context, payload models.Position) (err error) {
	if tx := uc.db.Model(&models.Position{}).Create(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (uc *psql) UpdatePosition(ctx context.Context, selectedField []string, payload models.Position) (err error) {
	if tx := uc.db.Model(&models.Position{}).Select(selectedField).Where("id = ?", payload.ID).Updates(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (uc *psql) DeletePosition(ctx context.Context, id int) (err error) {
	if tx := uc.db.Model(&models.Position{}).Delete(&models.Position{}, id); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}
