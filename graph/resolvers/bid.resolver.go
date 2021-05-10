package resolvers

import (
	"context"
	"errors"
	"time"

	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/graph/model"
	"github.com/azzzub/jobless/utils"
)

func (r *mutationResolver) CreateBid(ctx context.Context, input model.NewBid) (*model.Bid, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := utils.AuthMiddlewareProc(gc)
	if err != nil {
		return nil, err
	}

	db := config.DbConn()
	var bidProjectCounter int64
	db.Table("bids").Where("bidder_id = ? AND project_id = ?", token.ID, input.ProjectID).Count(&bidProjectCounter)

	if bidProjectCounter >= 2 {
		return nil, errors.New("you already made 2 bids on this project")
	}

	bid := &model.Bid{
		BidderID:  int(token.ID),
		ProjectID: input.ProjectID,
		Price:     input.Price,
		Comment:   input.Comment,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	result := db.Select("BidderID", "ProjectID", "Price", "Comment", "CreatedAt", "UpdatedAt").
		Create(&bid)

	if result.Error != nil {
		return nil, result.Error
	}

	return bid, nil
}

func (r *queryResolver) Bids(ctx context.Context) ([]*model.Bid, error) {
	var bids []*model.Bid

	db := config.DbConn()
	result := db.Preload("Project").Preload("Bidder").Find(&bids)
	if result.Error != nil {
		return nil, result.Error
	}

	return bids, nil
}
