package handlers

import (
	"BookKhoone/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Users struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type AllUsersResponse struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
}

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

func GetAllUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := services.GetAllUsersService(db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var response []AllUsersResponse
		for _, user := range users {
			response = append(response, AllUsersResponse{
				Username: user.Username,
				Email:    user.Email,
			})
		}

		c.JSON(200, gin.H{
			"users": response,
		})

	}
}
