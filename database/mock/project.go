package mock

import (
	"log"
	"math/rand"
	"time"

	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/bxcodec/faker/v3"
	"github.com/gosimple/slug"
)

func ProjectMock() *model.Project {
	db := config.DbConn()
	var user model.User
	err := db.First(&user).Error
	if err != nil {
		log.Fatal(err)
		return &model.Project{}
	}

	name := faker.Name()

	return &model.Project{
		CreatorID: user.ID,
		Slug:      slug.Make(name),
		Name:      name,
		Desc:      faker.Paragraph(),
		Price:     uint(rand.Int()),
		Deadline:  time.Now().Add(time.Duration(time.Now().Day() + 1)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
