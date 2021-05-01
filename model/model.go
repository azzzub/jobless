package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// String pointer for gorm unique
type Auth struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  *string   `json:"username" gorm:"unique"`
	Email     *string   `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}
