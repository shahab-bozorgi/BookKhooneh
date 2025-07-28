package handlers

import (
	"BookKhoone/internal/models"
	"BookKhoone/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"

	"gorm.io/gorm"
)

type CreateBook struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description"`
	UserID      *uint  `json:"user_id"`
}

func CreateBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateBook

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
