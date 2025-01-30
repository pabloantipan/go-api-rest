package handler

import (
	"fmt"
	"net/http"
	"poc/internal/domain"
	"poc/internal/handler/interfaces"

	serviceInterface "poc/internal/service/interfaces"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	// service *service.PlayerServi
	service serviceInterface.PlayerService
}

func NewPlayerHandler(s serviceInterface.PlayerService) interfaces.PlayerHandler {
	return &PlayerHandler{service: s}
}

// CreatePlayer godoc
// @Summary Create a new player
// @Description Create a new player with the provided input data
// @Tags players
// @Accept json
// @Produce json
// @Param player body domain.Player true "Player object"
// @Success 201 {object} domain.Player
// @Failure 400 {object} gin.H
// @Router /players [post]
func (h *PlayerHandler) Create(c *gin.Context) {
	var newPlayer domain.Player
	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		fmt.Println("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Creating player 1: ", newPlayer)

	player, err := h.service.Create(newPlayer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, player)
}

// GetPlayers godoc
// @Summary Get all players
// @Description Get all players
// @Tags players
// @Produce json
// @Success 200 {array} domain.Player
// @Failure 500 {object} gin.H
// @Router /players [get]
func (h *PlayerHandler) GetAll(c *gin.Context) {
	players, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func (h *PlayerHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	player, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var updatedPlayer domain.Player
	if err := c.ShouldBindJSON(&updatedPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPlayer.ID = id

	player, err := h.service.Update(updatedPlayer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
