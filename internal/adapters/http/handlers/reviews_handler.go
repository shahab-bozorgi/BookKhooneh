package handlers

import (
	"BookKhoone/internal/application"
	"BookKhoone/internal/domain"
	"BookKhoone/internal/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreateReviewBookHandler godoc
// @Summary Create review for a book
// @Description Create a review for a specific book. User must be authenticated.
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateReviewRequest true "Review data"
// @Success 200 {object} dto.CreateReviewResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /review/create [post]
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

		var book domain.Book
		if err := db.First(&book, input.BookID).Error; err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Book not found"})
			return
		}

		var user domain.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "User not found"})
			return
		}

		review := domain.Review{
			BookID:  input.BookID,
			UserID:  userID,
			Rating:  input.Rating,
			Comment: input.Comment,
		}

		createdReview, err := application.CreateBookReviewsService(db, review)
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
