package utils

import (
	"errors"
	"os"

	"github.com/azzzub/jobless/model"
	"github.com/dgrijalva/jwt-go"
)

func TokenValidator(tokenString string) (*model.Token, error) {
	claims := &model.Token{}

	result, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !result.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
