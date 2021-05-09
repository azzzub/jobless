package model

import (
	"time"

	"gorm.io/gorm"
)

type Bid struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	BidderID  uint `json:"bidder_id"`
	Bidder    *User
	ProjectID uint `json:"project_id"`
	Project   *Project
	Price     uint           `json:"price"`
	Comment   string         `json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
