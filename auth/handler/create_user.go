package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/auth/model"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	// Other fields specific to the creation request
}

type CreateUserOutput struct {
	ID int `json:"id"`
}

func (a *authHandler) CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Create a new user
	user := &model.User{
		Username: input.Username,
		Name:     input.Name,
		Email:    input.Email,
		// Other fields specific to the creation request
	}

	// Call the CreateUser method from the repository
	user, err := a.repo.CreateUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var output = CreateUserOutput{
		ID: user.ID,
	}

	c.JSON(http.StatusOK, output)
}
