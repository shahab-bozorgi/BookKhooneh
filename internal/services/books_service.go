package services

import (
	"BookKhoone/internal/dto"
	"BookKhoone/internal/models"
	"errors"
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

func GetBook(db *gorm.DB, id string) (models.Book, error) {
	var book models.Book
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
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

func FilterBookService(db *gorm.DB, filter dto.FilterBooksRequest) ([]dto.BookResponse, error) {
	var books []models.Book
	var response []dto.BookResponse

	if len(filter.Author) == 0 && len(filter.Title) == 0 {
		return nil, errors.New("at least one filter is required")
	}

	query := db.Model(&models.Book{})

	if len(filter.Author) > 0 {
		query = query.Where("author IN ?", filter.Author)
	}

	if len(filter.Title) > 0 {
		query = query.Where("title IN ?", filter.Title)
	}

	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}

	for _, book := range books {
		response = append(response, dto.BookResponse{
			Title:  book.Title,
			Author: book.Author,
		})
	}

	return response, nil
}
