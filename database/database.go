package database

import (
	"fmt"
	"gomod/helpers"
	"gomod/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func getDbValues() {
	err := godotenv.Load()
	helpers.CheckError(err)
}
func Init() {

	getDbValues()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	DB.AutoMigrate(models.Blog{}, models.Comment{})
}
