package mock

import (
	"math/rand"
	"time"

	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/bxcodec/faker/v3"
)

func BidMock() *model.Bid {
	var user model.User
	db := config.DbConn()
	err := db.First(&user).Error
	if err != nil {
		return &model.Bid{}
	}

	var project model.Project

	err = db.First(&project).Error
	if err != nil {
		return &model.Bid{}
	}

	return &model.Bid{
		BidderID:  user.ID,
		ProjectID: project.ID,
		Price:     uint(rand.Int()),
		Comment:   faker.Paragraph(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

}
