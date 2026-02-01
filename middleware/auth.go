package middleware

import (
	"blog-platform/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "authorization header missing",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{
				"error": "Invalid authorization format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		userId, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}

func OptionalMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {

			c.Next()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Next()
			return
		}

		userId, err := utils.ParseToken(parts[1])
		if err != nil {
			c.Next()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
