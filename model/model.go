package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// String pointer for gorm unique
type User struct {
	ID        string         `json:"id" gorm:"uniqueIndex"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Projects  []Project      `gorm:"foreignKey:CreatorID;references:ID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Bids      []Bid          `gorm:"foreignKey:BidderID"`

type Project struct {
	ID        string `json:"id" gorm:"uniqueIndex"`
	CreatorID string `json:"creator_id"`
	Creator   User
	// Bids      []Bid `gorm:"foreignKey:ProjectID"`
	// Bid       *Bid
	Name      string         `json:"name" gorm:"unique"`
	Desc      string         `json:"desc"`
	Price     uint           `json:"price"`
	Deadline  time.Time      `json:"deadline"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Bid struct {
	ID string `json:"id" gorm:"uniqueIndex"`
	// BidderID  string `json:"bidder_id"`
	// Bidder    *User
	// ProjectID string `json:"project_id"`
	// Project   *Project
	Price     uint           `json:"price"`
	Comment   string         `json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Token struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type ContextKey struct {
	Name string
}
