package main

import (
	"context"
	"net/http"
	"strings"

	errorHandler "api.kalbe.crm/config/error_handler"
	"api.kalbe.crm/config/hisentry"
	"api.kalbe.crm/config/logger"
	"api.kalbe.crm/config/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler.ErrorHandler,
	})
	app.Use(recover.New())

	env := viper.NewViper()
	if strings.EqualFold(env.GetString("APP_ENV"), "production") {
		hisentry.NewHisentry().Init()
	}

	// validator := validator.New()

	// rest.NewRest(

	// ).RegisterRoute(app)

	err := app.Listen(env.GetString("APP_PORT"))
	if err != nil {
		logger.LogError(context.Background(), err, http.StatusInternalServerError)
	}
}
