package token

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenType  = "access_token"
	RefreshTokenType = "refresh_token"
)

type Token struct {
	Algorithm           string
	SigningKey          any
	VerifyingKey        string
	AccessTokenTimeout  time.Duration
	RefreshTokenTimeout time.Duration
}

func (t *Token) GenerateAccessToken(payload map[string]any) (string, error) {
	payload["type"] = AccessTokenType
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * t.AccessTokenTimeout).Unix(),
		"nbf":     time.Now().Unix(),
		"iat":     time.Now().Unix(),
		"jti":     "refresh_token",
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(t.Algorithm), tokenContent)
	accessToken, err := jwtToken.SignedString(t.SigningKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (t *Token) GenerateRefreshToken(payload map[string]any) (string, error) {
	payload["type"] = RefreshTokenType
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(t.AccessTokenTimeout).Unix(),
		"nbf":     time.Now().Unix(),
		"iat":     time.Now().Unix(),
		"jti":     "refresh_token",
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(t.Algorithm), tokenContent)
	return jwtToken.SignedString(t.SigningKey)
}

func (t *Token) VerifyToken(accessToken string) (map[string]any, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return t.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	payload := token.Claims.(jwt.MapClaims)["payload"].(map[string]any)
	if payload["type"] != AccessTokenType {
		return nil, err
	}

	return payload, nil
}

func (t *Token) ParseRefreshToken(refreshToken string) (map[string]any, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return t.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	payload := token.Claims.(jwt.MapClaims)["payload"].(map[string]any)
	if payload["type"] != RefreshTokenType {
		return nil, err
	}

	return payload, nil
}

func (t *Token) VerifyTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the header
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  "INVALID_INPUT",
				"error": "Missing access token",
			})
			c.Abort()
			return
		}

		accessToken = strings.Replace(accessToken, "Bearer ", "", 1)

		// Verify the token
		claims, err := t.VerifyToken(accessToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  "INVALID_INPUT",
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Set the user ID to the context
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

func NewToken(
	algorithm,
	signingKey,
	verifyingKey string,
	accessTokenTimeout,
	refreshTokenTimeout time.Duration) JwtToken {

	var key any
	if strings.Contains(algorithm, "HS") {
		key = []byte(signingKey)
	} else if strings.Contains(algorithm, "RS") {
		key = GetRSAPrivateKey(GetKeyBytes(signingKey))
	} else if strings.Contains(algorithm, "Ed") {
		key = GetEdDSAPrivateKey(GetKeyBytes(signingKey))
	}

	return &Token{
		Algorithm:           algorithm,
		SigningKey:          key,
		VerifyingKey:        verifyingKey,
		AccessTokenTimeout:  accessTokenTimeout,
		RefreshTokenTimeout: refreshTokenTimeout,
	}
}
