package handler

import (
	"fmt"
	"net/http"
	"poc/internal/domain"
	"poc/internal/handler/interfaces"
	serviceInterfaces "poc/internal/service/interfaces"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	service serviceInterfaces.TeamService
}

func NewTeamHandler(s serviceInterfaces.TeamService) interfaces.TeamHandler {
	return &TeamHandler{service: s}
}

// CreateTeam godoc
// @Summary Create a new player
// @Description Create a new player with the provided input data
// @Tags players
// @Accept json
// @Produce json
// @Param player body domain.Team true "Team object"
// @Success 201 {object} domain.Team
// @Failure 400 {object} gin.H
// @Router /players [post]
func (h *TeamHandler) Create(c *gin.Context) {
	var newTeam domain.Team
	if err := c.ShouldBindJSON(&newTeam); err != nil {
		fmt.Println("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Creating player 1: ", newTeam)

	player, err := h.service.Create(newTeam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, player)
}

// GetTeams godoc
// @Summary Get all players
// @Description Get all players
// @Tags players
// @Produce json
// @Success 200 {array} domain.Team
// @Failure 500 {object} gin.H
// @Router /players [get]
func (h *TeamHandler) GetAll(c *gin.Context) {
	players, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func (h *TeamHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("userID")
	players, err := h.service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func (h *TeamHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	player, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *TeamHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var updatedTeam domain.Team
	if err := c.ShouldBindJSON(&updatedTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTeam.ID = id

	player, err := h.service.Update(updatedTeam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *TeamHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
