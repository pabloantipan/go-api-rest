package handler

import (
	"fmt"
	"net/http"
	"poc/internal/domain"
	"poc/internal/handler/interfaces"
	serviceInterfaces "poc/internal/service/interfaces"

	"github.com/gin-gonic/gin"
)

type AchievementHandler struct {
	service serviceInterfaces.AchievementService
}

func NewAchievementHandler(s serviceInterfaces.AchievementService) interfaces.AchievementHandler {
	return &AchievementHandler{service: s}
}

// CreateAchievement godoc
// @Summary Create a new Achievement
// @Description Create a new Achievement with the provided input data
// @Tags Achievements
// @Accept json
// @Produce json
// @Param Achievement body domain.Achievement true "Achievement object"
// @Success 201 {object} domain.Achievement
// @Failure 400 {object} gin.H
// @Router /Achievements [post]
func (h *AchievementHandler) Create(c *gin.Context) {
	var newAchievement domain.Achievement
	if err := c.ShouldBindJSON(&newAchievement); err != nil {
		fmt.Println("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Creating Achievement 1: ", newAchievement)

	Achievement, err := h.service.Create(newAchievement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Achievement)
}

// GetAchievements godoc
// @Summary Get all Achievements
// @Description Get all Achievements
// @Tags Achievements
// @Produce json
// @Success 200 {array} domain.Achievement
// @Failure 500 {object} gin.H
// @Router /Achievements [get]
func (h *AchievementHandler) GetAll(c *gin.Context) {
	Achievements, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Achievements)
}

func (h *AchievementHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	Achievement, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Achievement)
}

func (h *AchievementHandler) GetByUserID(c *gin.Context) {
	userId := c.Param("userId")
	Achievements, err := h.service.GetByUserID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Achievements)
}

func (h *AchievementHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var updatedAchievement domain.Achievement
	if err := c.ShouldBindJSON(&updatedAchievement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAchievement.ID = id

	Achievement, err := h.service.Update(updatedAchievement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Achievement)
}

func (h *AchievementHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
