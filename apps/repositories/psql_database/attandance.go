package psqldatabase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (psql *psql) GetAllAttandance(ctx context.Context, startDate time.Time, endDate time.Time) (resp []models.Attandance, err error) {

	tx := psql.db.Model(&models.Attandance{})

	if !startDate.IsZero() {
		tx = tx.Where("created_at >= ?", startDate)
	}

	if !endDate.IsZero() {
		tx = tx.Where("created_at <= ?", endDate)
	}

	tx = tx.Preload("Employee.Position").Preload("Employee.Departement").Preload("Location").Find(&resp)
	if tx.Error != nil {
		log.Println(tx.Error)
		return resp, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (psql *psql) GetAttandance(ctx context.Context, id int) (resp models.Attandance, err error) {
	if tx := psql.db.Where("id = ?", id).Preload("Employee.Position").Preload("Employee.Departement").Preload("Location").First(&resp); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return resp, fiber.NewError(fiber.StatusNotFound, fiber.ErrNotFound.Message)
		}
		log.Println(tx.Error)
		return resp, fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (psql *psql) CreateAttandance(ctx context.Context, payload models.Attandance) (err error) {
	if tx := psql.db.Create(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (psql *psql) UpdateAttandance(ctx context.Context, selectedField []string, payload models.Attandance) (err error) {
	fmt.Println(payload.UpdatedBy)
	if tx := psql.db.Model(&models.Attandance{}).Select(selectedField).Where("id = ?", payload.ID).Updates(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (psql *psql) DeleteAttandance(ctx context.Context, id int) (err error) {
	if tx := psql.db.Where("id = ?", id).Delete(&models.Attandance{}); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}
