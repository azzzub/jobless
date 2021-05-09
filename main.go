package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/database"
	"github.com/azzzub/jobless/database/mock"
	"github.com/azzzub/jobless/graph"
	"github.com/azzzub/jobless/graph/generated"
	"github.com/azzzub/jobless/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
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
	cmdApp := cli.NewApp()

	// CLI command for mocking
	// Usage: go run main.go <COMMAND>
	cmdApp.Commands = []*cli.Command{
		{
			Name: "migrate",
			Action: func(c *cli.Context) error {
				for _, dbList := range database.DBList() {
					if err := db.Debug().AutoMigrate(dbList.Model); err != nil {
						panic(err)
					}
				}
				os.Exit(0)
				return nil
			},
		},
		{
			Name: "user_mock",
			Action: func(c *cli.Context) error {
				if err := db.Debug().Create(mock.UserMock()).Error; err != nil {
					panic(err)
				}
				os.Exit(0)
				return nil
			},
		},
		{
			Name: "project_mock",
			Action: func(c *cli.Context) error {
				if err := db.Debug().Create(mock.ProjectMock()).Error; err != nil {
					panic(err)
				}
				os.Exit(0)
				return nil
			},
		},
		{
			Name: "bid_mock",
			Action: func(c *cli.Context) error {
				if err := db.Debug().Create(mock.BidMock()).Error; err != nil {
					panic(err)
				}
				os.Exit(0)
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// Creating the server
	// Move to gin-gonic framework
	router := gin.Default()
	router.Use(utils.GinContextToContextMiddleware())
	router.GET("/_gql", playgroundHandler())

	// Version 1 API
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "status ok",
			})
		})
	}

	// GraphQL gqlgen
	gqlRouter := router.Group("/_gql")
	{
		gqlRouter.POST("/query", graphqlHandler())
	}

	// Run on port 9000
	router.Run(":9000")
}
