package middlewares

import (
	"BookKhoone/infrastructure/utils"
	"BookKhoone/internal/domain"
	"gorm.io/gorm"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		var user domain.User
		if err := db.First(&user, userID).Error; err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("user_id", user.ID)
		c.Set("user_role", user.Role)

		c.Next()
	}
}
