package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/volatiletech/sqlboiler/boil"
)

type DbEnv struct {
	DB_URL      string `default:"localhost"`
	DB_PORT     string `default:"3306"`
	DB_NAME     string `default:"ambitious"`
	DB_USER     string `default:"root"`
	DB_PASSWORD string `default:"base0210"`
}

func Init() {
	var dbEnv DbEnv
	envconfig.Process("", &dbEnv)

	// DB接続
	db, err := sql.Open("mysql", dbEnv.DB_USER+":"+dbEnv.DB_PASSWORD+"@tcp("+dbEnv.DB_URL+":"+dbEnv.DB_PORT+")/"+dbEnv.DB_NAME+"?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}
