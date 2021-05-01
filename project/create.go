package project

import (
	"errors"
	"time"

	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/go-playground/validator"

	"strings"

	"github.com/azzzub/jobless/utils"
	"github.com/gofiber/fiber/v2"
)

type projectField struct {
	Name     string    `json:"name" validate:"required"`
	Desc     string    `json:"desc" validate:"required"`
	Price    uint      `json:"price" validate:"required,number,gte=0"`
	Deadline time.Time `json:"deadline" validate:"required"`
}

func CreateProject(c *fiber.Ctx) error {
	db := config.DbConn()

	// Get the bearer auth header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.ErrorHandler(c, fiber.StatusUnauthorized, errors.New("bearer auth must be provided"))
	}

	token := strings.Split(authHeader, " ")
	decodedToken, err := utils.TokenValidator(token[1])
	if err != nil {
		return utils.ErrorHandler(c, fiber.StatusUnauthorized, err)
	}

	// Check whether the creator id is exist
	var auth model.Auth
	checkCreatorId := db.First(&auth).Where("ID = ?", decodedToken.ID)
	if checkCreatorId.Error != nil {
		return utils.ErrorHandler(c, fiber.StatusBadRequest,
			errors.New("cannot find the creator, recheck your token"))
	}

	var body projectField
	if err := c.BodyParser(&body); err != nil {
		return utils.ErrorHandler(c, fiber.StatusBadRequest, err)
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		return utils.ErrorHandler(c, fiber.StatusBadRequest, err)
	}

	project := model.Project{
		CreatorId: decodedToken.ID,
		Name:      body.Name,
		Desc:      body.Desc,
		Price:     body.Price,
		Deadline:  body.Deadline,
	}

	result := db.Create(&project)
	if result.Error != nil {
		return utils.ErrorHandler(c, fiber.StatusBadGateway, result.Error)
	}

	return c.JSON(fiber.Map{
		"data":    project,
		"message": "success add new project",
	})
}
