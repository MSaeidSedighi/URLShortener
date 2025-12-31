package database

import (
	"log"
	"urlshortener/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully!")
	DB = db

	DB.AutoMigrate(&models.Link{}, &models.User{})
}
