package services

import (
	"BookKhoone/internal/models"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book models.Book) (*models.Book, error) {
	if err := db.Create(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
