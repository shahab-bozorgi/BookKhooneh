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

func GetBook(db *gorm.DB, name string) (models.Book, error) {
	var book models.Book
	if err := db.Where("name = ?", name).First(&book).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}
