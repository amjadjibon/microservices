package auth

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/auth/conf"
	"github.com/amjadjibon/microservices/auth/handler"
	"github.com/amjadjibon/microservices/auth/repo"
	"github.com/amjadjibon/microservices/pkg/app"
	"github.com/amjadjibon/microservices/pkg/db"
	"github.com/amjadjibon/microservices/pkg/google"
	"github.com/amjadjibon/microservices/pkg/logger"
	"github.com/amjadjibon/microservices/pkg/token"
)

// var (
// 	googleOauthConfig = oauth2.Config{
// 		ClientID:     "YOUR_CLIENT_ID",
// 		ClientSecret: "YOUR_CLIENT_SECRET",
// 		RedirectURL:  "YOUR_REDIRECT_URI",
// 		Scopes:       []string{"profile", "email"},
// 		Endpoint:     google.Endpoint,
// 	}
// )

func Run() {
	cfg := conf.GetConfig()

	logger.InitLogger(cfg.LogLevel)

	pg, err := db.NewPostgres(cfg.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	defer pg.Pool.Close()

	err = pg.Pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	jwtToken := token.NewToken(
		cfg.JWTAlgorithm,
		cfg.JWTSigningKey,
		cfg.JWTVerifyingKey,
		cfg.JWTAccessTokenTimeout,
		cfg.JWTRefreshTokenTimeout,
	)

	googleOauthConfig := google.NewOAuth2Client(
		cfg.GoogleOAuthClientID,
		cfg.GoogleOAuthClientSecret,
		cfg.GoogleOAuthRedirectURL,
		cfg.GoogleOAuthScopes,
	)

	router := Router(repo.NewAuthRepo(pg), jwtToken, googleOauthConfig)
	app.Run(cfg.Address, router)
}

func getDBConn(cfg conf.Config) *db.Postgres {
	pg, err := db.NewPostgres(cfg.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	defer pg.Pool.Close()

	err = pg.Pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	return pg
}

func Router(repository repo.AuthRepo, jwtToken token.JwtToken, oauth2Config *google.OAuth2Client) *gin.Engine {
	router := gin.Default()
	handlers := handler.NewAuthHandler(repository, jwtToken, oauth2Config)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/user/create", handlers.CreateUser)
	router.POST("/user/login", handlers.LoginUser)
	router.POST("/user/login/refresh", handlers.LoginRefresh)

	// add router group for user to verify token
	userGroup := router.Group("/api/user")
	userGroup.Use(jwtToken.VerifyTokenMiddleware())
	{
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.GET("/", handlers.GetAllUser)
	}

	return router
}
