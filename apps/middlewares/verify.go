package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"api.kalbe.crm/apps/domain"
	"api.kalbe.crm/config/logger"
	"api.kalbe.crm/config/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	secretKey := viper.NewViper().GetString("JWT_SECRET_KEY")

	tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		logger.LogError(c.Context(), err, http.StatusBadRequest)
		c.Status(http.StatusBadRequest).JSON(domain.Response{
			Data: map[string]interface{}{
				"token_verified": false,
			},
			Message: jwt.ErrInvalidKey.Error(),
			Status:  false,
		})
		return nil
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("name", claims["name"])
	c.Locals("employee_code", claims["employee_code"])

	return c.Next()
}
