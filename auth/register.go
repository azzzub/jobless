package auth

import (
	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *fiber.Ctx) error {
	password := "thisisthepassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		panic(err)
	}

	auth := model.Auth{
		Username: "username",
		Email:    "username@email.com",
		Password: string(hashedPassword),
	}

	db := config.DbConn()
	db.Create(&auth)

	finalResponse := fiber.Map{
		"message": "Success add new user!",
		"data":    &auth,
	}

	return c.JSON(finalResponse)
}
