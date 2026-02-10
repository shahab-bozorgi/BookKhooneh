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

func GetBook(db *gorm.DB, title string) (models.Book, error) {
	var book models.Book
	if err := db.Where("title = ?", title).First(&book).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func UpdateBook(db *gorm.DB, id uint, updatedData map[string]interface{}) (*models.Book, error) {
	var book models.Book

	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&book).Updates(updatedData).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func DeleteBook(db *gorm.DB, id uint) error {
	var book models.Book

	if err := db.First(&book, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
