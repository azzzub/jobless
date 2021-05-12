package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/database"
	"github.com/azzzub/jobless/database/mock"
	"github.com/azzzub/jobless/graph/generated"
	"github.com/azzzub/jobless/graph/resolvers"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/segmentio/ksuid"
	"github.com/urfave/cli/v2"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	// h := handler.NewDefaultServer()
	h := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &resolvers.Resolver{}}),
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
			Name: "drop_all_tables",
			Action: func(c *cli.Context) error {
				err := db.Exec("DROP TABLE IF EXISTS bids").Error
				if err != nil {
					panic(err)
				}
				err = db.Exec("DROP TABLE IF EXISTS projects").Error
				if err != nil {
					panic(err)
				}
				err = db.Exec("DROP TABLE IF EXISTS users").Error
				if err != nil {
					panic(err)
				}
				os.Exit(0)
				return nil
			},
		},
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
	router.Use(utils.CORSMiddleware())
	router.Use(utils.GinContextToContextMiddleware())
	router.GET("/_gql", playgroundHandler())

	// Version 1 API
	v1 := router.Group("/v1")
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_AUTH_KEY"), os.Getenv("GOOGLE_AUTH_SECRET"), "http://localhost:9000/v1/auth/google/callback"),
	)
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "status ok",
			})
		})

		v1.GET("/auth/google", func(c *gin.Context) {
			q := c.Request.URL.Query()
			q.Add("provider", "google")
			c.Request.URL.RawQuery = q.Encode()
			gothic.BeginAuthHandler(c.Writer, c.Request)
		})

		v1.GET("/auth/google/callback", func(c *gin.Context) {
			user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
				return
			}

			checkEmail := model.User{}

			err = db.Where("email = ? AND provider = ?", user.Email, "google").
				First(&checkEmail).Error
			if err != nil && err.Error() == "record not found" {
				// Make a random username by using ksuid package
				newUser := model.User{
					Username:  ksuid.New().String(),
					Email:     user.Email,
					Provider:  "google",
					Avatar:    user.AvatarURL,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				// Register the user if the email with google provider doesn't exist
				err = db.Create(&newUser).Error
				if err != nil {
					c.JSON(http.StatusBadGateway, gin.H{
						"error": err.Error(),
					})
					return
				}
				err = db.Where("email = ? AND provider = ?", user.Email, "google").
					First(&checkEmail).Error
			}
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
				return
			}

			claims := &model.Token{
				ID:    checkEmail.ID,
				Email: checkEmail.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signedToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

			c.JSON(http.StatusOK, gin.H{
				"data": map[string]map[string]string{
					"login": {
						"token": signedToken,
					},
				},
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
