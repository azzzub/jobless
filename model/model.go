package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// String pointer for gorm unique
type Auth struct {
	gorm.Model
	Username *string `json:"username" gorm:"unique"`
	Email    *string `json:"email" gorm:"unique"`
	Password string  `json:"password"`
}

type Project struct {
	gorm.Model
	CreatorId uint      `json:"creator_id"`
	Name      string    `json:"name" gorm:"unique"`
	Desc      string    `json:"desc"`
	Price     uint      `json:"price"`
	Deadline  time.Time `json:"deadline"`
}

type Token struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ContextKey struct {
	Name string
}
