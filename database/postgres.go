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

func CreateAnime(a *Anime) (*Anime, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &Anime{}, errors.New("anime not created")
	}
	return a, nil
}

func ReadAnime(id string) (*Anime, error) {
	var anime Anime
	res := db.First(&anime, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("anime not found")
	}
	return &anime, nil
}

func ReadAnimes() ([]*Anime, error) {
	var animes []*Anime
	res := db.Find(&animes)
	if res.Error != nil {
		return nil, errors.New("animes not found")
	}
	return animes, nil
}

func UpdateAnime(anime *Anime) (*Anime, error) {
	var updateAnime Anime
	result := db.Model(&updateAnime).Where(anime.ID).Updates(anime)
	if result.RowsAffected == 0 {
		return &Anime{}, errors.New("anime not updated")
	}
	return &updateAnime, nil

}

func DeleteAnime(id string) error {
	var deleteAnime Anime
	result := db.Where(id).Delete(&deleteAnime)
	if result.RowsAffected == 0 {
		return errors.New("anime data not deleted")
	}
	return nil
}