package middlewares

import (
	"net/http"
	"strings"

	"BookKhoone/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
