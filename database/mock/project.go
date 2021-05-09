package mock

import (
	"log"
	"math/rand"
	"time"

	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func ProjectMock() model.Project {
	db := config.DbConn()
	var user model.User
	err := db.First(&user).Error
	if err != nil {
		log.Fatal(err)
		return model.Project{}
	}

	return model.Project{
		ID:        uuid.New().String(),
		CreatorID: user.ID,
		Name:      faker.Name(),
		Desc:      faker.Paragraph(),
		Price:     uint(rand.Int()),
		Deadline:  time.Now().Add(time.Duration(time.Now().Day() + 1)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
