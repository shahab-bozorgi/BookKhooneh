package application

import (
	"BookKhoone/internal/domain"
	"BookKhoone/internal/dto"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"math"
)

func CreateBook(db *gorm.DB, book domain.Book) (*domain.Book, error) {
	if err := db.Create(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func GetAllBooksService(db *gorm.DB, books []domain.Book) ([]dto.BookResponse, error) {
	var result []dto.BookResponse

	for _, book := range books {
		var avg sql.NullFloat64
		db.Model(&domain.Review{}).
			Where("book_id = ?", book.ID).
			Select("AVG(rating)").
			Scan(&avg)

		avgRating := 0.0
		if avg.Valid {
			avgRating = math.Round(avg.Float64*10) / 10
		}

		result = append(result, dto.BookResponse{
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			AvgRating:   avgRating,
		})
	}

	return result, nil
}

func GetBook(db *gorm.DB, id string) (*dto.BookWithStats, error) {
	var book domain.Book
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	var avgRating float64

	db.Model(&domain.Review{}).Where("book_id = ?", book.ID).Select("AVG(rating)").Scan(&avgRating)
	avgRating = math.Round(avgRating*10) / 10

	return &dto.BookWithStats{
		Book:      book,
		AvgRating: avgRating,
	}, nil
}

func UpdateBook(db *gorm.DB, id uint, updatedData map[string]interface{}) (*domain.Book, error) {
	var book domain.Book

	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&book).Updates(updatedData).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func DeleteBook(db *gorm.DB, id uint) error {
	var book domain.Book

	if err := db.First(&book, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

func FilterBookService(db *gorm.DB, filter dto.FilterBooksRequest) ([]dto.BookResponse, error) {
	var books []domain.Book
	var response []dto.BookResponse

	if len(filter.Author) == 0 && len(filter.Title) == 0 {
		return nil, errors.New("at least one filter is required")
	}

	query := db.Model(&domain.Book{})

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
