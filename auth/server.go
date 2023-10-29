package auth

import (
	"github.com/amjadjibon/microservices/pkg/app"
	"github.com/amjadjibon/microservices/pkg/db"
	"github.com/amjadjibon/microservices/pkg/logger"
)

func Run() {
	logger.InitLogger("info")
	pg, err := db.NewPostgres("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pg.Pool.Close()
	addr := ":8080"
	router := Router()
	app.Run(addr, router)
}
