package main

import (
	"context"
	"net/http"
	"strings"

	psqldatabase "api.kalbe.crm/apps/repositories/psql_database"
	"api.kalbe.crm/apps/router/rest"
	"api.kalbe.crm/apps/service/auth"
	"api.kalbe.crm/apps/service/departement"
	"api.kalbe.crm/apps/service/location"
	"api.kalbe.crm/apps/service/position"
	"api.kalbe.crm/config/dbconnection"
	errorHandler "api.kalbe.crm/config/error_handler"
	"api.kalbe.crm/config/hisentry"
	"api.kalbe.crm/config/logger"
	"api.kalbe.crm/config/viper"
	"github.com/go-playground/validator/v10"
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

	db := dbconnection.DbConn()
	psql := psqldatabase.NewPsql(db)

	validator := validator.New()

	rest.NewRest(
		auth.NewAuth(psql, validator),
		departement.NewDepartement(psql, validator),
		position.NewPosition(psql, validator),
		location.NewLocation(psql, validator),
	).RegisterRoute(app)

	err := app.Listen(env.GetString("APP_PORT"))
	if err != nil {
		logger.LogError(context.Background(), err, http.StatusInternalServerError)
	}
}
