package main

import (
	"net/http"

	"github.com/azzzub/jobless/auth"
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/gql"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/project"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := config.DbConn()
	db.AutoMigrate(&model.Auth{}, &model.Project{})

	// Creating the server
	// Move to gin-gonic framework
	router := gin.Default()

	// Version 1 API
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "status ok",
			})
		})
		// GraphQL
		v1.GET("/gql", gql.GraphQL())
		v1.POST("/gql", gql.GraphQL())
	}
	// Auth router V1
	authRouterV1 := v1.Group("/auth")
	{
		authRouterV1.POST("/login", auth.LoginHandler)
		authRouterV1.POST("/register", auth.RegisterHandler)
	}
	// Project router V1
	projectRouterV1 := v1.Group("/project")
	{
		projectRouterV1.GET("/", project.ReadAllProject)
		projectRouterV1.POST("/", project.CreateProject)
	}

	// Run on port 9000
	router.Run(":9000")
}
