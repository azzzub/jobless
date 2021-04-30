package main

import (
	"github.com/azzzub/jobless/auth"
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.DbConn()

	db.AutoMigrate(&model.Auth{})

	// Creating the server
	app := fiber.New()

	v1 := app.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	authRouteV1 := v1.Group("/auth")

	authRouteV1.Post("/login", auth.LoginHandler)
	authRouteV1.Post("/register", auth.RegisterHandler)

	app.Listen(":9000")
}