package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Gthamsrim1/forms-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`).Error; err != nil {
		return fmt.Errorf("failed to create pgcrypto extension: %w", err)
	}
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "citext";`).Error; err != nil {
		return fmt.Errorf("failed to create citext extension: %w", err)
	}

	if err := DB.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Form{},
		&models.Question{},
	); err != nil {
		return fmt.Errorf("failed to migrate tables: %w", err)
	}

	log.Println("Database connected and migrated successfully!")
	return nil
}
