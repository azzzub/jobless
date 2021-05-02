package auth

import (
	"net/http"

	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type registerField struct {
	Username string `validate:"required,min=4,max=32"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func validateRegister(body registerField) error {
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err
	}

	return nil
}

func RegisterHandler(c *gin.Context) {
	db := config.DbConn()
	var body registerField

	if err := c.BindJSON(&body); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	if err := validateRegister(body); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		utils.ErrorHandler(c, http.StatusUnauthorized, err)
		return
	}

	auth := model.Auth{
		Username: &body.Username,
		Email:    &body.Email,
		Password: string(hashedPassword),
	}

	result := db.Create(&auth)
	if result.Error != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success add new user",
		"data":    auth,
	})
}
