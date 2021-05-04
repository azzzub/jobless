package gql

import (
	"time"

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
			"creator_id": &graphql.Field{
				Type:        graphql.Int,
				Description: "The creator's identifier of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.CreatorId, nil
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
			"desc": &graphql.Field{
				Type:        graphql.String,
				Description: "The description of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.Desc, nil
					}

					return nil, nil
				},
			},
			"price": &graphql.Field{
				Type:        graphql.Int,
				Description: "The price of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.Price, nil
					}

					return nil, nil
				},
			},
			"deadline": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "The deadline of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.Deadline, nil
					}

					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "The created date of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.CreatedAt, nil
					}

					return nil, nil
				},
			},
			"updated_at": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "The modified date of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.UpdatedAt, nil
					}

					return nil, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "The delete date of the project.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if project, ok := p.Source.(model.Project); ok {
						return project.DeletedAt.Time, nil
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

	mutationQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "MutationQuery",
		Fields: graphql.Fields{
			"createProject": &graphql.Field{
				Type:        projectType,
				Description: "Create a new project",
				Args: graphql.FieldConfigArgument{
					"creator_id": &graphql.ArgumentConfig{
						Description: "The project creator identifier number",
						Type:        graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Description: "The name of the project",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"desc": &graphql.ArgumentConfig{
						Description: "The description of the project",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Description: "The price of the project",
						Type:        graphql.NewNonNull(graphql.Int),
					},
					"deadline": &graphql.ArgumentConfig{
						Description: "The deadline of the project",
						Type:        graphql.NewNonNull(graphql.DateTime),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					creatorId, _ := p.Args["creator_id"].(int)
					name, _ := p.Args["name"].(string)
					desc, _ := p.Args["desc"].(string)
					price, _ := p.Args["price"].(int)
					deadline, _ := p.Args["deadline"].(time.Time)

					project := model.Project{
						CreatorId: uint(creatorId),
						Name:      name,
						Desc:      desc,
						Price:     uint(price),
						Deadline:  deadline,
					}

					result := db.Create(&project)

					if result.Error != nil {
						return nil, result.Error
					}

					return project, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: mutationQuery,
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
