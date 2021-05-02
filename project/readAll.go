package project

import (
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	"github.com/gofiber/fiber/v2"
)

func ReadAllProject(c *fiber.Ctx) error {
	db := config.DbConn()

	var project []model.Project

	result := db.Find(&project)
	if result.Error != nil {
		return utils.ErrorHandler(c, fiber.StatusBadGateway, result.Error)
	}

	return c.JSON(fiber.Map{
		"data": project,
	})
}
