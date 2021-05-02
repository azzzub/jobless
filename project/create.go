package project

import (
	"errors"
	"net/http"
	"time"

	config "github.com/azzzub/jobless/config"
	"github.com/azzzub/jobless/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"strings"

	"github.com/azzzub/jobless/utils"
)

type projectField struct {
	Name     string    `json:"name" validate:"required"`
	Desc     string    `json:"desc" validate:"required"`
	Price    uint      `json:"price" validate:"required,number,gte=0"`
	Deadline time.Time `json:"deadline" validate:"required"`
}

func CreateProject(c *gin.Context) {
	db := config.DbConn()

	// Get the bearer auth header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ErrorHandler(c, http.StatusUnauthorized,
			errors.New("bearer auth must be provided"))
		return
	}

	token := strings.Split(authHeader, " ")
	decodedToken, err := utils.TokenValidator(token[1])
	if err != nil {
		utils.ErrorHandler(c, http.StatusUnauthorized, err)
		return
	}

	// Check whether the creator id is exist
	var auth model.Auth
	checkCreatorId := db.First(&auth).Where("ID = ?", decodedToken.ID)
	if checkCreatorId.Error != nil {
		utils.ErrorHandler(c, http.StatusBadRequest,
			errors.New("cannot find the creator, recheck your token"))
		return
	}

	var body projectField
	if err := c.BindJSON(&body); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	project := model.Project{
		CreatorId: decodedToken.ID,
		Name:      strings.ToLower(body.Name),
		Desc:      body.Desc,
		Price:     body.Price,
		Deadline:  body.Deadline,
	}

	result := db.Create(&project)
	if result.Error != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success add new project",
		"data":    project,
	})
}
