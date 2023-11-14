package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getUserOutput struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func (a *authHandler) GetUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_INPUT",
			"error": err.Error(),
		})
		return
	}

	// Call the GetUser method from the repository
	user, err := a.repo.GetUserById(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_INPUT",
			"error": err.Error(),
		})
		return
	}

	// Return the user
	c.JSON(http.StatusOK, gin.H{
		"code": "SUCCESS",
		"data": getUserOutput{
			ID:         user.ID,
			Username:   user.Username,
			Name:       user.Name,
			Email:      user.Email,
			IsVerified: user.IsVerified,
		},
	})
}

func (a *authHandler) GetAllUser(c *gin.Context) {
	// check context if user is admin
	role, ok := c.Get("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_INPUT",
			"error": "Missing user role",
		})
		return
	}

	fmt.Println(role)

	if role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_INPUT",
			"error": "You are not authorized to access this resource",
		})
		return
	}

	// Call the GetUser method from the repository
	user, err := a.repo.GetAllUser(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "FAILED_GET_ALL_USER",
			"error": err.Error(),
		})
		return
	}

	var users []getUserOutput
	for _, u := range user {
		users = append(users, getUserOutput{
			ID:         u.ID,
			Username:   u.Username,
			Name:       u.Name,
			Email:      u.Email,
			IsVerified: u.IsVerified,
		})
	}

	// Return the user
	c.JSON(http.StatusOK, gin.H{
		"code": "SUCCESS",
		"data": users,
	})
}
