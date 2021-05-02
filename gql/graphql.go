package gql

import (
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func GraphQL() gin.HandlerFunc {
	db := config.DbConn()

	projectType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Project",
		Description: "Project data.",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.Int,
				Description: "The identifier of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.Name, nil
					}

					return nil, nil
				},
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"projects": &graphql.Field{
				Type:        graphql.NewList(projectType),
				Description: "List of projects",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var projects []model.Project
					result := db.Find(&projects)
					if result.Error != nil {
						return nil, result.Error
					}

					return projects, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

	// h := gin.HandlerFunc{
	// 	&handler.Config{
	// 		Schema:   &schema,
	// 		Pretty:   true,
	// 		GraphiQL: true,
	// 	},
	// }

	// return h
	// http.Handle("/", h)
	// http.ListenAndServe(":9000", nil)
}
