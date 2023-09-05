package models

import (
	"log"

	"github.com/yankyhermawan/learn-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	connectionString := "postgresql://postgres:ckjeZw2rpJiYRacMLyo1@containers-us-west-103.railway.app:7742/railway"

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	database.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal(err)
	}

	return database, nil
}
