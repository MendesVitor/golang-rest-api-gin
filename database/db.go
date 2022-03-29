package database

import (
	"api-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnDB() {
	connString := "host=localhost user=root password=root dbname=root port=5431 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Panic("DB connection error")
	}

	DB.AutoMigrate(&models.Student{})
}
