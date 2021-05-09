package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"os"
	"time"

	db "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/graph/generated"
	"github.com/azzzub/jobless/graph/model"
	rawModel "github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	jwt "github.com/dgrijalva/jwt-go"
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

	db := db.DbConn()

	result := db.Select("Username", "Email", "Password", "CreatedAt", "UpdatedAt").
		Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.LoginResponse, error) {
	user := model.User{}
	loginResponse := &model.LoginResponse{}

	db := db.DbConn()

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
		ID: uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	loginResponse.Token = signedToken

	return loginResponse, nil
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := utils.AuthMiddlewareProc(gc)
	if err != nil {
		return nil, err
	}

	project := &model.Project{
		CreatorID: int(token.ID),
		Name:      input.Name,
		Desc:      input.Desc,
		Price:     input.Price,
		Deadline:  input.Deadline,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	db := db.DbConn()
	result := db.Select("CreatorID", "Name", "Desc", "Price", "Deadline", "CreatedAt", "UpdatedAt").
		Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}

func (r *mutationResolver) CreateBid(ctx context.Context, input model.NewBid) (*model.Bid, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := utils.AuthMiddlewareProc(gc)
	if err != nil {
		return nil, err
	}

	bid := &model.Bid{
		BidderID:  int(token.ID),
		ProjectID: input.ProjectID,
		Price:     input.Price,
		Comment:   input.Comment,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	db := db.DbConn()
	result := db.Select("BidderID", "ProjectID", "Price", "Comment", "CreatedAt", "UpdatedAt").
		Create(&bid)

	if result.Error != nil {
		return nil, result.Error
	}

	return bid, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project

	db := db.DbConn()
	result := db.Debug().Preload("Creator").Preload("Bids").Preload("Bids.Bidder").Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (r *queryResolver) Bids(ctx context.Context) ([]*model.Bid, error) {
	var bids []*model.Bid

	db := db.DbConn()
	result := db.Find(&bids)
	if result.Error != nil {
		return nil, result.Error
	}

	return bids, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
