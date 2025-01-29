package main

import (
	"learning/config"
	"learning/internal/service"
	"log"

	"learning/internal/handler"
	"learning/internal/repository/datastore"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Initialize Datastore client
	datastoreClient := datastore.NewDatastoreClient(cfg)

	// Initialize repositories
	playerRepo := datastore.NewDatastorePlayerRepository(datastoreClient)

	// Initialize services, controllers, handlers...
	playerService := service.NewPlayerService(playerRepo)
	playerHandler := handler.NewPlayerHandler(playerService)

	// Setup router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	api := router.Group("/api/v1")
	{
		players := api.Group("/players")
		{
			players.POST("/", playerHandler.Create)
			players.GET("/", playerHandler.GetAll)
			players.GET("/:id", playerHandler.GetByID)
			players.PUT("/:id", playerHandler.Update)
			players.DELETE("/:id", playerHandler.Delete)
		}
	}

	router.Run(":8080")
}
