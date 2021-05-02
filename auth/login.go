package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type loginField struct {
	Uoe      string `validate:"required,min=4"`
	Password string `validate:"required,min=8"`
}

func validateLogin(body loginField) error {
	validation := validator.New()
	if err := validation.Struct(body); err != nil {
		return err
	}

	return nil
}

func LoginHandler(c *gin.Context) {
	var body loginField
	db := config.DbConn()

	if err := c.BindJSON(&body); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	if err := validateLogin(body); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	var auth model.Auth

	result := db.Where("username = ?", body.Uoe).Or("email = ?", body.Uoe).First(&auth)
	if result.Error != nil {
		utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("wrong email/username"))
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(auth.Password),
		[]byte(body.Password)); err != nil {
		utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("wrong password"))
		return
	}

	claims := model.Token{
		ID: auth.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(http.StatusOK, gin.H{
		"token": signedToken,
	})
}
