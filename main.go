package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/azzzub/jobless/auth"
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/gql"
	"github.com/azzzub/jobless/graph"
	"github.com/azzzub/jobless/graph/generated"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/project"
	"github.com/azzzub/jobless/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	// h := handler.NewDefaultServer()
	h := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &graph.Resolver{}}),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	godotenv.Load()

	db := config.DbConn()
	db.AutoMigrate(&model.Auth{}, &model.Project{})

	// Creating the server
	// Move to gin-gonic framework
	router := gin.Default()
	router.GET("/_gql", playgroundHandler())

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
	// GraphQL gqlgen
	gqlRouter := router.Group("/_gql")
	{
		gqlRouter.Use(utils.AuthMiddleware())
		gqlRouter.POST("/query", graphqlHandler())
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
