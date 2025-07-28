package handlers

import (
	"BookKhoone/internal/models"
	"BookKhoone/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"

	"gorm.io/gorm"
)

type Books struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description"`
	UserID      *uint  `json:"user_id"`
}

type BookResponse struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func CreateBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input Books

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book := models.Book{
			Title:       input.Title,
			Author:      input.Author,
			Description: input.Description,
			UserId:      input.UserID,
		}

		createdBook, err := services.CreateBook(db, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"title":       createdBook.Title,
			"author":      createdBook.Author,
			"description": createdBook.Description,
			"user_id":     createdBook.UserId,
		})
	}

}

func GetAllBooksHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := services.GetAllBooks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var response []BookResponse
		for _, book := range books {
			response = append(response, BookResponse{
				Title:  book.Title,
				Author: book.Author,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"books": response,
		})
	}
}
