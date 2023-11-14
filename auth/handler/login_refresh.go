package handler

import (
	"github.com/gin-gonic/gin"
)

type loginRefreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (a *authHandler) LoginRefresh(c *gin.Context) {
	var input loginRefreshInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"code":  "INVALID_INPUT",
			"error": err.Error(),
		})
		return
	}

	// Call the ParseRefreshToken method from the jwtToken
	payload, err := a.jwtToken.ParseRefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  "PARSE_REFRESH_TOKEN_FAILED",
			"error": err.Error(),
		})
		return
	}

	// Generate access token
	accessToken, err := a.jwtToken.GenerateAccessToken(payload)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  "GENERATE_ACCESS_TOKEN_FAILED",
			"error": err.Error(),
		})
		return
	}

	// Generate refresh token
	refreshToken, err := a.jwtToken.GenerateRefreshToken(payload)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  "GENERATE_REFRESH_TOKEN_FAILED",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
