package title

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/pkg/app"
	"github.com/amjadjibon/microservices/pkg/db"
	"github.com/amjadjibon/microservices/title/conf"
	"github.com/amjadjibon/microservices/title/handler"
	"github.com/amjadjibon/microservices/title/repo"
)

func Run() {
	cfg := conf.GetConfig()

	// Start the service.
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	postgres, err := db.NewPostgres(cfg.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	defer postgres.Pool.Close()

	repos := repo.NewTitleRepo(postgres)
	router := Router(repos)
	app.Run(addr, router)
}

func Router(repo *repo.TitleRepo) *gin.Engine {
	router := gin.Default()
	handlers := handler.NewTitleHandler(repo)

	// Add routes here.
	// Handlers for CRUD operations on titles
	titleRoutes := router.Group("/titles")
	{
		titleRoutes.GET("/", handlers.GetAllTitle)
		titleRoutes.GET("/:title_id", handlers.GetTitleByID)
		titleRoutes.POST("/", handlers.CreateTitle)
		titleRoutes.PUT("/:title_id", handlers.UpdateTitle) // Adjusted method name
		titleRoutes.DELETE("/:title_id", handlers.DeleteTitle)

		// Handlers for title's content operations (assuming these methods exist in handlers)
		titleRoutes.GET("/:title_id/content", handlers.GetContentForTitle)
		titleRoutes.POST("/:title_id/content", handlers.CreateContentForTitle)
		titleRoutes.PUT("/:title_id/content/:content_id", handlers.UpdateContentForTitle)
		titleRoutes.DELETE("/:title_id/content/:content_id", handlers.DeleteContentForTitle)

		// Search operations (assuming these methods exist in handlers)
		titleRoutes.GET("/search", handlers.SearchTitles)
		titleRoutes.GET("/:title_id/search", handlers.SearchContentForTitle)
	}

	return router
}
