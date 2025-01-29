package main

import (
	"learning/internal/service"

	"learning/internal/handler"
	"learning/internal/repository/datastore"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Datastore client
	datastoreClient := datastore.NewDatastoreClient()

	// Initialize repositories
	playerRepo := datastore.NewDatastorePlayerRepository(datastoreClient)

	// Initialize services, controllers, handlers...
	playerService := service.NewPlayerService(playerRepo)
	playerHandler := handler.NewPlayerHandler(playerService)

	// Setup router
	router := gin.Default()

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
