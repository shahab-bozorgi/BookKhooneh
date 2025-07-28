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

func GetAllBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
