package handlers

import (
	"BookKhoone/internal/config"
	"BookKhoone/internal/dto"
	"BookKhoone/internal/models"
	"BookKhoone/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterHandler(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input RegisterInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existing models.User
		if err := db.Where("email = ? OR username = ?", input.Email, input.Username).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		user, err := services.CreateUser(db, input.Username, input.Email, input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		token, err := services.GenerateUserToken(user.ID, cfg.JWTSecretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"token":   token,
		})
	}
}

// Login godoc
// @Summary Login user
// @Description Login with username and password and receive JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body dto.LoginRequest true "Login data"
// @Success 201 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/login [post]
func LoginHandler(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request"})
			return
		}

		user, err := services.LoginUser(db, req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: err.Error()})
			return
		}

		token, err := services.GenerateUserToken(user.ID, cfg.JWTSecretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to generate token"})
			return
		}

		c.JSON(http.StatusCreated, dto.LoginResponse{
			Message: "login successfully",
			Token:   token,
		})
	}
}
