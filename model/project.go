package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Slug      string `json:"slug" gorm:"unique"`
	CreatorID uint   `json:"creator_id"`
	Creator   User
	Bids      []Bid `gorm:"foreignKey:ProjectID"`
	Bid       *Bid
	Name      string         `json:"name"`
	Desc      string         `json:"desc"`
	Price     uint           `json:"price"`
	Deadline  time.Time      `json:"deadline"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
