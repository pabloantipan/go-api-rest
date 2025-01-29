package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Player struct {
// 	ID       string  `json:"id"`
// 	Name     string  `json:"name" binding:"required"`
// 	Age      int     `json:"age" binding:"required"`
// 	Position string  `json:"position" binding:"required"`
// 	Rating   float64 `json:"rating" binding:"required"`
// }

// var players = make(map[string]Player)

// func main() {
// 	router := gin.Default()

// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
// 	})

// 	v1 := router.Group("/api/V1")
// 	{
// 		v1.POST("/players", createPlayer)
// 		v1.GET("/players", getAllPlayers)
// 		v1.GET("/players/:id", getPlayerByID)
// 		v1.PUT("/players/:id", updatePlayer)
// 		// v1.DELETE("/players/:id", deletePlayer)
// 	}

// 	router.Run(":8080")
// }

// func createPlayer(c *gin.Context) {
// 	var newPlayer Player
// 	if err := c.ShouldBindJSON(&newPlayer); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	newPlayer.ID = fmt.Sprintf("%d", len(players)+1)
// 	players[newPlayer.ID] = newPlayer

// 	fmt.Println(players)

// 	c.JSON(http.StatusCreated, newPlayer)
// }

// func getAllPlayers(c *gin.Context) {
// 	playersList := make([]Player, 0, len(players))
// 	for _, player := range players {
// 		playersList = append(playersList, player)
// 	}

// 	c.JSON(http.StatusOK, playersList)
// }

// func getPlayerByID(c *gin.Context) {
// 	id := c.Param("id")
// 	player, ok := players[id]

// 	if !ok {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
// 	}

// 	c.JSON(http.StatusOK, player)
// }

// func updatePlayer(c *gin.Context) {
// 	id := c.Param("id")
// 	player, ok := players[id]

// 	if !ok {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
// 	}

// 	var updatedPlayer Player
// 	if err := c.ShouldBindJSON(&updatedPlayer); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	updatedPlayer.ID = player.ID
// 	players[id] = updatedPlayer

// 	c.JSON(http.StatusOK, updatedPlayer)
// }
