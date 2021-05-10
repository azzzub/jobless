package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Username        string         `json:"username"`
	Email           string         `json:"email"`
	IsEmailVerified bool           `json:"is_email_verified" gorm:"default:false"`
	IsUserVerified  bool           `json:"is_user_verified" gorm:"default:false"`
	Password        string         `json:"password"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Projects        []Project      `gorm:"foreignKey:CreatorID;references:ID"`
	Bids            []Bid          `gorm:"foreignKey:BidderID;references:ID"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
