package mock

import (
	"time"

	"github.com/azzzub/jobless/model"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func UserMock() model.User {
	return model.User{
		// Password: 123
		ID:        uuid.New().String(),
		Username:  faker.Username(),
		Email:     faker.Email(),
		Password:  "$2y$10$aaCw9WsT2aTR/9QE/SvpX./hol9R3BmA9Fh0Cl8o0rAuhiOhTaFJi",
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

}
