package database

import (
	"fmt"
	"log"
	"os"

	"example.com/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loadinf env file")
	}

	dbCon := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dbCon), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database %v", err)
	}

	//run the migration
	db.AutoMigrate(&models.User{}, &models.Profile{})
	DB = db
}
