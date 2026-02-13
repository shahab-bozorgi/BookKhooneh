package handlers

import (
	"BookKhoone/internal/dto"
	"BookKhoone/internal/models"
	"BookKhoone/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateReviewBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDValue, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "User not found in context"})
			return
		}
		userID, ok := userIDValue.(uint)
		if !ok {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Invalid user ID format"})
			return
		}

		var input dto.CreateReviewRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		if input.Rating < 1 || input.Rating > 10 {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Rating must be between 1 and 10"})
			return
		}

		var book models.Book
		if err := db.First(&book, input.BookID).Error; err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Book not found"})
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "User not found"})
			return
		}

		review := models.Review{
			BookID:  input.BookID,
			UserID:  userID,
			Rating:  input.Rating,
			Comment: input.Comment,
		}

		createdReview, err := services.CreateBookReviewsService(db, review)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.CreateReviewResponse{
			ID:     createdReview.ID,
			BookID: createdReview.BookID,
			User: dto.UserResponse{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
			Rating:  createdReview.Rating,
			Comment: createdReview.Comment,
		})
	}
}
