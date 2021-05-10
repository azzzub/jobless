package resolvers

import (
	"context"
	"errors"
	"time"

	"github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/graph/model"
	"github.com/azzzub/jobless/utils"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	db := config.DbConn()

	result := db.Select("Username", "Email", "Password", "CreatedAt", "UpdatedAt").
		Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	err = utils.SendMail(user.Email, "Email Confirmation", "Email confirmation")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) EmailVerification(ctx context.Context, input model.EmailVerification) (*model.EmailVerificationResponse, error) {
	db := config.DbConn()

	token, err := utils.TokenValidator(input.Token)
	if err != nil {
		return nil, err
	}

	var user model.User

	err = db.Where("email = ?", token.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.IsEmailVerified {
		return nil, errors.New("you've already verified your email")
	}

	updateUser := model.User{
		IsEmailVerified: true,
		UpdatedAt:       time.Now().Format(time.RFC3339),
	}

	err = db.Debug().Model(&user).Updates(&updateUser).Error
	if err != nil {
		return nil, err
	}

	return &model.EmailVerificationResponse{
		Message: "your email successfuly verified",
	}, nil
}
