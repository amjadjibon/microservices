package auth

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/auth/handler"
	"github.com/amjadjibon/microservices/auth/repo"
	"github.com/amjadjibon/microservices/pkg/app"
	"github.com/amjadjibon/microservices/pkg/db"
	"github.com/amjadjibon/microservices/pkg/logger"
)

func Run() {
	logger.InitLogger("info")
	pg, err := db.NewPostgres("postgres://rootuser:rootpassword@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pg.Pool.Close()

	err = pg.Pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	addr := ":8080"

	router := Router(repo.NewAuthRepo(pg))
	app.Run(addr, router)
}

func Router(repository repo.AuthRepo) *gin.Engine {
	router := gin.Default()
	handlers := handler.NewAuthHandler(repository)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	router.POST("/user/create", handlers.CreateUser)
	return router
}
