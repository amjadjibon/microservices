package token

import (
	"strings"
	"time"

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
