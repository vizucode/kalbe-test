package psqldatabase

import (
	"context"
	"net/http"

	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
)

func (psql *psql) GetAllDepartement(ctx context.Context) (resp []models.Departement, err error) {
	if tx := psql.db.Model(&models.Departement{}).Find(&resp); tx.Error != nil {
		return resp, fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}
	return resp, nil
}

func (psql *psql) GetDepartement(ctx context.Context, id int) (resp models.Departement, err error) {

	if tx := psql.db.Model(&models.Departement{}).Where("id = ?", id).First(&resp); tx.Error != nil {
		return resp, fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return resp, err
}

func (psql *psql) CreateDepartement(ctx context.Context, payload models.Departement) (err error) {

	if tx := psql.db.Model(&models.Departement{}).Create(&payload); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return nil
}

func (psql *psql) UpdateDepartement(ctx context.Context, selectedField []string, payload models.Departement) (err error) {

	if tx := psql.db.Model(&models.Departement{}).Select(selectedField).Updates(&payload); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return nil
}

func (psql *psql) DeleteDepartement(ctx context.Context, id int) (err error) {

	if tx := psql.db.Model(&models.Departement{}).Delete(&models.Departement{}, id); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return nil
}
