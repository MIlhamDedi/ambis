package db

import (
	"ambis/lib/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func New(appConfig *config.Config) *sql.DB {
	db, err := sql.Open(appConfig.DBMode, fmt.Sprintf("%s:%s@/%s", appConfig.DBUser, appConfig.DBPassword, appConfig.DBName))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
