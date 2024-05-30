package psqldatabase

import (
	"context"
	"log"
	"net/http"

	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
)

func (psql *psql) GetAllDepartement(ctx context.Context) (resp []models.Departement, err error) {
	if tx := psql.db.Model(&models.Departement{}).Find(&resp); tx.Error != nil {
		log.Println(tx.Error)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}
	return resp, nil
}

func (psql *psql) GetDepartement(ctx context.Context, id int) (resp models.Departement, err error) {

	if tx := psql.db.Model(&models.Departement{}).Where("id = ?", id).First(&resp); tx.Error != nil {
		log.Println(tx.Error)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, err
}

func (psql *psql) CreateDepartement(ctx context.Context, payload models.Departement) (err error) {

	if tx := psql.db.Model(&models.Departement{}).Create(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (psql *psql) UpdateDepartement(ctx context.Context, selectedField []string, payload models.Departement) (err error) {

	if tx := psql.db.Model(&models.Departement{}).Select(selectedField).Where("id = ?", payload.ID).Updates(&payload); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (psql *psql) DeleteDepartement(ctx context.Context, id int) (err error) {

	if tx := psql.db.Model(&models.Departement{}).Delete(&models.Departement{}, id); tx.Error != nil {
		log.Println(tx.Error)
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}
