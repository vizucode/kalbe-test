package dbconnection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"api.kalbe.crm/config/logger"
	"api.kalbe.crm/config/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConn() (db *gorm.DB) {
	v := viper.NewViper()

	dbhost := v.GetString("DB_HOST")
	dbuser := v.GetString("DB_USER")
	dbpassword := v.GetString("DB_PASSWORD")
	dbname := v.GetString("DB_NAME")
	dbport := v.GetInt("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

	return
}
