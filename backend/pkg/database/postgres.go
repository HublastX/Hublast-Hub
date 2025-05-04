package database

import (
	"fmt"
	"log"

	"github.com/HublastX/HubLast-Hub/config"
	"github.com/HublastX/HubLast-Hub/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	cfg := config.LoadConfig()

	dbURL := cfg.DB_URL
	fmt.Println("DB_URL:", dbURL)
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
	DB = database

	err = DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Roadmap{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database migration successful")
}
