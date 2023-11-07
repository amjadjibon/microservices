package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/auth/repo"
	"github.com/amjadjibon/microservices/pkg/google"
	"github.com/amjadjibon/microservices/pkg/token"
)

type AuthHandler interface {
	CreateUser(c *gin.Context)
	CreateRole(c *gin.Context)
	LoginUser(c *gin.Context)
}

type authHandler struct {
	repo         repo.AuthRepo
	jwtToken     token.JwtToken
	oAuth2Client *google.OAuth2Client
}

func NewAuthHandler(
	repo repo.AuthRepo,
	jwtToken token.JwtToken,
	oAuth2Client *google.OAuth2Client,
) AuthHandler {
	return &authHandler{
		repo:         repo,
		jwtToken:     jwtToken,
		oAuth2Client: oAuth2Client,
	}
}
