package handlers

import (
	"BookKhoone/infrastructure/config"
	"BookKhoone/internal/application"
	"BookKhoone/internal/domain"
	"BookKhoone/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Login user
// @Description Register with username, password and Email and receive JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body dto.RegisterRequest true "Register data"
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} dto.RegisterErrorResponse
// @Failure 401 {object} dto.RegisterErrorResponse
// @Router /auth/register [post]
func RegisterHandler(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.RegisterRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existing domain.User
		if err := db.Where("email = ? OR username = ?", input.Email, input.Username).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		user, err := application.CreateUser(db, input.Username, input.Email, input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		token, err := application.GenerateUserToken(user.ID, cfg.JWTSecretKey)
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
// @Failure 400 {object} dto.LoginErrorResponse
// @Failure 401 {object} dto.LoginErrorResponse
// @Router /auth/login [post]
func LoginHandler(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.LoginErrorResponse{Error: "Invalid request"})
			return
		}

		user, err := application.LoginUser(db, req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.LoginErrorResponse{Error: err.Error()})
			return
		}

		token, err := application.GenerateUserToken(user.ID, cfg.JWTSecretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.LoginErrorResponse{Error: "Failed to generate token"})
			return
		}

		c.JSON(http.StatusCreated, dto.LoginResponse{
			Message: "login successfully",
			Token:   token,
		})
	}
}
