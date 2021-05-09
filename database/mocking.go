package database

import (
	"github.com/azzzub/jobless/database/mock"
)

type Mocks struct {
	Mock interface{}
}

func NewMock() []Mocks {
	return []Mocks{
		{
			Mock: mock.ProjectMock(),
		},
		{
			Mock: mock.BidMock(),
		},
	}
}
