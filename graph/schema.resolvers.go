package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	db "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/graph/generated"
	"github.com/azzzub/jobless/graph/model"
	rawModel "github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
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
	result := db.Select("CreatorID", "Name", "Desc", "Price", "Deadline", "CreatedAt", "UpdatedAt").
		Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}

func (r *mutationResolver) CreateBid(ctx context.Context, input model.NewBid) (*model.Bid, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	userId, exist := gc.Get("UserID")
	if !exist {
		return nil, errors.New("header not exist")
	}

	id, ok := userId.(*rawModel.Token)
	if !ok {
		return nil, errors.New("type convert error")
	}

	bid := &model.Bid{
		BidderID:  int(id.ID),
		ProjectID: input.ProjectID,
		Price:     input.Price,
		Comment:   input.Comment,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	db := db.DbConn()
	result := db.Select("BidderID", "ProjectID", "Price", "Comment", "CreatedAt", "UpdatedAt").
		Create(&bid)

	if result.Error != nil {
		return nil, result.Error
	}

	return bid, nil
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
}

func (r *queryResolver) Bids(ctx context.Context) ([]*model.Bid, error) {
	var bids []*model.Bid

	db := db.DbConn()
	result := db.Find(&bids)
	if result.Error != nil {
		return nil, result.Error
	}

	return bids, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
