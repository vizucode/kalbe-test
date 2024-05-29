package psqldatabase

import (
	"context"
	"log"
	"net/http"

	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
)

func (psql *psql) GetAllEmployee(ctx context.Context) (resp []models.Employee, err error) {
	if tx := psql.db.Model(&models.Employee{}).Find(&resp); tx.Error != nil {
		return resp, fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return resp, nil
}
func (psql *psql) GetEmployee(ctx context.Context, employeeCode string) (resp models.Employee, err error) {
	if tx := psql.db.Model(&models.Employee{}).Where("employee_code = ?", employeeCode).First(&resp); tx.Error != nil {
		log.Println(tx.Error)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (psql *psql) CreateEmployee(ctx context.Context, payload models.Employee) (err error) {
	if tx := psql.db.Model(&models.Employee{}).Create(&payload); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return nil
}

func (psql *psql) UpdateEmployee(ctx context.Context, selectedString []string, payload models.Location) (err error) {
	if tx := psql.db.Model(&models.Employee{}).Select(selectedString).Updates(&payload); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return nil
}

func (psql *psql) DeleteEmployee(ctx context.Context, id int) (err error) {
	if tx := psql.db.Model(&models.Employee{}).Delete(&models.Employee{}, id); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, tx.Error.Error())
	}

	return nil
}
