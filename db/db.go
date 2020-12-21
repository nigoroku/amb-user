package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/volatiletech/sqlboiler/boil"
)

type DbEnv struct {
	DbUrl      string `envconfig:"DB_URL"`
	DbPort     string `envconfig:"DB_PORT"`
	DbName     string `envconfig:"DB_NAME"`
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
}

func Init() {
	var dbEnv DbEnv
	envconfig.Process("", &dbEnv)

	// DB接続
	fmt.Println(dbEnv.DbUrl)
	db, err := sql.Open("mysql", dbEnv.DbUser+":"+dbEnv.DbPassword+"@tcp("+dbEnv.DbUrl+":"+dbEnv.DbPort+")/"+dbEnv.DbName+"?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}
