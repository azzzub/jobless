package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/azzzub/jobless/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the bearer auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "bearer auth must be provided",
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

		c.Set("UserID", decodedToken)
		c.Next()
	}
}

// Use to process the authorization header
// returning the token model and error
func AuthMiddlewareProc(c *gin.Context) (*model.Token, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("bearer auth must be provided")
	}

	token := strings.Split(authHeader, " ")
	if len(token) < 2 {
		return nil, errors.New("invalid token, must be bearer token")
	}

	decodedToken, err := TokenValidator(token[1])
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return decodedToken, nil
}
