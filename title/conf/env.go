package conf

import (
	"os"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Host        string `env:"HOST" envDefault:"0.0.0.0"`
	Port        int    `env:"PORT" envDefault:"8080"`
	Address     string `env:"ADDRESS,expand" envDefault:"$HOST:${PORT}"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	DatabaseDSN string `env:"DATABASE_DSN"`

	// JWT
	JWTAlgorithm           string        `env:"JWT_ALGORITHM" envDefault:"RS256"`
	JWTSigningKey          string        `env:"JWT_SIGNING_KEY"`
	JWTVerifyingKey        string        `env:"JWT_VERIFYING_KEY"`
	JWTAccessTokenTimeout  time.Duration `env:"JWT_ACCESS_TOKEN_TIMEOUT"`
	JWTRefreshTokenTimeout time.Duration `env:"JWT_REFRESH_TOKEN_TIMEOUT"`
}

func GetConfig() Config {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load("title/conf/.env")
		if err != nil {
			panic(err)
		}
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	log.Debug().Any("config", cfg).Msg("get config success")
	return cfg
}
