package psqldatabase

import (
	"context"
	"errors"
	"log"
	"net/http"

	"api.kalbe.crm/apps/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (psql *psql) GetAllEmployee(ctx context.Context) (resp []models.Employee, err error) {
	if tx := psql.db.Model(&models.Employee{}).Find(&resp); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return resp, fiber.NewError(http.StatusNotFound, fiber.ErrNotFound.Message)
		}
		log.Println(tx.Error)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}
func (psql *psql) GetEmployee(ctx context.Context, employeeCode string) (resp models.Employee, err error) {
	if tx := psql.db.Model(&models.Employee{}).Where("employee_code = ?", employeeCode).First(&resp); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return resp, fiber.NewError(http.StatusNotFound, fiber.ErrNotFound.Message)
		}
		log.Println(tx.Error)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return resp, nil
}

func (psql *psql) CreateEmployee(ctx context.Context, payload models.Employee) (err error) {
	if tx := psql.db.Model(&models.Employee{}).Create(&payload); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (psql *psql) UpdateEmployee(ctx context.Context, selectedString []string, payload models.Employee) (err error) {
	if tx := psql.db.Model(&models.Employee{}).Select(selectedString).Where("employee_code = ?", payload.EmployeeCode).Updates(&payload); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}

func (psql *psql) DeleteEmployee(ctx context.Context, employeeCode string) (err error) {
	if tx := psql.db.Model(&models.Employee{}).Where("employee_code = ?", employeeCode).Delete(&models.Employee{}); tx.Error != nil {
		return fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	return nil
}
