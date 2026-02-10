package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		if role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "Forbidden: admin access only",
			})
			return
		}

		c.Next()
	}
}
