package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/auth/model"
	"github.com/amjadjibon/microservices/pkg/password"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserOutput struct {
	ID int `json:"id"`
}

func (a *authHandler) CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_INPUT",
			"error": err.Error(),
		})
		return
	}

	pass, err := password.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_INPUT",
			"error": err.Error(),
		})
	}

	// Create a new user
	user := &model.User{
		Username:   input.Username,
		Name:       input.Name,
		Email:      input.Email,
		IsVerified: false,
		Gender:     input.Gender,
		Role:       input.Role,
		Password:   pass,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	// Call the CreateUser method from the repository
	user, err = a.repo.CreateUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "CREATE_USER_FAILED",
			"error": err.Error()},
		)
		return
	}

	var output = CreateUserOutput{
		ID: user.ID,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "SUCCESS",
		"message": "User created successfully",
		"data":    output,
	})
}
