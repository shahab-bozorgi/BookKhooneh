package handlers

import (
	"BookKhoone/internal/dto"
	"BookKhoone/internal/services"
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

		user, err := services.GetUserService(db, username)
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
		users, err := services.GetAllUsersService(db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var response []dto.AllUsersResponse
		for _, user := range users {
			response = append(response, dto.AllUsersResponse{
				Username: user.Username,
				Email:    user.Email,
			})
		}

		c.JSON(200, gin.H{
			"users": response,
		})

	}
}
