package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/amjadjibon/microservices/pkg/password"
)

var (
	googleOauthConfig = oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "YOUR_REDIRECT_URI",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
)

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (a *authHandler) LoginUser(c *gin.Context) {
	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"code":  "INVALID_INPUT",
			"error": err.Error(),
		})
		return
	}

	// Call the GetUserByUsername method from the repository
	user, err := a.repo.GetUserByUsername(c.Request.Context(), input.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  "GET_USER_FAILED",
			"error": err.Error()},
		)
		return
	}

	// Compare the given password with the password from the database
	if !password.VerifyPassword(user.Password, input.Password) {
		c.JSON(400, gin.H{
			"code":  "INVALID_PASSWORD",
			"error": "invalid password"},
		)
		return
	}

	// Generate access token
	payload := map[string]any{
		"id":   user.ID,
		"role": user.Role,
	}

	accessToken, err := a.jwtToken.GenerateAccessToken(payload)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  "GENERATE_ACCESS_TOKEN_FAILED",
			"error": err.Error(),
		})
		return
	}

	refreshToken, err := a.jwtToken.GenerateRefreshToken(payload)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  "GENERATE_REFRESH_TOKEN_FAILED",
			"error": err.Error(),
		})
	}

	output := loginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(200, gin.H{
		"code": "SUCCESS",
		"data": output,
	})
}

func (a *authHandler) GoogleCallbackHandler(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, "Missing code")
		return
	}

	token, err := a.oAuth2Client.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Use the token to fetch user information from Google
	userInfo, err := a.oAuth2Client.FetchUserInfo(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(userInfo)

	// Register the user if not already registered, or log in the user.

	// Redirect or respond to the client as needed.
}
