package handler

import (
	"learning/internal/domain"
	"learning/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	service *service.PlayerService
}

func NewPlayerHanlder(s *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: s}
}

func (h *PlayerHandler) Create(c *gin.Context) {
	var newPlayer domain.Player
	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := h.service.Create(newPlayer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, player)
}
