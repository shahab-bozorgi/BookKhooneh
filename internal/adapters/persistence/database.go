package persistence

import (
	"BookKhoone/infrastructure/config"
	"fmt"
	"log"

	"BookKhoone/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		//hard code nabashe
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to adapters: %v", err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Book{}, &domain.Review{})
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	return db
}
