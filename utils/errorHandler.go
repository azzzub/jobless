package utils

import "github.com/gin-gonic/gin"

func ErrorHandler(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"errorCode": status,
		"message":   err.Error(),
	})
}
