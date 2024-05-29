package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/config/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (auth *auth) SignIn(ctx context.Context, payload domain.AuthRequest) (token string, err error) {

	err = auth.validation.StructCtx(ctx, &payload)
	if err != nil {
		log.Println(err)
		return "", err
	}

	employeeModel, err := auth.db.GetEmployee(ctx, payload.EmployeeCode)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if employeeModel.ID > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(employeeModel.Password), []byte(payload.Password))
		if err != nil {
			log.Println(err)
			return "", fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
		}

		tokenizer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"employee_code": payload.EmployeeCode,
			"exp":           time.Now().Add(time.Hour * 24).Unix(),
		})

		secretKey := []byte(viper.NewViper().GetString("SECRET_JWT_KEY"))
		token, err := tokenizer.SignedString(secretKey)
		if err != nil {
			log.Println(err)
			return "", fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
		}

		return token, nil
	}

	return "", nil
}
