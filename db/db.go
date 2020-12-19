package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
)

func Init() {

	EnvLoad()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	url := os.Getenv("DB_URL")
	port := os.Getenv("DB_PORT")

	// DB接続
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+url+":"+port+")/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}

func EnvLoad() {

	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf("../.env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file" + os.Getenv("GO_ENV"))
	}
}
