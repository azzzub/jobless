package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	database "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type loginField struct {
	Uoe      string `validate:"required,min=4"`
	Password string `validate:"required,min=8"`
}

func LoginHandler(c *fiber.Ctx) error {
	body := new(loginField)

	if err := c.BodyParser(body); err != nil {
		return utils.ErrorHandler(c, http.StatusBadRequest, err)
	}

	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		return utils.ErrorHandler(c, http.StatusBadRequest, err)
	}

	db := database.DbConn()

	auth := new(model.Auth)

	result := db.Where("username = ?", body.Uoe).Or("email = ?", body.Uoe).First(&auth)

	if result.Error != nil {
		return utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("wrong email/username"))
	}

	err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(body.Password))

	if err != nil {
		return utils.ErrorHandler(c, http.StatusUnauthorized, errors.New("wrong password"))
	}

	claims := model.Token{
		ID: auth.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return c.JSON(fiber.Map{
		"token": signedToken,
	})
}
