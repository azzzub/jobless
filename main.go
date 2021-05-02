package main

import (
	"github.com/azzzub/jobless/auth"
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/project"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := config.DbConn()
	db.AutoMigrate(&model.Auth{}, &model.Project{})

	// Creating the server
	app := fiber.New()
	v1 := app.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	// Auth router V1

	authRouteV1 := v1.Group("/auth")
	authRouteV1.Post("/login", auth.LoginHandler)
	authRouteV1.Post("/register", auth.RegisterHandler)

	// Project router V1

	projectRouteV1 := v1.Group("/project")
	projectRouteV1.Post("/", project.CreateProject)
	projectRouteV1.Get("/", project.ReadAllProject)

	app.Listen(":9000")
}
