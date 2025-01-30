package main

import (
	"log"
	"poc/config"
	"poc/internal/service"

	"poc/internal/handler"
	"poc/internal/repository/datastore"

	"github.com/gin-gonic/gin"

	_ "poc/docs/swagger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Player API
// @version         1.0
// @description     API Server for Player Management
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

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
	statRepo := datastore.NewDatastoreStatRepository(datastoreClient)
	achivementRepo := datastore.NewDatastoreAchievementRepository(datastoreClient)
	teamRepo := datastore.NewDatastoreTeamRepository(datastoreClient)

	// Initialize services
	playerService := service.NewPlayerService(playerRepo)
	statService := service.NewStatService(statRepo)
	achivementService := service.NewAchievementService(achivementRepo)
	teamRepoService := service.NewTeamService(teamRepo)

	// Initialize handlers
	playerHandler := handler.NewPlayerHandler(playerService)
	statHandler := handler.NewStatHandler(statService)
	achivementHandler := handler.NewAchievementHandler(achivementService)
	teamHandler := handler.NewTeamHandler(teamRepoService)

	profileHandler := handler.NewProfileHandler(statService, achivementService, teamRepoService)

	// Setup router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		// ginSwagger.URL("/swagger/doc.json"), // Point to your swagger json
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)

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

		teams := api.Group("/teams")
		{
			teams.POST("/", teamHandler.Create)
			teams.GET("/", teamHandler.GetAll)
			teams.GET("/:id", teamHandler.GetByID)
			teams.PUT("/:id", teamHandler.Update)
			teams.DELETE("/:id", teamHandler.Delete)
		}

		stats := api.Group("/stats")
		{
			stats.POST("/", statHandler.Create)
			stats.GET("/", statHandler.GetAll)
			stats.GET("/:id", statHandler.GetByID)
			stats.GET("/user/:id", statHandler.GetByUserID)
			stats.PUT("/:id", statHandler.Update)
			stats.DELETE("/:id", statHandler.Delete)
		}

		achivements := api.Group("/achivements")
		{
			achivements.POST("/", achivementHandler.Create)
			achivements.GET("/", achivementHandler.GetAll)
			achivements.GET("/:id", achivementHandler.GetByID)
			achivements.PUT("/:id", achivementHandler.Update)
			achivements.DELETE("/:id", achivementHandler.Delete)
		}

		profile := api.Group("/profile")
		{
			profile.GET("/:id", profileHandler.GetByID)
		}
	}

	router.Run(":8080")
}
