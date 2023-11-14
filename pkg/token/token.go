package token

type JwtToken interface {
	GenerateAccessToken(payload map[string]any) (string, error)
	GenerateRefreshToken(payload map[string]any) (string, error)
	ParseRefreshToken(refreshToken string) (map[string]any, error)
}
