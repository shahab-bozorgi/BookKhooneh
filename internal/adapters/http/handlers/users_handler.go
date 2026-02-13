package handlers

import (
	"BookKhoone/infrastructure/utils"
	"BookKhoone/internal/application"
	"BookKhoone/internal/domain"
	"BookKhoone/internal/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetUserHandler godoc
// @Summary Get user by username
// @Description Get user details by providing the username
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "Username of the user"
// @Success 200 {object} dto.GetUserResponse
// @Failure 404 {object} dto.UserErrorResponse
// @Router /users/{username} [get]
func GetUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		user, err := application.GetUserService(db, username)
		if err != nil {
			c.JSON(404, gin.H{"message": "user not found"})
			return
		}

		c.JSON(200, gin.H{"user": user})
	}

}

// GetAllUsersHandler godoc
// @Summary Get all users
// @Description Get all user details (requires authentication)
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.AllUsersResponse
// @Failure 401 {object} dto.UserErrorResponse
// @Failure 404 {object} dto.UserErrorResponse
// @Router /users/get_all [get]
func GetAllUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		paginationResult, err := utils.Paginate[domain.User](c, db.Model(&domain.User{}))
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{Error: err.Error()})
			return
		}

		var response []dto.AllUsersResponse
		for _, user := range paginationResult.Data {
			response = append(response, dto.AllUsersResponse{
				Username: user.Username,
				Email:    user.Email,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"page":  paginationResult.Page,
			"size":  paginationResult.Size,
			"total": paginationResult.Total,
			"users": response,
		})
	}
}
