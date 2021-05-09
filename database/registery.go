package database

import (
	"github.com/azzzub/jobless/model"
)

type dbList struct {
	Model interface{}
}

func DBList() []dbList {
	return []dbList{
		{Model: model.User{}},
		{Model: model.Project{}},
		{Model: model.Bid{}},
	}
}
