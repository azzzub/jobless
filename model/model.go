package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ContextKey struct {
	Name string
}
