package handlers

import (
	"BookKhoone/internal/models"
	"BookKhoone/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Books struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description"`
}

type BookResponse struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func CreateBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDValue, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}

		userIDStr, ok := userIDValue.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
			return
		}

		userIDUint64, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
			return
		}
		userID := uint(userIDUint64)

		var input Books
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book := models.Book{
			Title:       input.Title,
			Author:      input.Author,
			Description: input.Description,
			UserID:      &userID,
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
			"user_id":     createdBook.UserID,
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

func GetBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Param("title")

		book, err := services.GetBook(db, title)
		if err != nil {
			c.JSON(404, gin.H{"error": "Book not found in library"})
			return
		}

		c.JSON(200, gin.H{"book": book})
	}
}

func UpdateBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		bookID, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		updatedBook, err := services.UpdateBook(db, uint(bookID), input)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or update failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"title":       updatedBook.Title,
			"author":      updatedBook.Author,
			"description": updatedBook.Description,
			"user_id":     updatedBook.UserID,
		})
	}
}

func DeleteBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		bookID, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		if err := services.DeleteBook(db, uint(bookID)); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or deletion failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}
