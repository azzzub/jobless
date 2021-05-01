package utils

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message": err.Error(),
	})
}
