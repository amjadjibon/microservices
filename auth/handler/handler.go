package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/auth/repo"
)

type AuthHandler interface {
	CreateUser(c *gin.Context)
}

type authHandler struct {
	repo repo.AuthRepo
}

func NewAuthHandler(repo repo.AuthRepo) AuthHandler {
	return &authHandler{repo}
}
