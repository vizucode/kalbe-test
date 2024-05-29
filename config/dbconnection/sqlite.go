package dbconnection

import (
	"context"
	"fmt"
	"net/http"

	"api.kalbe.crm/config/logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

func SqliteConn(env *viper.Viper) (resp *sqlstore.Container) {

	fmt.Println(env.GetString("DB_SQLITE_DRIVER"))
	fmt.Println(env.GetString("DB_SQLITE_ADDR"))
	resp, err := sqlstore.New(env.GetString("DB_SQLITE_DRIVER"), env.GetString("DB_SQLITE_ADDR"), nil)
	if err != nil {
		logger.LogError(context.TODO(), err, http.StatusInternalServerError)
		return
	}

	return
}
