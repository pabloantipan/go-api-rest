package handler

import (
	"net/http"
	"poc/internal/domain"
	"poc/internal/handler/interfaces"

	serviceInterface "poc/internal/service/interfaces"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	statService        serviceInterface.StatService
	achievementService serviceInterface.AchievementService
	teamService        serviceInterface.TeamService
}

func NewProfileHandler(
	statService serviceInterface.StatService,
	achievementService serviceInterface.AchievementService,
	teamService serviceInterface.TeamService,
) interfaces.ProfileHandler {
	return &ProfileHandler{
		statService:        statService,
		achievementService: achievementService,
		teamService:        teamService,
	}
}

func (h *ProfileHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	type Result struct {
		Data  interface{}
		Error error
	}

	statsChan := make(chan Result)
	achievementsChan := make(chan Result)
	teamsChan := make(chan Result)

	go func() {
		stats, err := h.statService.GetByUserID(id)
		statsChan <- Result{Data: stats, Error: err}
	}()

	go func() {
		achievements, err := h.achievementService.GetByUserID(id)
		achievementsChan <- Result{Data: achievements, Error: err}
	}()

	go func() {
		teams, err := h.teamService.GetByID(id)
		teamsChan <- Result{Data: teams, Error: err}
	}()

	timeout := time.After(3 * time.Second)
	remaining := 3

	var profile domain.Profile

	for remaining > 0 {
		select {
		case result := <-statsChan:
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats"})
				return
			}

			profile.Stats = result.Data
			remaining--

		case result := <-achievementsChan:
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch achievements"})
				return
			}

			profile.Achievements = result.Data
			remaining--

		case result := <-teamsChan:
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teams"})
				return
			}
			profile.Team = result.Data
			remaining--

		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
			return

		case <-ctx.Done():
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request cancelled"})
			return
		}
	}

	c.JSON(http.StatusOK, profile)

}
