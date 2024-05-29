package dbconnection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"api.kalbe.crm/apps/models"
	"api.kalbe.crm/config/logger"
	"api.kalbe.crm/config/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() (db *gorm.DB) {
	v := viper.NewViper()

	dbhost := v.GetString("DB_HOST")
	dbuser := v.GetString("DB_USER")
	dbpassword := v.GetString("DB_PASSWORD")
	dbname := v.GetString("DB_NAME")
	dbport := v.GetInt("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbhost, dbuser, dbpassword, dbname, dbport)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		logger.LogError(context.Background(), errors.New(err.Error()), http.StatusInternalServerError)
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		logger.LogError(context.Background(), errors.New(err.Error()), http.StatusInternalServerError)
		return
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Minute)

	db.AutoMigrate(
		models.Departement{},
		models.Location{},
		models.Position{},
		models.Employee{},
		models.Attandance{},
	)

	return db
}
