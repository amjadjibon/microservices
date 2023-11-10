package handler

import (
	error_parser "github.com/amjadjibon/microservices/auth/error"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/auth/model"
)

type createRoleInput struct {
	Name string `json:"name" binding:"required"`
}

type createRoleOutput struct {
	ID int `json:"id"`
}

func (a *authHandler) CreateRole(c *gin.Context) {
	var input createRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   "INVALID_INPUT",
			"errors": error_parser.ParseError(err),
		})
		return
	}

	// Create a new role
	role := &model.Role{
		Name:      input.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// Call the CreateRole method from the repository
	role, err := a.repo.CreateRole(c.Request.Context(), role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "CREATE_ROLE_FAILED",
			"error": err.Error()},
		)
		return
	}

	var output = createRoleOutput{
		ID: role.ID,
	}

	c.JSON(http.StatusCreated, output)
}
