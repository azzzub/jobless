package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the bearer auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "bearer auth must be privided",
			})
			return
		}

		token := strings.Split(authHeader, " ")
		if len(token) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token, must be bearer token",
			})
			return
		}

		decodedToken, err := TokenValidator(token[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		c.Set("token", decodedToken)
		c.Next()
	}
}
