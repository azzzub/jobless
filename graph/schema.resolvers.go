package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	db "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/graph/generated"
	"github.com/azzzub/jobless/graph/model"

	"github.com/icza/gox/fmtx"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	project := &model.Project{
		CreatorID: input.CreatorID,
		Name:      input.Name,
		Desc:      input.Desc,
		Price:     input.Price,
		Deadline:  input.Deadline,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	db := db.DbConn()
	result := db.Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil

	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project

	db := db.DbConn()
	result := db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, project := range projects {
		priceString := "Rp" + fmtx.FormatInt(int64(project.Price), 3, '.') + ",-"
		project.PriceString = &priceString
	}

	return projects, nil
	// panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
