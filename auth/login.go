package auth

import (
	"github.com/gofiber/fiber/v2"
	// database "github.com/azzzub/jobless/config"
)

func LoginHandler(c *fiber.Ctx) error {
	// db := database.DbConn()

	return c.SendString("Login")
}
