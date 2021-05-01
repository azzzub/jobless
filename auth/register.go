package auth

import (
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/azzzub/jobless/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type registerField struct {
	Username string `validate:"required,min=4,max=32"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func validateStruct(body registerField) error {
	validate := validator.New()

	err := validate.Struct(body)

	if err != nil {
		return err
	}

	return nil
}

func RegisterHandler(c *fiber.Ctx) error {
	body := new(registerField)

	if err := c.BodyParser(body); err != nil {
		return utils.ErrorHandler(c, 500, err)
	}

	if err := validateStruct(*body); err != nil {
		return utils.ErrorHandler(c, 500, err)
	}

	username := body.Username
	email := body.Email
	password := body.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return utils.ErrorHandler(c, 500, err)
	}

	auth := model.Auth{
		Username: &username,
		Email:    &email,
		Password: string(hashedPassword),
	}

	db := config.DbConn()
	result := db.Create(&auth)

	if result.Error != nil {
		return utils.ErrorHandler(c, 500, result.Error)
	}

	finalResponse := fiber.Map{
		"message": "Success add new user!",
		"data":    &auth,
	}

	return c.JSON(finalResponse)
}
