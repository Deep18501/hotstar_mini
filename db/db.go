package db

import (
	"log"
	"os"

	"github.com/Deep18501/hotstar_mini/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := os.Getenv("GORM_DB_URL")
	if dsn == "" {
		log.Fatal("gorm db url not found in environemnt")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting to database", err)
	}
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Genre{})
	db.AutoMigrate(&models.Media{})
	db.AutoMigrate(&models.Rating{})

	return db
}
