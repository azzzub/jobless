package project

import (
	"net/http"

	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func ReadAllProject(c *gin.Context) {
	db := config.DbConn()

	var project []model.Project

	result := db.Find(&project)
	if result.Error != nil {
		utils.ErrorHandler(c, fiber.StatusBadGateway, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}
