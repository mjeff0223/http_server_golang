package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Anime struct {
	ID			uint 	`json: "id" gorm:"primary_key`
	Title		string 	`json: "title"`
	Description string 	`json: "description`
	Rate 		int 	`json: rate`
}

func getEnvVariable(key string) string {
	err := godotenv.Load("env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func NewPostgreSQLClient() {
	var (
		host = getEnvVariable("DB_HOST")
		port = getEnvVariable("DB_PORT")
		user = getEnvVariable("DB_USER")
		dbname = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	host,
	port,
	user,
	dbname,
	password,
)

	db, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(Anime{})
}