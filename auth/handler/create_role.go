package handler

import (
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
		c.JSON(400, gin.H{
			"code":  "INVALID_INPUT",
			"error": err.Error(),
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
		c.JSON(400, gin.H{
			"code":  "CREATE_ROLE_FAILED",
			"error": err.Error()},
		)
		return
	}

	var output = createRoleOutput{
		ID: role.ID,
	}

	c.JSON(200, output)
}
