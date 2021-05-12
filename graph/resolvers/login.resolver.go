package resolvers

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/graph/model"
	rawModel "github.com/azzzub/jobless/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.LoginResponse, error) {
	user := model.User{}
	loginResponse := &model.LoginResponse{}

	db := config.DbConn()

	result := db.Where("username = ?", input.Uoe).Or("email = ?", input.Uoe).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password)); err != nil {
		return nil, errors.New("wrong login information")
	}

	claims := &rawModel.Token{
		ID:    uint(user.ID),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
		},
	}

	refreshClaims := &rawModel.Token{
		ID:    uint(user.ID),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, _ := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET_REFRESH")))

	loginResponse.Token = signedToken
	loginResponse.RefreshToken = signedRefreshToken

	return loginResponse, nil
}
