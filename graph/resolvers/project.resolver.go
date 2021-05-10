package resolvers

import (
	"context"
	"fmt"
	"time"

	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/graph/model"
	"github.com/azzzub/jobless/utils"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := utils.AuthMiddlewareProc(gc)
	if err != nil {
		return nil, err
	}

	db := config.DbConn()

	// Check whether the user has created the project 4 times within a month.
	// Count from the last project's created date minus a month interval.
	var lastProjectCreated model.Project
	if err := db.Last(&lastProjectCreated).Error; err != nil {
		return nil, err
	}

	lastProjectCreatedDate, err := time.Parse(time.RFC3339, lastProjectCreated.CreatedAt)
	if err != nil {
		return nil, err
	}

	aMonthInterval := lastProjectCreatedDate.AddDate(0, -1, 0)

	var aMonthIntervalData []model.Project
	err = db.Where("created_at >= ? AND created_at <= ?", aMonthInterval, lastProjectCreatedDate).
		Order("ID ASC").
		Find(&aMonthIntervalData).Error
	if err != nil {
		return nil, err
	}

	nextMonth, err := time.Parse(time.RFC3339, aMonthIntervalData[0].CreatedAt)
	if err != nil {
		return nil, err
	}

	if len(aMonthIntervalData) >= 4 {
		return nil, fmt.Errorf(
			"you are already created 4 projects within a month, kindly please wait until %v to create a new project",
			nextMonth.AddDate(0, 1, 0),
		)
	}
	// End of counter project checking

	project := &model.Project{
		CreatorID: int(token.ID),
		Name:      input.Name,
		Desc:      input.Desc,
		Price:     input.Price,
		Deadline:  input.Deadline,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	result := db.Select("CreatorID", "Name", "Desc", "Price", "Deadline", "CreatedAt", "UpdatedAt").
		Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project

	db := config.DbConn()
	result := db.Debug().Preload("Creator").Preload("Bids").Preload("Bids.Bidder").Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}
